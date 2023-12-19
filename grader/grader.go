// Package grader contains the Grader struct and its methods used for processing the tables and grading them.
package grader

import (
	"amazingTimetable/counter"
	"amazingTimetable/table"
)

// Grader is a struct that contains the counters and the channel for finishing the program
type Grader struct {
	Counters     *counter.ThreadSafeCounters
	ShouldFinish chan bool
}

// Start starts the grader worker that is processing incoming tables
func (g *Grader) Start(queue chan table.Table) {
	for {
		select {
		case t, ok := <-queue:
			if !ok {
				return
			}
			if t.IsTableValid() {
				g.Counters.IncrementValid()
				t.GradeTable()
				g.Counters.SetBestOption(t)
			}
			g.Counters.IncrementChecked()
		}
	}
}
