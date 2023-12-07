package main

import (
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"time"
)

type Graders struct {
	NumberOfWorkers int
	ShouldFinish    chan bool
	Counters        *ThreadSafeCounters
	ProcessingQueue *ProcessingQueue
}

func (g *Graders) Start() {
	for i := 0; i < g.NumberOfWorkers; i++ {
		go g.GradeTimeTablesStartWorker()
	}
}

func (g *Graders) GradeTimeTablesStartWorker() {
	for {
		log.Debug("Checking new time table")
		t := g.ProcessingQueue.GetFromQueue()
		if t.isEmpty() {
			log.Debug("No more time tables to check")
			time.Sleep(1 * time.Millisecond)
			continue
		}
		log.Debug("Checking time table: " + t.prettyPrint())
		score := t.getScore()
		if score > g.ProcessingQueue.GetHighestScore() {
			g.ProcessingQueue.SetHighestScore(score)
			g.ProcessingQueue.bestTable = t
			log.Debug("New highest score: " + strconv.Itoa(score))
		}
		g.Counters.incrementChecked()
	}
}

func (tb *Table) getScore() int {
	score := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			if tb.TimeTable[i][j].IsPractical {
				score += rand.Intn(250)
			} else {
				score += rand.Intn(100)
			}
		}
	}
	return score
}
