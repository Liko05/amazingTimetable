package main

import (
	c "amazingTimetable/counter"
	g "amazingTimetable/generator"
	p "amazingTimetable/processing"
	t "amazingTimetable/table"
	w "amazingTimetable/watchdog"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// CreateVariablesForWorkers creates the variables for the workers
func CreateVariablesForWorkers() (chan bool, *c.ThreadSafeCounters, *p.ProcessingQueue) {
	var shouldFinish = make(chan bool)
	var counters = c.ThreadSafeCounters{
		Mu:                       sync.Mutex{},
		GeneratedOptions:         0,
		CheckedOptions:           0,
		ValidOptions:             0,
		OptionsBetterThanDefault: 0,
	}

	var processingQueue = p.ProcessingQueue{
		Mu:                 sync.Mutex{},
		BestTable:          t.Table{},
		OriginalTable:      t.Table{},
		ThreadSafeCounters: nil,
		Hashes:             make(map[uint32]bool, 200_000_000),
	}

	return shouldFinish, &counters, &processingQueue
}

// CreateWorkers creates the workers
func CreateWorkers(shouldFinish chan bool, counters *c.ThreadSafeCounters, processingQueue *p.ProcessingQueue) (w.Watchdog, g.Generator) {
	var watchdog = w.Watchdog{
		DesiredDuration:             0,
		DelayBetweenProgressUpdates: 0,
		ShouldFinish:                shouldFinish,
		Counters:                    counters,
	}

	var generator = g.Generator{
		ShouldFinish:    shouldFinish,
		Counters:        counters,
		ProcessingQueue: processingQueue,
	}

	return watchdog, generator
}

// GetArgsAndApply gets the arguments from the command line and applies them to the workers
func GetArgsAndApply(watchdog *w.Watchdog, generator *g.Generator) {
	var timeLimit int
	var timeBetweenProgressUpdates int
	var numberOfGenerators int
	var debugLevel bool
	var help bool

	numberOfAvailableCPUs := runtime.NumCPU()

	flag.IntVar(&timeLimit, "t", 180, "The time limit in seconds")
	flag.IntVar(&timeBetweenProgressUpdates, "p", 10, "The time between progress updates in seconds")
	flag.IntVar(&numberOfGenerators, "g", numberOfAvailableCPUs-1, "The number of generators")
	flag.BoolVar(&debugLevel, "d", false, "Enable debug level logging")
	flag.BoolVar(&help, "h", false, "Show help")

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if debugLevel {
		InitLogger(log.DebugLevel)
		log.Debug("Time limit: " + strconv.Itoa(timeLimit) + " seconds")
		log.Debug("Time between progress updates: " + strconv.Itoa(timeBetweenProgressUpdates) + " seconds")
		log.Debug("Number of generators: " + strconv.Itoa(numberOfGenerators))
	} else {
		InitLogger(log.InfoLevel)
	}
	log.Info("Number of available CPUs: " + strconv.Itoa(numberOfAvailableCPUs))

	watchdog.DesiredDuration = timeLimit
	watchdog.DelayBetweenProgressUpdates = timeBetweenProgressUpdates
	generator.NumberOfWorkers = numberOfGenerators

	log.Info("Starting with time limit: " + strconv.Itoa(timeLimit) + " seconds")
}

// main is the entry point of the program
func main() {
	var watchdog, generator = CreateWorkers(CreateVariablesForWorkers())
	GetArgsAndApply(&watchdog, &generator)

	watchdog.Start(time.Now())
	generator.Start()

	<-watchdog.ShouldFinish

	log.Info("Finished")
	log.Info("Generated options: " + strconv.FormatUint(generator.Counters.GeneratedOptions, 10))
}
