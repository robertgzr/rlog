package rlog

import (
	"io"
	"os"
)

const defaultEnvKey = "GO_LOG"

var envKey string

func init() {
	envKey = defaultEnvKey
	globalLogger = NewLogger(os.Stdout, LvlDebug)
}

var globalLogger *Logger

func Init() {
	maxLvl := parseEnv()
	globalLogger.SetMaxLvl(maxLvl)
}

func parseEnv() Lvl {
	lvlstr := os.Getenv(envKey)
	if lvlstr == "" {
		return LvlDebug
	}
	lvl, err := LvlFromString(lvlstr)
	if err != nil {
		return LvlDebug
	}
	return lvl
}

func New(ctx ...interface{}) *Logger {
	return globalLogger.New(ctx...)
}

func SetOutput(wr io.Writer) {
	globalLogger.SetOutput(wr)
}

func SetColor(b bool) {
	globalLogger.SetColor(b)
}

func SetMaxLvl(lvl Lvl) {
	globalLogger.SetMaxLvl(lvl)
}

func SetLogTime(b bool) {
	globalLogger.SetLogTime(b)
}

func Crit(a ...interface{}) {
	globalLogger.write(LvlCrit, a...)
}

func Error(a ...interface{}) {
	globalLogger.write(LvlError, a...)
}

func Warn(a ...interface{}) {
	globalLogger.write(LvlWarn, a...)
}

func Info(a ...interface{}) {
	globalLogger.write(LvlInfo, a...)
}

func Debug(a ...interface{}) {
	globalLogger.write(LvlDebug, a...)
}
