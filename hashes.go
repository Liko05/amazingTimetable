package main

import (
	"sync"
)

type ThreadSafeListOfHashes struct {
	mu     sync.Mutex
	hashes map[string]bool
}

func (l *ThreadSafeListOfHashes) add(hash string) {
	l.mu.Lock()
	l.hashes[hash] = true
	l.mu.Unlock()
}

func (l *ThreadSafeListOfHashes) contains(hash string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.hashes[hash]
}
