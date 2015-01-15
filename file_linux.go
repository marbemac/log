package log

import (
	stdLog "log"
	"os"
	"strings"
)

// TODO
func NewFileLogger(config *LogConfig) (Logger, error) {
	f, err := os.OpenFile("/var/log/"+strings.ToLower(config.DirectoryName)+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		stdLog.Fatalf("error opening file: %v", err)
		return nil, err
	}
	stdLog.SetOutput(f)

	return &fileLogger{w: f}, nil
}
