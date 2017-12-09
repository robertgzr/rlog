package rlog

import (
	"io"
	"time"
)

type Option interface {
	Apply(*logger)
}

type output struct {
	w io.Writer
}

func OutputOpt(w io.Writer) Option {
	return &output{w}
}
func (o *output) Apply(l *logger) {
	l.out = o.w
}

type maxLvl struct {
	max Lvl
}

func MaxLvlOpt(max Lvl) Option {
	return &maxLvl{max}
}
func (o *maxLvl) Apply(l *logger) {
	l.maxLvl = o.max
}

type disableColor struct {
	disable bool
}

func DisableColorOpt(disable bool) Option {
	return &disableColor{disable}
}
func (o *disableColor) Apply(l *logger) {
	l.color = o.disable
}

type ctx struct {
	c []interface{}
}

func CtxOpt(c ...interface{}) Option {
	return &ctx{c}
}
func (o *ctx) Apply(l *logger) {
	if l.ctx == nil {
		l.ctx = newCtx()
	}
	l.ctx.Add(o.c)
}

type logtime struct{}

func LogTimeOpt() Option {
	return new(logtime)
}

const timeFormat = "2006-01-02T15:04:05-0700"

func (o *logtime) Apply(l *logger) {
	l.With(CtxOpt("t", func() string { return time.Now().Format(timeFormat) }))
}
