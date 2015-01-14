package log

import (
	"os"
)

// TODO
func NewFileLogger(config *LogConfig) (Logger, error) {
	return &fileLogger{w: os.Stdout}, nil
}
