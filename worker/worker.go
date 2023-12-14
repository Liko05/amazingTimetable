// Package worker takes care of creating workers and starting them in a new go routine as a pair of generator and grader.
package worker

import (
	"amazingTimetable/counter"
	"amazingTimetable/generator"
	"amazingTimetable/grader"
	"amazingTimetable/hash"
	"amazingTimetable/table"
	"amazingTimetable/watchdog"
	log "github.com/sirupsen/logrus"
	"time"
)

// Workers is a struct that contains the necessary variables for the workers
type Workers struct {
	Counters        *counter.ThreadSafeCounters
	Hashes          *hash.Hashes
	ShouldFinish    chan bool
	NumberOfWorkers int
	TimeBetweenLogs int
	TimeLimit       int
}

// initWorkers initializes the worker with the necessary variables and returns them
func (w *Workers) initWorkers() (generator.Generator, grader.Grader, watchdog.Watchdog) {
	gen := generator.Generator{
		Counters:     w.Counters,
		Hashes:       w.Hashes,
		ShouldFinish: w.ShouldFinish,
	}

	grade := grader.Grader{
		Counters:     w.Counters,
		ShouldFinish: w.ShouldFinish,
	}

	watch := watchdog.Watchdog{
		DesiredDuration:             w.TimeLimit,
		DelayBetweenProgressUpdates: w.TimeBetweenLogs,
		ShouldFinish:                w.ShouldFinish,
		Counters:                    w.Counters,
	}

	return gen, grade, watch
}

// Start starts the workers as a pair of generator and grader with channel to communicate between them
func (w *Workers) Start() {
	gen, grade, watch := w.initWorkers()
	log.Info("Starting workers")
	go watch.Start(time.Now())
	for i := 0; i < w.NumberOfWorkers; i++ {
		channelForWorkers := make(chan table.Table, 10000)
		go gen.Start(channelForWorkers)
		go grade.Start(channelForWorkers)
	}
}
