// Package counter is a package that holds the counters for the program
package counter

import "sync"

// ThreadSafeCounters is a struct that holds counters for the program
// It is thread safe and can be used by multiple goroutines at the same time
type ThreadSafeCounters struct {
	Mu                       sync.Mutex
	GeneratedOptions         uint64
	CheckedOptions           uint64
	ValidOptions             uint64
	OptionsBetterThanDefault uint64
}

// IncrementGenerated increments the number of generated options
func (l *ThreadSafeCounters) IncrementGenerated() {
	l.Mu.Lock()
	l.GeneratedOptions++
	l.Mu.Unlock()
}

// IncrementChecked increments the number of checked options
func (l *ThreadSafeCounters) IncrementChecked() {
	l.Mu.Lock()
	l.CheckedOptions++
	l.Mu.Unlock()
}

// GetGenerated returns the number of generated options
func (l *ThreadSafeCounters) GetGenerated() uint64 {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	return l.GeneratedOptions
}

// GetChecked returns the number of checked options
func (l *ThreadSafeCounters) GetChecked() uint64 {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	return l.CheckedOptions
}

// IncrementValid increments the number of valid options
func (l *ThreadSafeCounters) IncrementValid() {
	l.Mu.Lock()
	l.ValidOptions++
	l.Mu.Unlock()
}

// GetValid returns the number of valid options
func (l *ThreadSafeCounters) GetValid() uint64 {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	return l.ValidOptions
}

// IncrementOptionsBetterThanDefault increments the number of options better than default
func (l *ThreadSafeCounters) IncrementOptionsBetterThanDefault() {
	l.Mu.Lock()
	l.OptionsBetterThanDefault++
	l.Mu.Unlock()
}

// GetOptionsBetterThanDefault returns the number of options better than default
func (l *ThreadSafeCounters) GetOptionsBetterThanDefault() uint64 {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	return l.OptionsBetterThanDefault
}
