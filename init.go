package rlog

var global Logger

func init() {
	global = newlogger()
}

func New(opt ...interface{}) Logger {
	return global.New(opt...)
}
func With(opt ...interface{}) Logger {
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
