package main

import (
	"os"

	"github.com/robertgzr/rlog"
)

func main() {
	rlog.ParseEnv()
	println("rlog demo\n=========")
	println("\nlog levels:")
	rlog.Debug("debug")
	rlog.Info("info")
	rlog.Warn("warning")
	rlog.Error("error")
	rlog.Crit("critical")

	println("\nworks with more than `string`")
	val_1 := 1000
	val_2 := struct{ a int }{a: 9}
	rlog.Warn("debug", val_1, val_2)

	println("\nwith context:")
	rlog.With("key", "value").Debug("logging with context")
	rlog.With("val_1", val_1, "val_2", val_2).Error("ouch!")

	println("\nyou can customize the output:")
	rl := rlog.New(rlog.LogTime(true), rlog.Output(os.Stderr))
	rl.Debug("using a new logger that also logs the time")
	rl = rl.With(rlog.OpenerCloser(">> ", ""), rlog.Delimiter(", "), "has_name", "rl")
	rl.Info("and has persistent context")
	rl.Warn("and custom ctx opener/closer/delimiter")
	rl.With("more_ctx", 0.0007).Error("critical")
}
