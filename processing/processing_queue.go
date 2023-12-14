package processing

import (
	"amazingTimetable/counter"
	"sync"
)

// Queue is a thread safe Queue for storing time tables
type Queue struct {
	Mu                 sync.Mutex
	ThreadSafeCounters *counter.ThreadSafeCounters
	Queue              []interface{}
	Hashes             map[uint32]bool
}

// Push pushes an element to the queue
func (p *Queue) Push(element interface{}) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.Queue = append(p.Queue, element)
}

// Pop pops an element from the queue
func (p *Queue) Pop() interface{} {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	if len(p.Queue) == 0 {
		return nil
	}

	element := p.Queue[0]
	p.Queue = p.Queue[1:]

	return element
}

func (p *Queue) CheckHash(hash uint32) bool {
	p.Mu.Lock()
	defer p.Mu.Unlock()

	if p.Hashes[hash] {
		return true
	}
	p.Hashes[hash] = true
	return false
}
