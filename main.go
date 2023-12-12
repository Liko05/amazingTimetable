package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// CreateVariablesForWorkers creates the variables that will be used and shared by the workers
func CreateVariablesForWorkers() (chan bool, *ThreadSafeCounters, *ProcessingQueue) {
	var shouldFinish = make(chan bool)
	var counters = ThreadSafeCounters{
		Mu:               sync.Mutex{},
		GeneratedOptions: 0,
		CheckedOptions:   0,
		ValidOptions:     0,
	}
	var processingQueue = ProcessingQueue{
		Mu:    sync.Mutex{},
		Queue: make([]interface{}, 0),
		BestTable: Table{
			Score: -10000,
		},
		BestTables:   make([]Table, 0),
		ShouldFinish: &shouldFinish,
	}
	return shouldFinish, &counters, &processingQueue
}

// CreateWorkers creates the workers that will be used to generate and grade the timetables and the watchdog that will monitor the progress
func CreateWorkers(shouldFinish chan bool, counters *ThreadSafeCounters, processingQueue *ProcessingQueue) (Watchdog, Generator, Grader) {
	var watchdog = Watchdog{
		ShouldFinish: shouldFinish,
		Counters:     counters,
	}
	var generators = Generator{
		Counters:        counters,
		ProcessingQueue: processingQueue,
	}
	var graders = Grader{
		Counters:        counters,
		ProcessingQueue: processingQueue,
	}
	return watchdog, generators, graders
}

// GetArgsAndApplyArgs gets the arguments from the command line and applies them to the workers
func GetArgsAndApplyArgs(watchdog *Watchdog, generators *Generator, graders *Grader) {
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
		log.Debug("Number of graders: " + strconv.Itoa(numberOfGraders))
	} else {
		InitLogger(log.InfoLevel)
	}
	log.Info("Number of available CPUs: " + strconv.Itoa(numberOfAvailableCPUs))

	watchdog.DesiredDuration = timeLimit
	watchdog.DelayBetweenProgressUpdates = timeBetweenProgressUpdates

	generators.NumberOfWorkers = numberOfGenerators
	graders.NumberOfWorkers = numberOfGraders

	log.Info("Starting with time limit: " + strconv.Itoa(timeLimit) + " seconds")
}

// main is the entry point of the program
func main() {
	shouldFinish, counters, processingQueue := CreateVariablesForWorkers()
	watchdog, generators, graders := CreateWorkers(shouldFinish, counters, processingQueue)

	GetArgsAndApplyArgs(&watchdog, &generators, &graders)

	timeStart := time.Now()

	watchdog.Start(timeStart)
	generators.Start()
	graders.Start()

	<-shouldFinish
	log.Info("Time limit reached, finished processing at: " + time.Now().Format("2006-01-02 15:04:05"))
	log.Info("Total time taken: " + time.Since(timeStart).String())
	log.Info("Generated options: " + strconv.FormatUint(counters.GeneratedOptions, 10) + ", checked options: " + strconv.FormatUint(counters.CheckedOptions, 10) + ", valid options: " + strconv.FormatUint(counters.ValidOptions, 10))
	log.Info("Best table has score: " + strconv.Itoa(processingQueue.BestTable.Score))
	log.Info(strconv.Itoa(len(processingQueue.BestTables)) + "were better than the original table with a score of: " + strconv.Itoa(processingQueue.OriginalTable.Score))
	log.Info("Best table: ")
	println(processingQueue.BestTable.String())

	if processingQueue.OriginalTable.Hash() == processingQueue.BestTable.Hash() {
		log.Info("Best table is the original table")
	} else {
		log.Info("Original table: ")
		println(processingQueue.OriginalTable.String())
		log.Info("Original table score: " + strconv.Itoa(processingQueue.OriginalTable.Score))
	}
}
