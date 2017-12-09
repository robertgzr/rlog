package main

import (
	"github.com/robertgzr/rlog"
)

func init() {
	rlog.ParseEnv()
}

func main() {
	l := rlog.New().With(rlog.LogTimeOpt(), rlog.CtxOpt("magic", 1000))
	l.Debug("debug")
	l.Info("info")
	l.Warn("warning")
	l.Error("error")
	l.Crit("critical")
}
