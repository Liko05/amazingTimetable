package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

// MyFormatter is a custom formatter for logrus
// Taken from here: https://stackoverflow.com/questions/48971780/how-to-change-the-format-of-log-output-in-logrus
// all it does is configure the logrus logger to output pretty log format
type MyFormatter struct {
	log.TextFormatter
}

// Format formats the log entry
func (f *MyFormatter) Format(entry *log.Entry) ([]byte, error) {
	// this whole mess of dealing with ansi color codes is required if you want the colored output otherwise you will lose colors in the log levels
	var levelColor int
	switch entry.Level {
	case log.DebugLevel, log.TraceLevel:
		levelColor = 31 // gray
	case log.WarnLevel:
		levelColor = 33 // yellow
	case log.ErrorLevel, log.FatalLevel, log.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
	return []byte(fmt.Sprintf("[%s] - \x1b[%dm%s\x1b[0m - %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

// InitLogger initializes the logger with the given log level
func InitLogger(level log.Level) {
	log.SetFormatter(&MyFormatter{log.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		ForceColors:            true,
		DisableLevelTruncation: true,
	}})

	log.SetOutput(os.Stdout)
	log.SetLevel(level)
}
