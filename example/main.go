package main

import (
	"github.com/robertgzr/rlog"
)

func init() {
	rlog.Init()
	rlog.SetColor(false)
}

func main() {
	l := rlog.New()
	l.SetLogTime(true)
	l.AddCtx("magic", 1000)
	l.Debug("debug")
	l.Info("info")
	l.Warn("warning")
	l.Error("error")
	l.Crit("critical")
}
