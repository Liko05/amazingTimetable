package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"sync"
	"time"
)

func main() {
	initLogger(log.InfoLevel)

	//thread shared variables
	var shouldFinish = make(chan bool)
	var counters = ThreadSafeCounters{}
	var processingQueue = ProcessingQueue{
		mu:    sync.Mutex{},
		queue: make([]interface{}, 0),
		bestTable: Table{
			Score: -10000,
		},
	}

	// worker variables
	var watchdog = Watchdog{
		DesiredDuration:             180,
		DelayBetweenProgressUpdates: 10,
		ShouldFinish:                shouldFinish,
	}
	var generators = Generators{
		Counters:        &counters,
		NumberOfWorkers: 7,
		ProcessingQueue: &processingQueue,
	}
	var graders = Graders{
		NumberOfWorkers: 12,
		Counters:        &counters,
		ProcessingQueue: &processingQueue,
	}

	log.Info("Starting to generate time tables")
	timeStart := time.Now()

	watchdog.Start(timeStart, &counters)
	generators.Start()
	graders.Start()

	<-shouldFinish
	log.Info("Time elapsed for " + strconv.FormatUint(counters.getGenerated(), 10) + " options: " + time.Since(timeStart).String() + " time tables generated and " + strconv.FormatUint(counters.getChecked(), 10) + " time tables checked")

	//wait for input to exit
	log.Info("Finished \n \n \n \n ")
	log.Info("Best table: \n" + processingQueue.bestTable.String())
	log.Info("Score: " + strconv.Itoa(processingQueue.bestTable.Score))
	log.Info(counters.getValid())
	var input string
	log.Info("Press enter to exit")
	_, _ = fmt.Scanln(&input)
}
