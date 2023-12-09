package main

import (
	log "github.com/sirupsen/logrus"
)

type Generators struct {
	Hashes          *ThreadSafeListOfHashes
	ShouldFinish    chan bool
	Counters        *ThreadSafeCounters
	NumberOfWorkers int
	ProcessingQueue *ProcessingQueue
}

func (g *Generators) Start() {
	for i := 0; i < g.NumberOfWorkers; i++ {
		go g.GenerateTimeTablesStartWorker()
	}
}

func (g *Generators) GenerateTimeTablesStartWorker() {
	defaultTimeTable := Table{}
	defaultTimeTable.createDefault()
	g.ProcessingQueue.queue <- defaultTimeTable
	g.Counters.incrementGenerated() //4860
	for {
		defaultTimeTable = defaultTimeTable.generateNewTimeTable(defaultTimeTable.retrieveArrayOfSubjectsWithoutPauses())
		log.Debug(defaultTimeTable.prettyPrint())
		//log.Debug(g.Hashes.hashes)
		//if !defaultTimeTable.checkIfHashAlreadyExists(g.Hashes) {
		g.Counters.incrementGenerated()
		g.ProcessingQueue.queue <- defaultTimeTable
		//	log.Debug("Sent new time table for checking")
		//	g.ProcessingQueue.AddToQueue(defaultTimeTable)
		//} else {
		//	log.Debug("Found duplicate hash")
		//	g.ShouldFinish <- true
		//}
		log.Debug("Generating again")
	}
}
