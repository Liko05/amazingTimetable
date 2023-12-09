package main

import (
	"math/rand"
	"strconv"
)

type Subject struct {
	Name        string
	Teacher     string
	Room        string
	IsPractical bool
}

type Table struct {
	TimeTable [50]Subject
	Score     int
}

func (tb *Table) createDefault() {
	tb.TimeTable = [50]Subject{}
	tb.Score = 0

	tb.TimeTable[0] = Subject{
		Name:        "TP",
		Teacher:     "No",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[1] = Subject{
		Name:        "AM",
		Teacher:     "Rk",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[2] = Subject{
		Name:        "A",
		Teacher:     "Jz",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[3] = Subject{
		Name:        "M",
		Teacher:     "Hr",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[4] = Subject{
		Name:        "M",
		Teacher:     "Hr",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[5] = Subject{
		Name:        "PIS",
		Teacher:     "Bc",
		Room:        "23",
		IsPractical: false,
	}

	tb.TimeTable[7] = Subject{
		Name:        "DS",
		Teacher:     "Vc",
		Room:        "19c",
		IsPractical: true,
	}
	tb.TimeTable[8] = Subject{
		Name:        "DS",
		Teacher:     "Vc",
		Room:        "19c",
		IsPractical: true,
	}
	tb.TimeTable[10] = Subject{
		Name:        "M",
		Teacher:     "Hr",
		Room:        "23",
		IsPractical: false,
	}

	tb.TimeTable[11] = Subject{
		Name:        "A",
		Teacher:     "Jz",
		Room:        "29",
		IsPractical: false,
	}
	tb.TimeTable[12] = Subject{
		Name:        "CIT",
		Teacher:     "Sv",
		Room:        "17b",
		IsPractical: true,
	}
	tb.TimeTable[13] = Subject{
		Name:        "CIT",
		Teacher:     "Sv",
		Room:        "17b",
		IsPractical: true,
	}

	tb.TimeTable[15] = Subject{
		Name:        "C",
		Teacher:     "Mr",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[16] = Subject{
		Name:        "PV",
		Teacher:     "Re",
		Room:        "18a",
		IsPractical: true,
	}

	tb.TimeTable[17] = Subject{
		Name:        "PV",
		Teacher:     "Re",
		Room:        "18a",
		IsPractical: true,
	}
	tb.TimeTable[18] = Subject{
		Name:        "TV",
		Teacher:     "Lc",
		Room:        "TV",
		IsPractical: false,
	}

	tb.TimeTable[21] = Subject{
		Name:        "A",
		Teacher:     "Jz",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[22] = Subject{
		Name:        "C",
		Teacher:     "Mr",
		Room:        "23",
		IsPractical: false,
	}

	tb.TimeTable[23] = Subject{
		Name:        "WA",
		Teacher:     "Ad",
		Room:        "23",
		IsPractical: false,
	}

	tb.TimeTable[24] = Subject{
		Name:        "DS",
		Teacher:     "Vc",
		Room:        "23",
		IsPractical: false,
	}

	tb.TimeTable[25] = Subject{
		Name:        "PSS",
		Teacher:     "Mo",
		Room:        "8a",
		IsPractical: true,
	}

	tb.TimeTable[26] = Subject{
		Name:        "PSS",
		Teacher:     "Mo",
		Room:        "8a",
		IsPractical: true,
	}

	tb.TimeTable[30] = Subject{
		Name:        "WA",
		Teacher:     "Na",
		Room:        "19b",
		IsPractical: true,
	}

	tb.TimeTable[31] = Subject{
		Name:        "WA",
		Teacher:     "Na",
		Room:        "19b",
		IsPractical: true,
	}

	tb.TimeTable[32] = Subject{
		Name:        "M",
		Teacher:     "Hr",
		Room:        "23",
		IsPractical: false,
	}

	tb.TimeTable[33] = Subject{
		Name:        "A",
		Teacher:     "Jz",
		Room:        "23",
		IsPractical: false,
	}

	tb.TimeTable[34] = Subject{
		Name:        "PV",
		Teacher:     "Re",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[35] = Subject{
		Name:        "PIS",
		Teacher:     "Bc",
		Room:        "19a",
		IsPractical: true,
	}
	tb.TimeTable[36] = Subject{
		Name:        "PIS",
		Teacher:     "Bc",
		Room:        "19a",
		IsPractical: true,
	}

	tb.TimeTable[41] = Subject{
		Name:        "C",
		Teacher:     "Mr",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[42] = Subject{
		Name:        "PSS",
		Teacher:     "Ms",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[43] = Subject{
		Name:        "AM",
		Teacher:     "Rk",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[44] = Subject{
		Name:        "PIS",
		Teacher:     "Bc",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[45] = Subject{
		Name:        "TV",
		Teacher:     "Lc",
		Room:        "TV",
		IsPractical: false,
	}
}

func (tb *Table) shuffle() {
	// Fisher-Yates shuffle algorithm
	for i := len(tb.TimeTable) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		tb.TimeTable[i], tb.TimeTable[j] = tb.TimeTable[j], tb.TimeTable[i]
	}
}

func (tb *Table) isClassEmpty(index int) bool {
	return tb.TimeTable[index].Name == ""
}

func (tb *Table) String() string {
	var str string
	for i := 0; i < len(tb.TimeTable); i++ {
		str += "[ " + strconv.Itoa(i) + " : " + tb.TimeTable[i].Name + " ] "
		if i%10 == 0 && i != 0 {
			str += "\n"
		}
	}
	return str
}

func (sb *Subject) floorNumber() int {
	room, _ := strconv.Atoi(sb.Room[:1])
	return room
}
