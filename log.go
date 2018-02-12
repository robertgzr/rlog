package rlog

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

const timeFormat = "060102.150405"

type Logger interface {
	New(...interface{}) Logger
	With(...interface{}) Logger

	Crit(...interface{})
	Error(...interface{})
	Warn(...interface{})
	Info(...interface{})
	Debug(...interface{})
}

type logger struct {
	out     io.Writer
	maxLvl  Lvl
	ctx     Ctx
	logtime bool
	color   bool
}

func newlogger() logger {
	return logger{
		ctx:    newCtx(),
		out:    os.Stdout,
		maxLvl: LvlDebug,
		color:  true,
	}
}

func (l logger) write(lvl Lvl, a ...interface{}) {
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

	a = append(a, l.ctx.String())
	fmtr.Fprintf(l.out, "%s|%s %s",
		strings.ToUpper(lvl.String()),
		func() string {
			if l.logtime {
				return fmt.Sprintf("%s|", time.Now().Format(timeFormat))
			}
			return ""
		}(),
		fmt.Sprintln(a...),
	)
}

// New creates a new logger and calls With(opt) on it.
//
// See documentation of With for details.
func (l logger) New(opt ...interface{}) Logger {
	// return l.With(ResetOpt()).With(opt...)
	return l.With(opt...)
}

// With returns a copy of the logger with changed options and context.
// It doesn't modify the receiver.
//
// Takes a list of interface{}.
// You can supply Option values optionally folled by (string, interface) pairs to set context
func (l logger) With(opt ...interface{}) Logger {
	for i, o := range opt {
		if option, ok := o.(Option); ok {
			option.Apply(&l)
		} else {
			l.ctx.Add(opt[i:])
			break
		}
	}
	return l
}

func (l logger) Crit(a ...interface{}) {
	l.write(LvlCrit, a...)
}

func (l logger) Error(a ...interface{}) {
	l.write(LvlError, a...)
}

func (l logger) Warn(a ...interface{}) {
	l.write(LvlWarn, a...)
}

func (l logger) Info(a ...interface{}) {
	l.write(LvlInfo, a...)
}

func (l logger) Debug(a ...interface{}) {
	l.write(LvlDebug, a...)
}
