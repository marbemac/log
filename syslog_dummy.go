// +build windows nacl plan9

package log

import (
	"os"
)

func NewSysLogger(config *LogConfig) (Logger, error) {
}

type dummyLogger struct{}

func NewConsoleLogger(config *LogConfig) (Logger, error) {
	return &dummyLogger{}, nil
}

func (l *dummyLogger) Writer(sev Severity) io.Writer {
	return os.Stdout
}

func (l *dummyLogger) Infof(format string, args ...interface{}) {}

func (l *dummyLogger) Warningf(format string, args ...interface{}) {}

func (l *dummyLogger) Errorf(format string, args ...interface{}) {}

func (l *dummyLogger) Fatalf(format string, args ...interface{}) {}
