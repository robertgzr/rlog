package rlog

import (
	"io"
)

type Option interface {
	Apply(*logger)
}

type output struct {
	w io.Writer
}

func SetOutputOpt(w io.Writer) Option {
	return &output{w}
}
func (o *output) Apply(l *logger) {
	l.out = o.w
}

type maxLvl struct {
	max Lvl
}

func SetMaxLvlOpt(max Lvl) Option {
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
	l.color = !o.disable
}

type logtime struct{}

func LogTimeOpt() Option {
	return new(logtime)
}
func (o *logtime) Apply(l *logger) {
	l.logtime = true
}

type resetCtx struct{}

func ResetOpt() Option {
	return new(resetCtx)
}
func (o *resetCtx) Apply(l *logger) {
	l.ctx = newCtx()
}

type setOpenerCloser struct {
	opener, closer string
}

func SetOpenerCloserOpt(opener, closer string) Option {
	return &setOpenerCloser{opener, closer}
}
func (o *setOpenerCloser) Apply(l *logger) {
	l.ctx.op = o.opener
	l.ctx.cl = o.closer
}

type setDelimiter struct {
	delimiter string
}

func SetDelimiterOpt(delimiter string) Option {
	return &setDelimiter{delimiter}
}
func (o *setDelimiter) Apply(l *logger) {
	l.ctx.del = o.delimiter
}
