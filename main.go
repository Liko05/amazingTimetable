package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

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

func initLogger() {
	log.SetFormatter(&myFormatter{log.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		ForceColors:            true,
		DisableLevelTruncation: true,
	}})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

var shouldFinish = make(chan bool)
var generatedOptions = make(chan int, 100000000)

func main() {
	initLogger()
	log.Debug("Number of processors: " + strconv.Itoa(runtime.NumCPU()))
	log.Debug("Number of threads: " + strconv.Itoa(runtime.GOMAXPROCS(0)))

	initTimeTable()

	subjects := retrieveArrayOfSubjectsWithouPauses()
	//for _, subject := range subjects {
	//	log.Debug(subject.name)
	//}

	//idk := copy(subjects, subjects)

	log.Info("Number of subjects: " + strconv.Itoa(len(subjects)))

	timeStart := time.Now()

	for i := 0; i < 3; i++ {
		go generateCalendars()
	}

	<-shouldFinish

	log.Info("Time elapsed for generating 1000000 options: " + time.Since(timeStart).String())

}

func generateCalendars() {
	for {
		_ = generateNewTimeTable(retrieveArrayOfSubjectsWithouPauses())
		generatedOptions <- 1
		log.Debug("Generated options: " + strconv.Itoa(len(generatedOptions)))
		if len(generatedOptions) == 1000000 {
			shouldFinish <- true
		}
	}
}
