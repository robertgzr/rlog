package subpkg

import "github.com/robertgzr/rlog"

func Do() {
	rlog.Debug("debug from subpackage")
	rlog.Error("error from subpackage")
}
