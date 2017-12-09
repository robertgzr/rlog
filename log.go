package rlog

import (
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
)

type Logger interface {
	New(...Option) Logger
	With(...Option) Logger

	Crit(...interface{})
	Error(...interface{})
	Warn(...interface{})
	Info(...interface{})
	Debug(...interface{})
}

type logger struct {
	out    io.Writer
	maxLvl Lvl
	ctx    *Ctx
	color  bool
}

func (l *logger) write(lvl Lvl, a ...interface{}) {
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

func (l *logger) New(opt ...Option) Logger {
	new := l
	return new.With(opt...)
}

func (l *logger) With(opt ...Option) Logger {
	for _, o := range opt {
		o.Apply(l)
	}
	return l
}

func (l *logger) Crit(a ...interface{}) {
	l.write(LvlCrit, a...)
}

func (l *logger) Error(a ...interface{}) {
	l.write(LvlError, a...)
}

func (l *logger) Warn(a ...interface{}) {
	l.write(LvlWarn, a...)
}

func (l *logger) Info(a ...interface{}) {
	l.write(LvlInfo, a...)
}

func (l *logger) Debug(a ...interface{}) {
	l.write(LvlDebug, a...)
}
