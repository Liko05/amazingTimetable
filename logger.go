package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

// Taken from here: https://stackoverflow.com/questions/48971780/how-to-change-the-format-of-log-output-in-logrus
// all it does is configure the logrus logger to output pretty log format

type myFormatter struct {
	log.TextFormatter
}

func (f *myFormatter) Format(entry *log.Entry) ([]byte, error) {
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

func initLogger(level log.Level) {
	log.SetFormatter(&myFormatter{log.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		ForceColors:            true,
		DisableLevelTruncation: true,
	}})

	//open file
	//file, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	//
	//if err == nil {
	//	log.SetOutput(file)
	//} else {
	//	log.Info("Failed to log to file, using default stderr")
	//}

	log.SetOutput(os.Stdout)
	log.SetLevel(level)
}
