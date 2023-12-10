package main

import (
	log "github.com/sirupsen/logrus"
)

// Generator is responsible for running the generation of new timetables
type Generator struct {
	Counters        *ThreadSafeCounters
	NumberOfWorkers int
	ProcessingQueue *ProcessingQueue
}

// Start starts the generation of new timetables based on the NumberOfWorkers
func (g *Generator) Start() {
	for i := 0; i < g.NumberOfWorkers; i++ {
		go g.GenerationWorkerStart()
	}
}

// GenerationWorkerStart is the worker that generates new timetables
func (g *Generator) GenerationWorkerStart() {
	defaultTimeTable := Table{}
	defaultTimeTable.CreateDefault()
	for {
		log.Debug("Generating new time table")
		defaultTimeTable.Shuffle()
		g.ProcessingQueue.Push(defaultTimeTable)
		g.Counters.IncrementGenerated()
		log.Debug("Generating again")
	}
}
