package main

import (
	"sync"
)

type threadSafeListOfHashes struct {
	mu     sync.Mutex
	hashes map[string]bool
}

func (l *threadSafeListOfHashes) add(hash string) {
	l.mu.Lock()
	l.hashes[hash] = true
	l.mu.Unlock()
}

func (l *threadSafeListOfHashes) contains(hash string) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.hashes[hash]
}
