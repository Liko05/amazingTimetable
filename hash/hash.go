// Package hash is creating thread safe map for the purpose of storing hashes.
package hash

import (
	"sync"
)

// Hashes is thread safe and contains map of used hash
type Hashes struct {
	mu     sync.Mutex
	Hashes map[uint32]bool
}

// ContainsAndAdd checks if hash is already in map and adds it if it's not
func (h *Hashes) ContainsAndAdd(hash uint32) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.Hashes[hash] {
		return true
	}
	h.Hashes[hash] = true
	return false
}
