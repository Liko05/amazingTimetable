package generator

import (
	"amazingTimetable/counter"
	"amazingTimetable/processing"
	"amazingTimetable/table"
	"amazingTimetable/watchdog"
	"sync"
	"testing"
	"time"
)

// Watch out for this test, it is very resource intensive
func TestGenerator_GenerationWorkerStart(t *testing.T) {
	shouldFinish := make(chan bool)
	counters := counter.ThreadSafeCounters{}
	processingQueue := processing.ProcessingQueue{
		Mu:         sync.Mutex{},
		Queue:      make([]interface{}, 0),
		BestTable:  table.Table{},
		BestTables: make([]table.Table, 0),
	}
	watchdog := watchdog.Watchdog{
		DesiredDuration:             5,
		DelayBetweenProgressUpdates: 100,
		ShouldFinish:                shouldFinish,
		Counters:                    &counters,
	}
	g := Generator{
		Counters:        &counters,
		NumberOfWorkers: 1,
		ProcessingQueue: &processingQueue,
	}

	go watchdog.Start(time.Now())
	go g.Start()
	<-shouldFinish

	if counters.GetGenerated() <= 0 {
		t.Errorf("Expected generated to be more, got %d", counters.GetGenerated())
	}
}
