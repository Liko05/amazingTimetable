package table

import (
	"testing"
)

func TestTable_Shuffle(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()
	hashOriginal := tb.Hash()
	tb.Shuffle()
	hashShuffled := tb.Hash()

	if hashOriginal == hashShuffled {
		t.Errorf("Expected different hash, got %v", hashOriginal)
	}
}

func TestTable_CheckConsecutiveClasses(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	for i := 0; i < 5; i++ {
		if !tb.CheckConsecutiveClasses(i) {
			t.Errorf("Expected true, got %v", tb.CheckConsecutiveClasses(i))
		}
	}
}

func TestTable_CheckConsecutiveClasses_Fails(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()
	tb.TimeTable[0] = Subject{
		Name:    50,
		Teacher: 100,
		Room:    100,
		Floor:   100,
	}
	tb.TimeTable[1] = Subject{
		Name:    100,
		Teacher: 100,
		Room:    100,
		Floor:   100,
	}
	tb.TimeTable[2] = Subject{
		Name:    99,
		Teacher: 100,
		Room:    100,
		Floor:   100,
	}

	if tb.CheckConsecutiveClasses(0) {
		t.Errorf("Expected false, got %v", tb.CheckConsecutiveClasses(0))
	}
}

func TestTable_LegalityOfTheDay(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	for i := 0; i < 5; i++ {
		if !tb.LegalityOfTheDay(i) {
			t.Errorf("Expected true, got %v", tb.LegalityOfTheDay(i))
		}
	}
}

func TestTable_IsThereLunchPause(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	for i := 0; i < 5; i++ {
		if !tb.IsThereLunchPause(i) {
			t.Errorf("Expected true, got %v", tb.IsThereLunchPause(i))
		}
	}
}
