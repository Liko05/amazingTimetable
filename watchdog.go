package main

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type Watchdog struct {
	DesiredDuration             int // seconds
	DelayBetweenProgressUpdates int // seconds
}

func (w *Watchdog) Start(shouldFinish chan bool, timeStart time.Time, counters *ThreadSafeCounters) {
	go func() {
		lastUpdate := time.Now()
		for {
			if time.Since(timeStart).Seconds() >= float64(w.DesiredDuration) {
				shouldFinish <- true
			} else if time.Since(lastUpdate).Seconds() >= float64(w.DelayBetweenProgressUpdates) {
				time.Sleep(time.Duration(w.DelayBetweenProgressUpdates) * time.Second)
				log.Info("Generated time tables: " + strconv.FormatUint(counters.getGenerated(), 10) + " Checked time tables: " + strconv.FormatUint(counters.getChecked(), 10))
				lastUpdate = time.Now()
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
}
