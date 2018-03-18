package rlog

import (
	"os"
	"strings"

	apex "github.com/apex/log"
)

const defaultEnvKey = "GO_LOG"

func ParseEnv() {
	env := os.Getenv(defaultEnvKey)
	if env == "" {
		return
	}

	lvlmap := levelFromString(env)
	std.Handler = NewHandlerWithPkgMap(defaultWriter, lvlmap)
}

func levelFromString(env string) LvlMap {
	lvlmap := LvlMap{"*": apex.DebugLevel}

	if env == "" {
		return lvlmap
	}

	lvlpairs := strings.Split(env, ",")
	for _, lvlpair := range lvlpairs {
		r := strings.Split(lvlpair, "=")
		if len(r) < 2 {
			lvl := parseLvl(r[0])
			lvlmap["*"] = lvl
			continue
		}
		pkg := r[0]
		lvl := parseLvl(r[1])
		lvlmap[pkg] = lvl
	}
	return lvlmap
}

func parseLvl(str string) apex.Level {
	if lvl, err := apex.ParseLevel(str); err != nil {
		return apex.DebugLevel
	} else {
		return lvl
	}
}
