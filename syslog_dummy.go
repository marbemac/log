// +build windows nacl plan9

package log

import (
	"io"
	"os"
)

func NewSysLogger(config *LogConfig) (Logger, error) {
	return &dummyLogger{}, nil
}

type dummyLogger struct{}

func (l *dummyLogger) Writer(sev Severity) io.Writer {
	return os.Stdout
}

func (l *dummyLogger) Infof(format string, args ...interface{}) {}

func (l *dummyLogger) Warningf(format string, args ...interface{}) {}

func (l *dummyLogger) Errorf(format string, args ...interface{}) {}

func (l *dummyLogger) Fatalf(format string, args ...interface{}) {}
