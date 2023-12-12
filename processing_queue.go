package main

import (
	"sync"
)

// ProcessingQueue is a thread safe Queue for storing time tables
type ProcessingQueue struct {
	Mu                 sync.Mutex
	BestTable          Table
	OriginalTable      Table
	ThreadSafeCounters *ThreadSafeCounters
	Hashes             map[uint32]bool
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

func (q *ProcessingQueue) AddHash(hash uint32) {
	q.Mu.Lock()
	defer q.Mu.Unlock()

	q.Hashes[hash] = true
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
