// Package generator is responsible for running the generation of new timetables
package generator

import (
	"amazingTimetable/counter"
	"amazingTimetable/hash"
	"amazingTimetable/table"
)

// Generator is struct that is holding the necessary data for the generation of new timetables
type Generator struct {
	Counters     *counter.ThreadSafeCounters
	Hashes       *hash.Hashes
	ShouldFinish chan bool
}

// Start starts the generation of new timetables and sends them to the channel for processing if their hash is not already in the hash map
func (g *Generator) Start(que chan table.Table) {
	defaultTimeTable := table.Table{}
	defaultTimeTable.CreateDefault()
	for {
		if g.Counters.StopGeneration {
			close(que)
			return
		}
		defaultTimeTable.Shuffle()
		if !g.Hashes.ContainsAndAdd(defaultTimeTable.Hash()) {
			que <- defaultTimeTable
			g.Counters.IncrementGenerated()
		}
	}
}
