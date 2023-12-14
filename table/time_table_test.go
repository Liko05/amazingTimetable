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
		if !tb.checkConsecutiveClasses(i) {
			t.Errorf("Expected true, got %v", tb.checkConsecutiveClasses(i))
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
		Name:    110,
		Teacher: 100,
		Room:    100,
		Floor:   100,
	}
	tb.TimeTable[2] = Subject{
		Name:    110,
		Teacher: 100,
		Room:    100,
		Floor:   100,
	}

	tb.TimeTable[7] = Subject{
		Name:    110,
		Teacher: 110,
		Room:    110,
		Floor:   110,
	}
	tb.TimeTable[8] = Subject{
		Name:    110,
		Teacher: 110,
		Room:    110,
		Floor:   110,
	}

	tb.TimeTable[9] = Subject{
		Name:    112,
		Teacher: 110,
		Room:    110,
		Floor:   110,
	}

	if tb.checkConsecutiveClasses(0) {
		t.Errorf("Expected false, got %v", tb.checkConsecutiveClasses(0))
	}
}

func TestTable_LegalityOfTheDay(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	for i := 0; i < 5; i++ {
		if !tb.legalityOfTheDay(i) {
			t.Errorf("Expected true, got %v", tb.legalityOfTheDay(i))
		}
	}
}

func TestTable_IsThereLunchPause(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	for i := 0; i < 5; i++ {
		if !tb.isThereLunchPause(i) {
			t.Errorf("Expected true, got %v", tb.isThereLunchPause(i))
		}
	}
}
