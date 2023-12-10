package main

import (
	"sync"
)

// ProcessingQueue is a thread safe Queue for storing time tables
type ProcessingQueue struct {
	Mu         sync.Mutex
	Queue      []interface{}
	BestTable  Table
	BestTables []Table // left here just for debugging purposes
}

// Push pushes an element to the queue
func (q *ProcessingQueue) Push(element interface{}) {
	q.Mu.Lock()
	defer q.Mu.Unlock()

	q.Queue = append(q.Queue, element)
}

// Pop pops an element from the queue
func (q *ProcessingQueue) Pop() interface{} {
	q.Mu.Lock()
	defer q.Mu.Unlock()
	if len(q.Queue) == 0 {
		return nil
	}

	element := q.Queue[0]
	q.Queue = q.Queue[1:]

	return element
}

// AddIfBetter adds an element to the BestTable if it is better than the current BestTable
func (q *ProcessingQueue) AddIfBetter(element interface{}) {
	q.Mu.Lock()
	defer q.Mu.Unlock()

	table, ok := element.(Table)
	if ok == false {
		return
	}

	if table.Score > q.BestTable.Score {
		q.BestTable = table
	}
}

// AddToBestTables adds an element to the BestTables
// left here just for debugging purposes
func (q *ProcessingQueue) AddToBestTables(element interface{}) {
	q.Mu.Lock()
	defer q.Mu.Unlock()

	table, ok := element.(Table)
	if ok == false {
		return
	}

	q.BestTables = append(q.BestTables, table)
}
