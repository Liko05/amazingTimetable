// Package counter is a package that holds the counters for the program
package counter

import (
	"amazingTimetable/table"
	"sync"
)

// ThreadSafeCounters is a struct that holds counters for the program
// It is thread safe and can be used by multiple goroutines at the same time
type ThreadSafeCounters struct {
	Mu                       sync.Mutex
	GeneratedOptions         uint64
	CheckedOptions           uint64
	ValidOptions             uint64
	OptionsBetterThanDefault uint64
	BestOption               table.Table
	OriginalOption           table.Table
	StopGeneration           bool
}

// Stop stops the generation of new timetables
func (l *ThreadSafeCounters) Stop() {
	l.Mu.Lock()
	l.StopGeneration = true
	l.Mu.Unlock()
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

// SetBestOption Checks if the table is better than the current best option and sets it if it is
func (l *ThreadSafeCounters) SetBestOption(table table.Table) {
	l.Mu.Lock()
	if table.Score > l.BestOption.Score {
		l.BestOption = table
	}
	if table.Score > l.OriginalOption.Score {
		l.OptionsBetterThanDefault++
	}
	l.Mu.Unlock()
}

// GetBestOption returns the best option
func (l *ThreadSafeCounters) GetBestOption() table.Table {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	return l.BestOption
}

// GetOriginalOption returns the original option
func (l *ThreadSafeCounters) GetOriginalOption() table.Table {
	l.Mu.Lock()
	defer l.Mu.Unlock()
	return l.OriginalOption
}
