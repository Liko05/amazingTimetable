package generator

import (
	"amazingTimetable/counter"
	"amazingTimetable/hash"
	"amazingTimetable/table"
)

// Generator is responsible for running the generation of new timetables
type Generator struct {
	Counters     *counter.ThreadSafeCounters
	Hashes       *hash.Hashes
	ShouldFinish chan bool
}

// Start starts the generation of new timetables based on the NumberOfWorkers
func (g *Generator) Start(que chan table.Table) {
	defaultTimeTable := table.Table{}
	defaultTimeTable.CreateDefault()
	for {
		defaultTimeTable.Shuffle()
		//if !g.Hashes.CheckHash(defaultTimeTable.Hash()) {
		que <- defaultTimeTable
		g.Counters.IncrementGenerated()

		//	}
	}
}
