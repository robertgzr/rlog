package rlog

import (
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
)

// Lvl type and funcs from github.com/inconshreveable/log15:
type Lvl int

const (
	LvlCrit Lvl = iota
	LvlError
	LvlWarn
	LvlInfo
	LvlDebug
)

func (l Lvl) String() string {
	switch l {
	case LvlDebug:
		return "dbug"
	case LvlInfo:
		return "info"
	case LvlWarn:
		return "warn"
	case LvlError:
		return "eror"
	case LvlCrit:
		return "crit"
	default:
		panic("bad level")
	}
}

func LvlFromString(lvlString string) (Lvl, error) {
	switch lvlString {
	case "debug", "dbug":
		return LvlDebug, nil
	case "info":
		return LvlInfo, nil
	case "warn":
		return LvlWarn, nil
	case "error", "eror":
		return LvlError, nil
	case "crit":
		return LvlCrit, nil
	default:
		return LvlDebug, fmt.Errorf("Unknown level: %v", lvlString)
	}
}

type Logger struct {
	out    io.Writer
	maxLvl Lvl
	ctx    *Ctx
	color  bool
}

func NewLogger(wr io.Writer, lvl Lvl, ctx ...interface{}) *Logger {
	l := &Logger{
		out:    wr,
		maxLvl: lvl,
		ctx:    newCtx(),
	}
	l.ctx.Add(ctx)
	return l
}

func (l *Logger) New(ctx ...interface{}) *Logger {
	new := &Logger{
		out:    l.out,
		maxLvl: l.maxLvl,
		ctx:    l.ctx,
	}
	new.ctx.Add(ctx)
	return new
}

func (l *Logger) SetOutput(wr io.Writer) {
	l.out = wr
}
func (l *Logger) SetColor(b bool) {
	l.color = b
}
func (l *Logger) SetMaxLvl(lvl Lvl) {
	l.maxLvl = lvl
}
func (l *Logger) AddCtx(ctx ...interface{}) {
	l.ctx.Add(ctx)
}

func (l *Logger) write(lvl Lvl, a ...interface{}) {
	// apply maxLvl filter
	if lvl > l.maxLvl {
		return
	}
	// format the output
	var attr color.Attribute
	switch lvl {
	case LvlCrit:
		attr = color.FgMagenta
	case LvlError:
		attr = color.FgRed
	case LvlWarn:
		attr = color.FgYellow
	case LvlInfo:
		attr = color.Reset
	case LvlDebug:
		attr = color.FgCyan
	}
	fmtr := color.New(attr)
	if !l.color {
		fmtr.DisableColor()
	}
	fmtr.Fprintf(l.out, "%s|%s %s", strings.ToUpper(lvl.String()), l.ctx.String(), fmt.Sprintln(a...))
}

func (l *Logger) Crit(a ...interface{}) {
	l.write(LvlCrit, a...)
}

func (l *Logger) Error(a ...interface{}) {
	l.write(LvlError, a...)
}

func (l *Logger) Warn(a ...interface{}) {
	l.write(LvlWarn, a...)
}

func (l *Logger) Info(a ...interface{}) {
	l.write(LvlInfo, a...)
}

func (l *Logger) Debug(a ...interface{}) {
	l.write(LvlDebug, a...)
}
