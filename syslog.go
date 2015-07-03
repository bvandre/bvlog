package bvlog

import (
	"fmt"
	"log/syslog"
	"os"
)

type syslogger struct {
	*syslog.Writer
}

func NewSyslogLogger() (Logger, error) {
	l, err := syslog.New(syslog.LOG_INFO, os.Args[0])
	if err != nil {
		return nil, err
	}
	return &syslogger{
		Writer: l,
	}, err
}

func (s *syslogger) Fatalf(exit bool, format string, v ...interface{}) error {
	err := s.Err(fmt.Sprintf(format, v...))
	if exit {
		os.Exit(1)
	}
	return err
}

func (s *syslogger) Fatal(exit bool, v ...interface{}) error {
	err := s.Err(fmt.Sprint(v...))
	if exit {
		os.Exit(1)
	}
	return err
}

func (s *syslogger) Infof(format string, v ...interface{}) error {
	return s.Writer.Info(fmt.Sprintf(format, v...))
}

func (s *syslogger) Info(v ...interface{}) error {
	return s.Writer.Info(fmt.Sprint(v...))
}

func (s *syslogger) Warnf(format string, v ...interface{}) error {
	return s.Writer.Warning(fmt.Sprintf(format, v...))
}

func (s *syslogger) Warn(v ...interface{}) error {
	return s.Writer.Warning(fmt.Sprint(v...))
}
