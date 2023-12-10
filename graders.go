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

func (g *Graders) start() {
	for i := 0; i < g.NumberOfWorkers; i++ {
		go g.gradeTimeTablesStartWorker()
	}
}

func (g *Graders) gradeTimeTablesStartWorker() {
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

		//mandatory rules which might return early so we dont waster resources
		tb.Score += tb.lunchBreaks(dayIndex)
		tb.Score += tb.legalityOfTheDay(dayIndex)
		if !tb.isPracticalSubjectConnecting(dayIndex) || tb.Score < -20000 {
			return false
		}

		//optional rules
		tb.Score += tb.gradeSubjectsInDay(dayIndex)
		tb.Score += tb.roomChangePoints(dayIndex)
		tb.Score += tb.gradeIfSubjectOccursMultipleTimes(dayIndex)
		tb.Score += tb.checkIfStartsWithProfileSubjects(dayIndex)
		tb.Score += tb.checkIfEndsWithProfileSubjects(dayIndex)
		tb.Score += tb.isFirstClassHighestFloor(dayIndex)
	}

	return true
}

func (tb *Table) roomChangePoints(dayIndex int) int {
	final := 0
	for i := dayIndex; i < dayIndex+9; i++ {
		currentSubject := tb.TimeTable[i]
		if i-dayIndex != 0 {
			previousSubject := tb.TimeTable[i-1]
			if previousSubject.Room == currentSubject.Room && previousSubject.Name != "" {
				final += 100
			} else if previousSubject.Floor == currentSubject.Floor {
				final += 50
			} else {
				final -= 100
			}
		}
	}
	return final
}

func (tb *Table) lunchBreaks(dayIndex int) int {
	for i := dayIndex + 4; i < dayIndex+9; i++ {
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

func (tb *Table) gradeIfSubjectOccursMultipleTimes(dayIndex int) int {
	finalScore := 0
	subjectOccurrences := make(map[string]int)
	for i := dayIndex; i <= dayIndex+9; i++ {
		if tb.TimeTable[i].Name != "" && !tb.TimeTable[i].IsPractical {
			subjectOccurrences[tb.TimeTable[i].Name]++
		}
	}

	for _, occurrence := range subjectOccurrences {
		if occurrence > 1 {
			finalScore -= 100
		}
	}

	return finalScore
}

func (tb *Table) checkIfStartsWithProfileSubjects(dayIndex int) int {
	finalScore := 0
	profileSubjects := []string{"M", "PV", "WA", "C", "DS", "PSS"}
	for i := dayIndex; i <= dayIndex+9; i++ {
		if i-dayIndex == 0 {
			for _, subject := range profileSubjects {
				if tb.TimeTable[i].Name == subject {
					finalScore -= 100
				}
			}
		}
	}
	return finalScore
}

func (tb *Table) checkIfEndsWithProfileSubjects(dayIndex int) int {
	finalScore := 0
	profileSubjects := []string{"M", "PV", "WA", "C", "DS", "PSS"}
	lunchBreakIndex := 0
	for i := dayIndex + 4; i <= dayIndex+9; i++ {
		if tb.TimeTable[i].Name == "" {
			lunchBreakIndex = i
			break
		}
	}

	for i := lunchBreakIndex + 1; i <= dayIndex+9; i++ {
		if tb.TimeTable[i].Name != "" {
			for _, subject := range profileSubjects {
				if tb.TimeTable[i].Name == subject {
					finalScore -= 100

				}
			}
			break
		}
	}
	return finalScore
}

func (tb *Table) isFirstClassHighestFloor(dayIndex int) int {
	if tb.TimeTable[dayIndex].Floor == 4 && tb.TimeTable[dayIndex].Name != "" {
		return -250
	} else {
		for i := dayIndex + 1; i <= dayIndex+9; i++ {
			if tb.TimeTable[i].Name != "" {
				if tb.TimeTable[i].Floor == 4 {
					return -250
				} else {
					return -0
				}
			}
		}
	}
	return 0
}
