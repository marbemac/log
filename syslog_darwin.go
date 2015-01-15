package log

import (
	"log/syslog"
)

// Syslogger sends all your logs to syslog
// Note: the logs are going to MAIL_LOG facility
type sysLogger struct {
	info *syslog.Writer
	warn *syslog.Writer
	err  *syslog.Writer
}
