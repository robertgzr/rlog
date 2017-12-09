package rlog

import (
	"os"
)

const defaultEnvKey = "GO_LOG"

var (
	envKey string
	global Logger
)

func init() {
	envKey = defaultEnvKey
	global = new(logger).With(
		OutputOpt(os.Stdout),
		MaxLvlOpt(LvlDebug),
		DisableColorOpt(false),
	)

}

func New(opt ...Option) Logger {
	return global.New(opt...)
}
func With(opt ...Option) Logger {
	return global.With(opt...)
}
func Crit(a ...interface{}) {
	global.Crit(a...)
}
func Error(a ...interface{}) {
	global.Error(a...)
}
func Warn(a ...interface{}) {
	global.Warn(a...)
}
func Info(a ...interface{}) {
	global.Info(a...)
}
func Debug(a ...interface{}) {
	global.Debug(a...)
}
