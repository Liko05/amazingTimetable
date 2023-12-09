package main

import "sync"

type ThreadSafeCounters struct {
	mu               sync.Mutex
	generatedOptions uint64
	checkedOptions   uint64
	validOptions     uint64
}

func (l *ThreadSafeCounters) incrementGenerated() {
	l.mu.Lock()
	l.generatedOptions++
	l.mu.Unlock()
}

func (l *ThreadSafeCounters) incrementChecked() {
	l.mu.Lock()
	l.checkedOptions++
	l.mu.Unlock()
}

func (l *ThreadSafeCounters) getGenerated() uint64 {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.generatedOptions
}

func (l *ThreadSafeCounters) getChecked() uint64 {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.checkedOptions
}

func (l *ThreadSafeCounters) incrementValid() {
	l.mu.Lock()
	l.validOptions++
	l.mu.Unlock()
}

func (l *ThreadSafeCounters) getValid() uint64 {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.validOptions
}
