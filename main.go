package main

import (
	c "amazingTimetable/counter"
	h "amazingTimetable/hash"
	"amazingTimetable/table"
	"amazingTimetable/utils"
	"amazingTimetable/worker"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
	"sync"
)

// CreateVariablesForWorkers creates the variables for the workers
func CreateVariablesForWorkers() (chan bool, *c.ThreadSafeCounters, *h.Hashes) {
	var shouldFinish = make(chan bool)
	var counters = c.ThreadSafeCounters{
		Mu:                       sync.Mutex{},
		GeneratedOptions:         0,
		CheckedOptions:           0,
		ValidOptions:             0,
		OptionsBetterThanDefault: 0,
		BestOption:               table.Table{},
		OriginalOption:           table.Table{},
	}

	counters.OriginalOption.CreateDefault()
	counters.OriginalOption.GradeTable()

	var hashes = h.Hashes{
		Hashes: make(map[uint32]bool, 200_000_000), // explanation on this is in README.md
	}

	return shouldFinish, &counters, &hashes
}

// GetArgs gets the arguments from the command line and returns the necessary variables
func GetArgs() (int, int, int) {
	var timeLimit int
	var timeBetweenProgressUpdates int
	var numberOfWorkers int
	var debugLevel bool
	var help bool

	numberOfAvailableCPUs := runtime.NumCPU()

	flag.IntVar(&timeLimit, "t", 180, "The time limit in seconds")
	flag.IntVar(&timeBetweenProgressUpdates, "p", 10, "The time between progress updates in seconds")
	flag.IntVar(&numberOfWorkers, "w", numberOfAvailableCPUs/2, "Number of generator - grader pairs")
	flag.BoolVar(&debugLevel, "d", false, "Enable debug level logging")
	flag.BoolVar(&help, "h", false, "Show help")

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if debugLevel {
		InitLogger(log.DebugLevel)
	} else {
		InitLogger(log.InfoLevel)
	}
	return timeLimit, timeBetweenProgressUpdates, numberOfWorkers
}

// main is the entry point of the program
func main() {
	shouldFinish, counters, hashes := CreateVariablesForWorkers()
	timeLimit, timeBetweenLogs, numberOfWorkers := GetArgs()
	workers := worker.Workers{
		Counters:        counters,
		Hashes:          hashes,
		ShouldFinish:    shouldFinish,
		NumberOfWorkers: numberOfWorkers,
		TimeBetweenLogs: timeBetweenLogs,
		TimeLimit:       timeLimit,
	}

	log.Info("Starting... Time limit: " + strconv.Itoa(timeLimit) + " seconds, Time between logs: " + strconv.Itoa(timeBetweenLogs) + " seconds, Number of workers: " + strconv.Itoa(numberOfWorkers))
	workers.Start()

	<-shouldFinish
	log.Info("Finished execution")
	log.Info("Generated options: " + strconv.FormatUint(counters.GetGenerated(), 10) + " Checked options: " + strconv.FormatUint(counters.GetChecked(), 10))
	log.Info("Valid options: " + strconv.FormatUint(counters.GetValid(), 10))
	log.Info("Options better than default: " + strconv.FormatUint(counters.GetOptionsBetterThanDefault(), 10))
	if counters.GetOptionsBetterThanDefault() > 0 {
		log.Info("Best option: ")
		println(utils.TableToString(counters.GetBestOption()))
	} else {
		log.Info("No option better than default")
		log.Info("Default option: ")
		println(utils.TableToString(counters.GetOriginalOption()))
	}
}
