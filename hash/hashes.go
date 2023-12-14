package hash

import (
	"sync"
)

// Hashes is thread safe and contains map of used hash
type Hashes struct {
	mu     sync.Mutex
	Hashes map[uint32]bool
}

func (h *Hashes) CheckHash(hash uint32) bool {
	h.mu.Lock()
	defer h.mu.Unlock()

	if h.Hashes[hash] {
		return true
	}
	h.Hashes[hash] = true
	return false
}
