package log

import (
	stdLog "log"
	"os"
	"os/user"
	"strings"
)

func NewFileLogger(config *LogConfig) (Logger, error) {
	currentUser, err := user.Current()
	if err != nil {
		stdLog.Fatalf("error reading current system user: %v", err)
		return nil, err
	}
	path := "C:\\Users\\" + currentUser.Username + "\\AppData\\" + config.DirectoryName
	err = os.MkdirAll(path, os.FileMode(0777))
	if err != nil {
		stdLog.Fatalf("error creating support directory: %v", err)
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
