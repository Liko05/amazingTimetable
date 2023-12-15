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

func TestTable_IsTableValid(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	if !tb.IsTableValid() {
		t.Errorf("Expected true, got %v", tb.IsTableValid())
	}
}

func TestTable_IsTableValid_Fails(t *testing.T) {
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

	if tb.IsTableValid() {
		t.Errorf("Expected false, got %v", tb.IsTableValid())
	}
}

func TestTable_FirstFloorStart(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	output := [5]int{}

	for i := 0; i < 5; i++ {
		output[i] = int(tb.firstFloorStart(i))
	}

	expected := [5]int{-100, -100, -100, -100, -100}

	if output != expected {
		t.Errorf("Expected %v, got %v", expected, output)
	}

}

func TestTable_IsFridayShort(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	if tb.isFridayShort() != 100 {
		t.Errorf("Expected 100, got %v", tb.isFridayShort())
	}
}

func TestTable_IsFridayShort_Fails(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()
	tb.TimeTable[40] = Subject{
		Name:    110,
		Teacher: 110,
		Room:    110,
		Floor:   110,
	}
	tb.TimeTable[41] = Subject{
		Name:    110,
		Teacher: 110,
		Room:    110,
		Floor:   110,
	}
	tb.TimeTable[42] = Subject{
		Name:    110,
		Teacher: 110,
		Room:    110,
		Floor:   110,
	}
	tb.TimeTable[43] = Subject{
		Name:    110,
		Teacher: 110,
		Room:    110,
		Floor:   110,
	}
	tb.TimeTable[44] = Subject{
		Name:    110,
		Teacher: 110,
		Room:    110,
		Floor:   110,
	}

	if tb.isFridayShort() != 0 {
		t.Errorf("Expected 0, got %v", tb.isFridayShort())
	}
}

func TestTable_RoomChanges(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	output := [5]int{}

	for i := 0; i < 5; i++ {
		output[i] = int(tb.roomChanges(i))
	}

	expected := [5]int{150, -450, 150, 150, 350}

	if output != expected {
		t.Errorf("Expected %v, got %v", expected, output)
	}
}

func TestTable_SameSubjectInDay(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	output := [5]int{}

	for i := 0; i < 5; i++ {
		output[i] = int(tb.sameSubjectInDay(i))
	}

	expected := [5]int{-100, 0, 0, 0, 0}

	if output != expected {
		t.Errorf("Expected %v, got %v", expected, output)
	}
}

func TestTable_ProfileSubjectsFirstOrAfterPause(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	output := [5]int{}

	for i := 0; i < 5; i++ {
		output[i] = int(tb.profileSubjectsFirstOrAfterPause(i))
	}

	expected := [5]int{-100, -100, 0, -100, 0}

	if output != expected {
		t.Errorf("Expected %v, got %v", expected, output)
	}
}

func TestTable_GradePresentClasses(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	output := [5]int{}

	for i := 0; i < 5; i++ {
		output[i] = int(tb.gradePresentClasses(i))
	}

	expected := [5]int{400, 200, 400, 500, 500}

	if output != expected {
		t.Errorf("Expected %v, got %v", expected, output)
	}
}
