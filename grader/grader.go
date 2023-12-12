package grader

import (
	"amazingTimetable/counter"
	"amazingTimetable/processing"
	"amazingTimetable/table"
)

// Grader is a struct which holds all the information about the grader
type Grader struct {
	Counters        *counter.ThreadSafeCounters
	ProcessingQueue *processing.ProcessingQueue
}

// GradeTimeTable Calls validation and grading functions on the table
func (g *Grader) GradeTimeTable(table table.Table) {

}
