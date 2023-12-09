package main

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

func main() {
	initLogger(log.InfoLevel)

	//thread shared variables
	hashes := ThreadSafeListOfHashes{hashes: make(map[string]bool)}
	var shouldFinish = make(chan bool)
	var counters = ThreadSafeCounters{}
	var queue = make(chan Table, 1000000)
	var processingQueue = ProcessingQueue{
		mu:                  sync.Mutex{},
		queue:               queue,
		currentHighestScore: 0,
		bestTable:           Table{},
	}

	// worker variables
	var watchdog = Watchdog{DesiredDuration: 180, DelayBetweenProgressUpdates: 10}
	var generators = Generators{
		Hashes:          &hashes,
		ShouldFinish:    shouldFinish,
		Counters:        &counters,
		NumberOfWorkers: 6,
		ProcessingQueue: &processingQueue,
	}
	var graders = Graders{
		NumberOfWorkers: 3,
		ShouldFinish:    shouldFinish,
		Counters:        &counters,
		ProcessingQueue: &processingQueue,
	}

	log.Info("Starting to generate time tables")
	timeStart := time.Now()

	watchdog.Start(shouldFinish, timeStart, &counters)
	generators.Start()
	graders.Start()

	<-shouldFinish
	log.Info("Time elapsed for options: " + time.Since(timeStart).String() + " seconds " + strconv.FormatUint(counters.getGenerated(), 10) + " time tables generated and " + strconv.FormatUint(counters.getChecked(), 10) + " time tables checked")
	log.Info("Best time table: " + processingQueue.bestTable.prettyPrint() + " with score: " + strconv.Itoa(processingQueue.GetHighestScore()))
}
