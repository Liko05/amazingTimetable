package main

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// Watchdog is a struct that represents a watchdog
type Watchdog struct {
	DesiredDuration             int // seconds
	DelayBetweenProgressUpdates int // seconds
	ShouldFinish                chan bool
	Counters                    *ThreadSafeCounters
}

// Start starts the watchdog
// It will send a message to the ShouldFinish channel when the time limit is reached
// It will also log the progress every DelayBetweenProgressUpdates seconds
func (w *Watchdog) Start(timeStart time.Time) {
	go func() {
		lastUpdate := time.Now()
		for {
			if time.Since(timeStart).Seconds() >= float64(w.DesiredDuration) {
				w.ShouldFinish <- true
				log.Info("Time limit reached")
				return
			} else if time.Since(lastUpdate).Seconds() >= float64(w.DelayBetweenProgressUpdates) {
				log.Info("Generated time tables: " + strconv.FormatUint(w.Counters.GetGenerated(), 10) + " Checked time tables: " + strconv.FormatUint(w.Counters.GetChecked(), 10))
				lastUpdate = time.Now()
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
}
