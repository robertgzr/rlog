package rlog

import (
	"os"

	apex "github.com/apex/log"
	"github.com/apex/log/handlers/level"
)

const defaultEnvKey = "GO_LOG"

func ParseEnv() {
	env := os.Getenv(defaultEnvKey)
	if env == "" {
		return
	}

	maxLvl := levelFromString(env)
	std.Handler = level.New(std.Handler, maxLvl)
}

func levelFromString(env string) apex.Level {
	if lvl, err := apex.ParseLevel(env); err != nil {
		return apex.DebugLevel
	} else {
		return lvl
	}
}
