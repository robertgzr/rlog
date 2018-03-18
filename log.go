package rlog

import (
	"os"

	apex "github.com/apex/log"
)

var (
	std = Logger{apex.Logger{
		Handler: NewHandler(os.Stderr),
		Level:   apex.DebugLevel,
	}}
)

type Logger struct {
	apex.Logger
}

type Fields apex.Fields

func Debug(msg string) {
	std.Debug(msg)
}
func Info(msg string) {
	std.Info(msg)
}
func Warn(msg string) {
	std.Warn(msg)
}
func Error(msg string) {
	std.Error(msg)
}
func Fatal(msg string) {
	std.Fatal(msg)
}

func Trace(msg string) *apex.Entry {
	return std.Trace(msg)
}
func WithError(err error) *apex.Entry {
	return std.WithError(err)
}
func WithField(key string, value interface{}) *apex.Entry {
	return std.WithField(key, value)
}
func WithFields(fields Fields) *apex.Entry {
	return std.WithFields(apex.Fields(fields))
}
