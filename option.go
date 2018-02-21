package rlog

import (
	"io"
)

type Option interface {
	Apply(*logger)
}

type outputOpt struct{ io.Writer }

func Output(w io.Writer) Option {
	return outputOpt{w}
}
func (o outputOpt) Apply(l *logger) {
	l.out = o.Writer
}

type maxLvl struct {
	max Lvl
}

func MaxLvl( max Lvl) Option {
	return maxLvl{max}
}
func (o maxLvl) Apply(l *logger) {
	l.maxLvl = o.max
}

type colorOpt bool

func Color(enabled bool) Option {
	return colorOpt(enabled)
}
func (o colorOpt) Apply(l *logger) {
	l.color = bool(o)
}

type logtimeOpt bool

func LogTime(enabled bool) Option {
	return logtimeOpt(enabled)
}
func (o logtimeOpt) Apply(l *logger) {
	l.logtime = bool(o)
}

type resetCtxOpt struct{}

func ResetOpt() Option {
	return resetCtxOpt{}
}
func (o resetCtxOpt) Apply(l *logger) {
	l.ctx = newCtx()
}

type openerCloserOpt struct {
	opener, closer string
}

func OpenerCloser(opener, closer string) Option {
	return openerCloserOpt{opener, closer}
}
func (o openerCloserOpt) Apply(l *logger) {
	l.ctx.op = o.opener
	l.ctx.cl = o.closer
}

type delimiterOpt string

func Delimiter(del string) Option {
	return delimiterOpt(del)
}
func (o delimiterOpt) Apply(l *logger) {
	l.ctx.del = string(o)
}
