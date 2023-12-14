package table

import (
	"github.com/spaolacci/murmur3"
	"math/rand"
)

// Subject is a struct that represents a subject
type Subject struct {
	Name    uint8
	Teacher uint8
	Room    uint8
	Floor   uint8
}

// Table is a struct that represents a timetable
type Table struct {
	TimeTable [50]Subject
	Score     int32
}

// CreateDefault creates a default timetable with the current timetable
func (tb *Table) CreateDefault() {
	tb.TimeTable = [50]Subject{}
	tb.Score = 0

	tb.TimeTable[0] = Subject{
		Name:    101,
		Teacher: 1,
		Room:    1,
		Floor:   4,
	}
	tb.TimeTable[1] = Subject{
		Name:    2,
		Teacher: 2,
		Room:    1,
		Floor:   4,
	}
	tb.TimeTable[2] = Subject{
		Name:    3,
		Teacher: 3,
		Room:    1,
		Floor:   4,
	}
	tb.TimeTable[3] = Subject{
		Name:    4,
		Teacher: 4,
		Room:    1,
		Floor:   4,
	}
	tb.TimeTable[4] = Subject{
		Name:    4,
		Teacher: 4,
		Room:    1,
		Floor:   4,
	}
	tb.TimeTable[5] = Subject{
		Name:    5,
		Teacher: 5,
		Room:    1,
		Floor:   4,
	}

	tb.TimeTable[7] = Subject{
		Name:    106,
		Teacher: 6,
		Room:    2,
		Floor:   3,
	}
	tb.TimeTable[8] = Subject{
		Name:    106,
		Teacher: 6,
		Room:    2,
		Floor:   3,
	}
	tb.TimeTable[10] = Subject{
		Name:    4,
		Teacher: 4,
		Room:    1,
		Floor:   4,
	}

	tb.TimeTable[11] = Subject{
		Name:    3,
		Teacher: 3,
		Room:    3,
		Floor:   4,
	}
	tb.TimeTable[12] = Subject{
		Name:    107,
		Teacher: 7,
		Room:    4,
		Floor:   3,
	}
	tb.TimeTable[13] = Subject{
		Name:    107,
		Teacher: 7,
		Room:    4,
		Floor:   3,
	}

	tb.TimeTable[15] = Subject{
		Name:    8,
		Teacher: 8,
		Room:    1,
		Floor:   4,
	}
	tb.TimeTable[16] = Subject{
		Name:    112,
		Teacher: 9,
		Room:    5,
		Floor:   3,
	}

	tb.TimeTable[17] = Subject{
		Name:    112,
		Teacher: 9,
		Room:    5,
		Floor:   3,
	}
	tb.TimeTable[18] = Subject{
		Name:    9,
		Teacher: 10,
		Room:    6,
		Floor:   0,
	}

	tb.TimeTable[21] = Subject{
		Name:    3,
		Teacher: 3,
		Room:    1,
		Floor:   4,
	}
	tb.TimeTable[22] = Subject{
		Name:    8,
		Teacher: 8,
		Room:    1,
		Floor:   4,
	}

	tb.TimeTable[23] = Subject{
		Name:    10,
		Teacher: 11,
		Room:    1,
		Floor:   4,
	}

	tb.TimeTable[24] = Subject{
		Name:    6,
		Teacher: 6,
		Room:    1,
		Floor:   4,
	}

	tb.TimeTable[25] = Subject{
		Name:    111,
		Teacher: 12,
		Room:    7,
		Floor:   2,
	}

	tb.TimeTable[26] = Subject{
		Name:    111,
		Teacher: 12,
		Room:    7,
		Floor:   2,
	}

	tb.TimeTable[30] = Subject{
		Name:    110,
		Teacher: 13,
		Room:    8,
		Floor:   3,
	}

	tb.TimeTable[31] = Subject{
		Name:    110,
		Teacher: 13,
		Room:    8,
		Floor:   3,
	}

	tb.TimeTable[32] = Subject{
		Name:    4,
		Teacher: 4,
		Room:    1,
		Floor:   4,
	}

	tb.TimeTable[33] = Subject{
		Name:    3,
		Teacher: 3,
		Room:    1,
		Floor:   4,
	}

	tb.TimeTable[34] = Subject{
		Name:    12,
		Teacher: 9,
		Room:    1,
		Floor:   4,
	}
	tb.TimeTable[35] = Subject{
		Name:    105,
		Teacher: 5,
		Room:    9,
		Floor:   3,
	}
	tb.TimeTable[36] = Subject{
		Name:    105,
		Teacher: 5,
		Room:    9,
		Floor:   3,
	}

	tb.TimeTable[41] = Subject{
		Name:    8,
		Teacher: 8,
		Room:    1,
		Floor:   4,
	}
	tb.TimeTable[42] = Subject{
		Name:    11,
		Teacher: 14,
		Room:    1,
		Floor:   4,
	}
	tb.TimeTable[43] = Subject{
		Name:    2,
		Teacher: 2,
		Room:    1,
		Floor:   4,
	}
	tb.TimeTable[44] = Subject{
		Name:    5,
		Teacher: 5,
		Room:    1,
		Floor:   4,
	}
	tb.TimeTable[45] = Subject{
		Name:    9,
		Teacher: 10,
		Room:    6,
		Floor:   0,
	}
}

// Shuffle shuffles the timetable
func (tb *Table) Shuffle() {
	for i := len(tb.TimeTable) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		tb.TimeTable[i], tb.TimeTable[j] = tb.TimeTable[j], tb.TimeTable[i]
	}
}

// Hash returns a hash of the timetable using the murmur3 algorithm
func (tb *Table) Hash() uint32 {
	hash := murmur3.New32()
	for i := 0; i < len(tb.TimeTable); i++ {
		_, err := hash.Write([]byte{tb.TimeTable[i].Name})
		if err != nil {
			return 0
		}
	}
	return hash.Sum32()
}

func (tb *Table) IsEmpty() bool {
	for i := 0; i < len(tb.TimeTable); i++ {
		if tb.TimeTable[i].Name > 0 {
			return false
		}
	}
	return true
}

func (tb *Table) CheckConsecutiveClasses(dayIndex int) bool {
	for hours := 1; hours < 9; hours++ {
		if tb.TimeTable[dayIndex*10+hours].Name < 100 {
			continue
		}

		if tb.TimeTable[dayIndex*10+hours-1].Name != tb.TimeTable[dayIndex*10+hours].Name && tb.TimeTable[dayIndex*10+hours+1].Name != tb.TimeTable[dayIndex*10+hours].Name {
			return false
		}
	}
	return true
}

func (tb *Table) LegalityOfTheDay(dayIndex int) bool {
	maxClasses := 0
	for hours := 0; hours < 10; hours++ {
		if tb.TimeTable[dayIndex*10+hours].Name > 0 {
			maxClasses++
		}
	}
	if maxClasses > 8 {
		return false
	}
	return true
}

func (tb *Table) IsThereLunchPause(dayIndex int) bool {
	if tb.TimeTable[dayIndex*10+4].Name == 0 || tb.TimeTable[dayIndex*10+5].Name == 0 || tb.TimeTable[dayIndex*10+6].Name == 0 || tb.TimeTable[dayIndex*10+7].Name == 0 {
		return true
	}
	return false
}
