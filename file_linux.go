package log

import (
	"os"
	"strings"
)

// TODO
func NewFileLogger(config *LogConfig) (Logger, error) {
	path := "/var/log/" + strings.ToLower(config.DirectoryName)
	err = os.MkdirAll(path, os.FileMode(0777))
	if err != nil {
		stdLog.Fatalf("error creating log directory: %v", err)
		return nil, err
	}

	f, err := os.OpenFile(path+"/log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		stdLog.Fatalf("error opening file: %v", err)
		return nil, err
	}
	stdLog.SetOutput(f)

	return &fileLogger{w: f}, nil
}
