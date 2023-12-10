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
	var counters = ThreadSafeCounters{
		mu:               sync.Mutex{},
		generatedOptions: 0,
		checkedOptions:   0,
		validOptions:     0,
	}
	var processingQueue = ProcessingQueue{
		mu:    sync.Mutex{},
		queue: make([]interface{}, 0),
		bestTable: Table{
			Score: -10000,
		},
		bestTables: make([]Table, 0),
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
	generators.start()
	graders.start()

	<-shouldFinish
	log.Info("Time elapsed for " + strconv.FormatUint(counters.getGenerated(), 10) + " options: " + time.Since(timeStart).String() + " time tables generated and " + strconv.FormatUint(counters.getChecked(), 10) + " time tables checked")
	log.Info("Best table:" + "\n")
	log.Info(processingQueue.bestTable.String())
	var input string
	log.Info("Press enter to exit")
	_, _ = fmt.Scanln(&input)
}
