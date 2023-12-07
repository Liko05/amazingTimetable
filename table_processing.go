package main

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

type ProcessingQueue struct {
	mu                  sync.Mutex
	queue               []Table
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

func (q *ProcessingQueue) AddToQueue(t Table) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, t)
}

func (q *ProcessingQueue) GetFromQueue() Table {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.queue) == 0 {
		log.Debug("Queue is empty")
		return Table{}
	}
	t := q.queue[0]
	q.queue = q.queue[1:]
	return t
}
