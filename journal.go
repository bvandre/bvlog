package bvlog

import (
	"errors"
	"fmt"
	"github.com/coreos/go-systemd/journal"
	"os"
)

type jlogger struct {
}

func NewJournalLogger() (Logger, error) {
	if !journal.Enabled() {
		return nil, errors.New("could not connect to journald socket")
	}
	return &jlogger{}, nil
}

func (j *jlogger) Write(b []byte) (int, error) {
	err := journal.Send(string(b), journal.PriInfo, nil)
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

func (j *jlogger) Close() error {
	//for now we will just make this a no-op
	return nil
}

func (j *jlogger) Fatalf(exit bool, format string, v ...interface{}) error {
	err := journal.Send(fmt.Sprintf(format, v...), journal.PriErr, nil)
	if exit {
		os.Exit(1)
	}
	return err
}

func (j *jlogger) Fatal(exit bool, v ...interface{}) error {
	err := journal.Send(fmt.Sprint(v...), journal.PriErr, nil)
	if exit {
		os.Exit(1)
	}
	return err
}

func (j *jlogger) Info(v ...interface{}) error {
	return journal.Send(fmt.Sprint(v...), journal.PriInfo, nil)
}

func (j *jlogger) Infof(format string, v ...interface{}) error {
	return journal.Send(fmt.Sprintf(format, v...), journal.PriInfo, nil)
}

func (j *jlogger) Warn(v ...interface{}) error {
	return journal.Send(fmt.Sprint(v...), journal.PriWarning, nil)
}

func (j *jlogger) Warnf(format string, v ...interface{}) error {
	return journal.Send(fmt.Sprintf(format, v...), journal.PriWarning, nil)
}
