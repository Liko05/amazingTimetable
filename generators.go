package main

import (
	log "github.com/sirupsen/logrus"
)

type Generators struct {
	Counters        *ThreadSafeCounters
	NumberOfWorkers int
	ProcessingQueue *ProcessingQueue
}

func (g *Generators) Start() {
	initTable := Table{}
	initTable.createDefault()
	g.ProcessingQueue.Push(initTable)
	for i := 0; i < g.NumberOfWorkers; i++ {
		go g.GenerateTimeTablesStartWorker()
	}
}

func (g *Generators) GenerateTimeTablesStartWorker() {
	defaultTimeTable := Table{}
	defaultTimeTable.createDefault()
	for {
		log.Debug("Generating new time table")
		defaultTimeTable.shuffle()
		g.ProcessingQueue.Push(defaultTimeTable)
		g.Counters.incrementGenerated()
		log.Debug("Generating again")
	}
}
