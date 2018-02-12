package main

import (
	"github.com/robertgzr/rlog"
)

func main() {
	rlog.ParseEnv()
	val_1 := 1000
	val_2 := struct{}{}

	rlog.Debug("using the global logger")
	rlog.With("ctx", "is here").Info("info with context")
	rlog.Warn("warning", val_1, val_2)
	rlog.Error("error")
	rlog.With("val_1", val_1, "val_2", val_2).Crit("critical")

	l := rlog.New(rlog.LogTimeOpt())
	l.Debug("using a new logger")
	l.Info("set the logtime option")
	l.Warn("warning")
	l = l.With(rlog.SetOpenerCloserOpt(">> ", ""), rlog.SetDelimiterOpt(", "), "fail", true)
	l.Error("error")
	l.With("more_ctx", 0.0007).Crit("critical")
}
