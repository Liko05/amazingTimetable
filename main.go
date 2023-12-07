package main

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	initLogger(log.DebugLevel)

	hashes := threadSafeListOfHashes{hashes: make(map[string]bool)}
	var shouldFinish = make(chan bool)
	var timeTables = make(chan Table)

	timeStart := time.Now()

	for i := 0; i < 3; i++ {
		go generateCalendars(&hashes, shouldFinish, timeTables)
	}

	<-shouldFinish
	log.Info("Time elapsed for generating  options: " + time.Since(timeStart).String())
}
