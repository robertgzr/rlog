package rlog

import (
	"fmt"
	"io"
	"os"
	"sync"

	apex "github.com/apex/log"

	"github.com/fatih/color"
	colorable "github.com/mattn/go-colorable"
)

var bold = color.New(color.Bold)

var colors = [...]*color.Color{
	apex.DebugLevel: color.New(color.FgWhite),
	apex.InfoLevel:  color.New(color.FgBlue),
	apex.WarnLevel:  color.New(color.FgYellow),
	apex.ErrorLevel: color.New(color.FgRed),
	apex.FatalLevel: color.New(color.FgMagenta),
}

// Strings mapping.
var strings = [...]string{
	apex.DebugLevel: "DBUG",
	apex.InfoLevel:  "INFO",
	apex.WarnLevel:  "WARN",
	apex.ErrorLevel: "EROR",
	apex.FatalLevel: "CRIT",
}

type Handler struct {
	mu     sync.Mutex
	Lvl    apex.Level
	Writer io.Writer
}

func NewHandler(w io.Writer) apex.Handler {
	if f, ok := w.(*os.File); ok {
		return &Handler{
			Writer: colorable.NewColorable(f),
		}
	}

	return &Handler{
		Writer: w,
	}
}

func (h *Handler) HandleLog(e *apex.Entry) error {
	var (
		color = colors[e.Level]
		level = strings[e.Level]
		names = e.Fields.Names()
		msg   = e.Message
	)

	h.mu.Lock()
	defer h.mu.Unlock()

	if e.Level == apex.FatalLevel {
		msg = bold.Sprint(msg)
	}
	color.Fprintf(h.Writer, "%s %-25s", bold.Sprintf("%*s|", 4, level), msg)

	for _, name := range names {
		if name == "source" {
			continue
		}
		fmt.Fprintf(h.Writer, " %s=%#v", color.Sprint(name), e.Fields.Get(name))
	}

	fmt.Fprintln(h.Writer)
	return nil
}
