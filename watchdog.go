package main

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Watchdog struct {
	DesiredDuration             int // seconds
	DelayBetweenProgressUpdates int // seconds
	ShouldFinish                chan bool
}

func (w *Watchdog) Start(timeStart time.Time, counters *ThreadSafeCounters) {
	go func() {
		lastUpdate := time.Now()
		for {
			if time.Since(timeStart).Seconds() >= float64(w.DesiredDuration) {
				w.ShouldFinish <- true
			} else if time.Since(lastUpdate).Seconds() >= float64(w.DelayBetweenProgressUpdates) {
				log.Info("Generated time tables: " + strconv.FormatUint(counters.getGenerated(), 10) + " Checked time tables: " + strconv.FormatUint(counters.getChecked(), 10))
				lastUpdate = time.Now()
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
}
