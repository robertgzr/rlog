package rlog

import "os"

func ParseEnv() {
	maxLvl := parseEnv()
	global = global.With(SetMaxLvlOpt(maxLvl))
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
