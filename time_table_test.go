package main

import "testing"

func TestTable_CreateDefault(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	if tb.TimeTable[0].Name != "TP" {
		t.Errorf("Expected TP, got %s", tb.TimeTable[0].Name)
	}
	if tb.TimeTable[1].Name != "AM" {
		t.Errorf("Expected AM, got %s", tb.TimeTable[1].Name)
	}
	if tb.TimeTable[2].Name != "A" {
		t.Errorf("Expected A, got %s", tb.TimeTable[2].Name)
	}
	if tb.TimeTable[3].Name != "M" {
		t.Errorf("Expected M, got %s", tb.TimeTable[3].Name)
	}
	if tb.TimeTable[4].Name != "M" {
		t.Errorf("Expected M, got %s", tb.TimeTable[4].Name)
	}
}

func TestTable_Shuffle(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()
	hashOriginal := tb.Hash()
	tb.Shuffle()
	hashShuffled := tb.Hash()

	if hashOriginal == hashShuffled {
		t.Errorf("Expected different hash, got %s", hashOriginal)
	}
}
