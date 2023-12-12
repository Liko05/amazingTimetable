package main

// Grader is a struct which holds all the information about the grader
type Grader struct {
	Counters        *ThreadSafeCounters
	ProcessingQueue *ProcessingQueue
}

func (g *Grader) GradeTimeTable(table Table) {

}
