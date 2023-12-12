package main

import (
	"sync"
)

// ProcessingQueue is a thread safe Queue for storing time tables
type ProcessingQueue struct {
	Mu                 sync.Mutex
	Queue              []interface{}
	BestTable          Table
	OriginalTable      Table
	BestTables         []Table
	ThreadSafeCounters *ThreadSafeCounters
	Hashes             map[Table]bool
}

// Push pushes an element to the queue
func (q *ProcessingQueue) Push(element interface{}) {
	q.Mu.Lock()
	defer q.Mu.Unlock()
	table, ok := element.(Table)
	if ok == false {
		return
	}

	q.Hashes[table] = true

	q.Queue = append(q.Queue, table)
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
		q.ThreadSafeCounters.IncrementOptionsBetterThanDefault()
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

	if q.CheckBestTableUnique(&table) == false {
		return
	}

	q.BestTables = append(q.BestTables, table)

}

func (q *ProcessingQueue) AddOriginal(element interface{}) {
	q.Mu.Lock()
	defer q.Mu.Unlock()

	table, ok := element.(Table)
	if ok == false {
		return
	}

	q.OriginalTable = table
}

func (q *ProcessingQueue) CheckBestTableUnique(table *Table) bool {
	for i := 0; i < len(q.BestTables); i++ {
		if q.BestTables[i].Hash() == table.Hash() {
			return false
		}
	}
	return true
}
