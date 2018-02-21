package rlog

import "os"

const defaultEnvKey = "GO_LOG"

func ParseEnv(a ...string) {
	envKey := defaultEnvKey
	if len(a) == 1 {
		envKey = a[0]
	}
	maxLvl := parseEnv(envKey)
	global = global.With(MaxLvl(maxLvl))
}

func parseEnv(e string) Lvl {
	lvlstr := os.Getenv(e)
	if lvlstr == "" {
		return LvlDebug
	}
	lvl, err := LvlFromString(lvlstr)
	if err != nil {
		return LvlDebug
	}
	return lvl
}
