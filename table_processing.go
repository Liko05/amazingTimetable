package main

import (
	"sync"
)

type ProcessingQueue struct {
	mu        sync.Mutex
	queue     []interface{}
	bestTable Table
}

func (q *ProcessingQueue) Push(element interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.queue = append(q.queue, element)
}

func (q *ProcessingQueue) Pop() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.queue) == 0 {
		return nil
	}

	element := q.queue[0]
	q.queue = q.queue[1:]

	return element
}

func (q *ProcessingQueue) AddIfBetter(element interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()

	table, ok := element.(Table)
	if ok == false {
		return
	}

	if table.Score > q.bestTable.Score {
		q.bestTable = table
	}
}
