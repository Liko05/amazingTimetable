package main

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type Graders struct {
	NumberOfWorkers int
	Counters        *ThreadSafeCounters
	ProcessingQueue *ProcessingQueue
}

func (g *Graders) Start() {
	for i := 0; i < g.NumberOfWorkers; i++ {
		go g.GradeTimeTablesStartWorker()
	}
}

func (g *Graders) GradeTimeTablesStartWorker() {
	for {
		log.Debug("Checking new time table")
		table, ok := g.ProcessingQueue.Pop().(Table)
		if ok == false {
			time.Sleep(10 * time.Nanosecond)
			continue
		}

		if table.gradeTimeTable() {

			g.Counters.incrementValid()
			g.ProcessingQueue.AddIfBetter(table)
		}
		g.Counters.incrementChecked()
	}
}

func (tb *Table) gradeTimeTable() bool {

	//dayIndex := 0
	//for i, subject := range tb.TimeTable {
	//	dayIndex = getDayIndex(i)
	//	switch {
	//	case i > 0 && i%10 != 0:
	//		tb.Score += tb.roomChangePoints(i)
	//	}
	//}

	return tb.isWeekReasonable()
}

func (tb *Table) isWeekReasonable() bool {
	dayIndex := 0
	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		tb.Score += tb.lunchBreaks(dayIndex)
		tb.Score += tb.legalityOfTheDay(dayIndex)
		if !tb.isPracticalSubjectConnecting(dayIndex) {
			return false
		}
	}

	if tb.Score >= -20000 {
		return true
	}
	return false
}

func (tb *Table) roomChangePoints(index int) int {
	previous, current := tb.TimeTable[index-1], tb.TimeTable[index]
	if previous.Room == current.Room {
		return 100
	}
	if previous.floorNumber() == current.floorNumber() {
		return 50
	}
	return -25
}

//day rules

func (tb *Table) lunchBreaks(dayIndex int) int {
	for i := dayIndex + 5; i < dayIndex+9; i++ {
		if tb.TimeTable[i].Name == "" {
			return 100
		}
	}
	return -100000
}

func (tb *Table) legalityOfTheDay(dayIndex int) int {
	classes := 0
	for i := dayIndex; i <= dayIndex+9; i++ {
		if tb.TimeTable[i].Name != "" {
			classes++
		}
	}

	switch {
	case classes < 7:
		return 25
	case classes == 7:
		return 0
	case classes == 8 || classes == 9:
		return -25
	default:
		return -100000
	}
}

func (tb *Table) isPracticalSubjectConnecting(dayIndex int) bool {
	for i := dayIndex; i < dayIndex+9; i++ {
		currentSubject := tb.TimeTable[i]
		if i-dayIndex != 0 && i != dayIndex+9 {
			if !currentSubject.IsPractical {
				continue
			}
			previousSubject := tb.TimeTable[i-1]
			upcomingSubject := tb.TimeTable[i+1]
			isOkay := false

			if previousSubject.Name == currentSubject.Name && previousSubject.IsPractical {
				isOkay = true
			}
			if upcomingSubject.Name == currentSubject.Name && upcomingSubject.IsPractical {
				isOkay = true
			}

			if isOkay == false {
				return false
			}
		}
	}
	return true
}

func (tb *Table) gradeSubjectsInDay(dayIndex int) int {
	finalScore := 0
	for i := dayIndex; i <= dayIndex+9; i++ {
		if tb.TimeTable[i].Name != "" {
			if i-dayIndex < 6 {
				finalScore += 100
			} else {
				finalScore -= 100
			}
		}
	}
	return finalScore
}
