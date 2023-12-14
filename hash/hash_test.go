package hash

import (
	"sync"
	"testing"
)

func Test_HashContains(t *testing.T) {
	hashes := Hashes{
		mu:     sync.Mutex{},
		Hashes: make(map[uint32]bool),
	}

	hashes.Hashes[1] = true

	if !hashes.ContainsAndAdd(1) {
		t.Errorf("Expected true, got false")
	}

	if hashes.ContainsAndAdd(2) {
		t.Errorf("Expected false, got true")
	}
}
