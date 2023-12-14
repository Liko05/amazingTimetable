package main

import (
	c "amazingTimetable/counter"
	g "amazingTimetable/generator"
	gr "amazingTimetable/grader"
	p "amazingTimetable/processing"
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
func CreateVariablesForWorkers() (chan bool, *c.ThreadSafeCounters, *p.Queue) {
	var shouldFinish = make(chan bool)
	var counters = c.ThreadSafeCounters{
		Mu:                       sync.Mutex{},
		GeneratedOptions:         0,
		CheckedOptions:           0,
		ValidOptions:             0,
		OptionsBetterThanDefault: 0,
	}

	var processingQueue = p.Queue{
		Mu:                 sync.Mutex{},
		ThreadSafeCounters: &counters,
		Queue:              make([]interface{}, 0),
		Hashes:             make(map[uint32]bool, 200_000_000),
	}

	return shouldFinish, &counters, &processingQueue
}

// CreateWorkers creates the workers
func CreateWorkers(shouldFinish chan bool, counters *c.ThreadSafeCounters, processingQueue *p.Queue) (w.Watchdog, g.Generator, gr.Grader) {
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

	var grader = gr.Grader{
		Counters:        counters,
		ProcessingQueue: processingQueue,
	}

	return watchdog, generator, grader
}

// GetArgsAndApply gets the arguments from the command line and applies them to the workers
func GetArgsAndApply(watchdog *w.Watchdog, generator *g.Generator, grader *gr.Grader) {
	var timeLimit int
	var timeBetweenProgressUpdates int
	var numberOfGenerators int
	var numberOfGraders int
	var debugLevel bool
	var help bool

	numberOfAvailableCPUs := runtime.NumCPU()

	flag.IntVar(&timeLimit, "t", 180, "The time limit in seconds")
	flag.IntVar(&timeBetweenProgressUpdates, "p", 10, "The time between progress updates in seconds")
	flag.IntVar(&numberOfGenerators, "g", numberOfAvailableCPUs/2, "The number of generators")
	flag.IntVar(&numberOfGraders, "r", numberOfAvailableCPUs/2, "The number of graders")
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
	grader.NumberOfWorker = numberOfGraders

	log.Info("Starting with time limit: " + strconv.Itoa(timeLimit) + " seconds")
}

// main is the entry point of the program
func main() {
	var watchdog, generator, grader = CreateWorkers(CreateVariablesForWorkers())
	GetArgsAndApply(&watchdog, &generator, &grader)

	watchdog.Start(time.Now())
	generator.Start()
	grader.Start()

	<-watchdog.ShouldFinish

	log.Info("Finished")
	log.Info("Checked options: " + strconv.FormatUint(grader.Counters.CheckedOptions, 10))
	log.Info("Generated options: " + strconv.FormatUint(generator.Counters.GeneratedOptions, 10))
	log.Info("Valid options: " + strconv.FormatUint(grader.Counters.ValidOptions, 10))
	log.Info("Options with unique hash: " + strconv.FormatUint(grader.Counters.OptionsBetterThanDefault, 10))
}
