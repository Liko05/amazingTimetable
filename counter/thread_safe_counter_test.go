package counter

import (
	"testing"
)

func TestThreadSafeCounters_IncrementChecked(t *testing.T) {
	counter := ThreadSafeCounters{}
	counter.IncrementChecked()
	if counter.CheckedOptions != 1 {
		t.Error("Expected 1, got", counter.CheckedOptions)
	}
}

func TestThreadSafeCounters_IncrementGenerated(t *testing.T) {
	counter := ThreadSafeCounters{}
	counter.IncrementGenerated()
	if counter.GeneratedOptions != 1 {
		t.Error("Expected 1, got", counter.GeneratedOptions)
	}
}

func TestThreadSafeCounters_IncrementValid(t *testing.T) {
	counter := ThreadSafeCounters{}
	counter.IncrementValid()
	if counter.ValidOptions != 1 {
		t.Error("Expected 1, got", counter.ValidOptions)
	}
}

func TestThreadSafeCounters_GetChecked(t *testing.T) {
	counter := ThreadSafeCounters{}
	counter.CheckedOptions = 1
	if counter.GetChecked() != 1 {
		t.Error("Expected 1, got", counter.GetChecked())
	}
}

func TestThreadSafeCounters_GetGenerated(t *testing.T) {
	counter := ThreadSafeCounters{}
	counter.GeneratedOptions = 1
	if counter.GetGenerated() != 1 {
		t.Error("Expected 1, got", counter.GetGenerated())
	}
}

func TestThreadSafeCounters_GetValid(t *testing.T) {
	counter := ThreadSafeCounters{}
	counter.ValidOptions = 1
	if counter.GetValid() != 1 {
		t.Error("Expected 1, got", counter.GetValid())
	}
}
