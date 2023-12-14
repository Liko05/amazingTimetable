package grader

import (
	"amazingTimetable/counter"
	"amazingTimetable/table"
)

type Grader struct {
	Counters     *counter.ThreadSafeCounters
	ShouldFinish chan bool
}

func (g *Grader) Start(queue chan table.Table) {
	for {
		select {
		case t := <-queue:
			if t.IsTableValid() {
				g.Counters.IncrementValid()
				t.GradeTable()
				t.Score = 500
				g.Counters.SetBestOption(t)
			}
			g.Counters.IncrementChecked()
		}
	}
}
