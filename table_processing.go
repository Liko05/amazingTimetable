package main

import (
	"sync"
)

type ProcessingQueue struct {
	mu                  sync.Mutex
	queue               chan Table
	currentHighestScore int
	bestTable           Table
}

func (q *ProcessingQueue) SetHighestScore(score int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.currentHighestScore = score
}

func (q *ProcessingQueue) GetHighestScore() int {
	q.mu.Lock()
	defer q.mu.Unlock()
	return q.currentHighestScore
}
