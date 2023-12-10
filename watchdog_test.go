package main

import (
	"testing"
	"time"
)

func TestWatchdog_Start(t *testing.T) {
	shouldFinish := make(chan bool)
	watchdog := Watchdog{
		DesiredDuration:             3,
		DelayBetweenProgressUpdates: 100,
		ShouldFinish:                shouldFinish,
	}
	timeStart := time.Now()
	watchdog.Start(timeStart)
	<-shouldFinish
	timeTotal := time.Since(timeStart).Seconds()

	if timeTotal < 3 {
		t.Errorf("Expected time to be greater than 3, got %f", timeTotal)
	}
}
