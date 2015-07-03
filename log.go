package bvlog

import (
	"io"
)

var l Logger

//The logger interface will be implemented
//by a few different types of loggers
//
//This should abstract away the individual
//backend loggers
type Logger interface {
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

//Infof and Info use the info prefix for logging.
func Infof(format string, v ...interface{}) error {
	return l.Infof(format, v...)
}
func Info(v ...interface{}) error {
	return l.Info(v...)
}

//Warnf and Warn use the warning prefix for logging.
func Warnf(format string, v ...interface{}) error {
	return l.Warnf(format, v...)
}
func Warn(v ...interface{}) error {
	return l.Warn(v...)
}

//Fatalf and Fatal use the error prefix for logging.
//If you want the process to exit with an error code
//pass true for exit.
func Fatalf(exit bool, format string, v ...interface{}) error {
	return l.Fatalf(exit, format, v...)
}
func Fatal(exit bool, v ...interface{}) error {
	return l.Fatal(exit, v...)
}
