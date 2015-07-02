package bvlog

import (
	"io"
)

var l logger

//The logger interface will be implemented
//by a few different types of loggers
//
//This should abstract away the individual
//logging details
type logger interface {
	io.Writer
	io.Closer
	Infof(format string, v ...interface{}) error
	Info(v ...interface{}) error
	Warnf(format string, v ...interface{}) error
	Warn(v ...interface{}) error
	Fatalf(exit bool, format string, v ...interface{}) error
	Fatal(exit bool, v ...interface{}) error
}

func init() {
	li, err := NewJournalLogger()
	if err == nil {
		l = li
		return
	}
	li, err = NewSyslogLogger()
	l = li
}

func Infof(format string, v ...interface{}) error {
	return l.Infof(format, v...)
}
func Info(v ...interface{}) error {
	return l.Info(v...)
}
func Warnf(format string, v ...interface{}) error {
	return l.Warnf(format, v...)
}
func Warn(v ...interface{}) error {
	return l.Warn(v...)
}
func Fatalf(exit bool, format string, v ...interface{}) error {
	return l.Fatalf(exit, format, v...)
}
func Fatal(exit bool, v ...interface{}) error {
	return l.Fatal(exit, v...)
}
