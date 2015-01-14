package log

import (
	"fmt"
	"io"
	"time"
)

// fileLogger outputs the logs to the underlying writer
type fileLogger struct {
	w io.Writer
}

func (l *fileLogger) Writer(sev Severity) io.Writer {
	return l
}

func (l *fileLogger) Write(val []byte) (int, error) {
	return io.WriteString(
		l.w,
		fmt.Sprintf("%v: %v\n", time.Now().UTC().Format(time.StampMilli),
			string(val)))
}

func (l *fileLogger) Infof(format string, args ...interface{}) {
	infof(1, l, format, args...)
}

func (l *fileLogger) Warningf(format string, args ...interface{}) {
	warningf(1, l, format, args...)
}

func (l *fileLogger) Errorf(format string, args ...interface{}) {
	errorf(1, l, format, args...)
}

func (l *fileLogger) Fatalf(format string, args ...interface{}) {
	fatalf(1, l, format, args...)
}
