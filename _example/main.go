package main

import (
	"errors"

	"github.com/apex/log"
	"github.com/robertgzr/rlog"

	"./subpkg"
)

func main() {
	log.SetHandler(rlog.Default)
	log.Info("rlog can be used from apex/log by setting it as the log handler")
	log.Warn("or on it's own... via a limited wrapper around apex/log\n")

	println("rlog demo\n=========")
	println("\nlog levels:")
	rlog.Debug("debug")
	rlog.Info("info")
	rlog.Warn("warning")
	rlog.Error("error")

	// println("\nwith context:")
	val_1 := 1000
	val_2 := struct{ a int }{a: 9}

	rlog.WithField("key", "value").Debug("structured logging")
	rlog.WithFields(rlog.Fields{
		"val_1": val_1,
		"val_2": val_2,
	}).Error("ouch!")

	errorFunc()

	println("\nwe can log from subpackages as well")
	subpkg.Do()
	println("\n")

	rlog.Fatal("fatal log exits process")
}

func errorFunc() error {
	var err = errors.New("fatal error error")
	rlog.Trace("does something that produces an error").Stop(&err)
	return nil
}
