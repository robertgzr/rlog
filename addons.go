package rlog

import "time"

var timeFormat = "2006-01-02T15:04:05-0700"

func (l *Logger) SetLogTime(b bool) {
	if b {
		l.AddCtx("t", func() string { return time.Now().Format(timeFormat) })
	} else {
		l.ctx.Delete("t")
	}
}
