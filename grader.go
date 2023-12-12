package main

import (
	log "github.com/sirupsen/logrus"
	"time"
)

// Grader is a struct which holds all the information about the grader
type Grader struct {
	NumberOfWorkers int
	Counters        *ThreadSafeCounters
	ProcessingQueue *ProcessingQueue
}

// Start starts the grading of new timetables based on the NumberOfWorkers
func (g *Grader) Start() {
	defaultTb := Table{}
	defaultTb.CreateDefault()
	defaultTb.IsWeekReasonable()
	g.ProcessingQueue.AddOriginal(defaultTb)
	g.ProcessingQueue.AddIfBetter(defaultTb)
	for i := 0; i < g.NumberOfWorkers; i++ {
		go g.GradeTimeTablesStartWorker()
	}
}

// GradeTimeTablesStartWorker is a worker which grades the timetables
func (g *Grader) GradeTimeTablesStartWorker() {
	for {
		log.Debug("Checking new time table")
		table, ok := g.ProcessingQueue.Pop().(Table)
		if ok == false {
			time.Sleep(10 * time.Nanosecond)
			continue
		}

		if table.IsWeekReasonable() {
			g.Counters.IncrementValid()
			g.ProcessingQueue.AddIfBetter(table)
			g.ProcessingQueue.AddToBestTables(table)
		}
		g.Counters.IncrementChecked()
	}
}

// IsWeekReasonable grades the week and returns true if it is reasonable
// It also sets the score of the table
// If the week is not reasonable it returns false and goes back to checking new timetable
func (tb *Table) IsWeekReasonable() bool {
	dayIndex := 0
	if !tb.IsPracticalSubjectConnecting() {
		return false
	}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		//mandatory rules which might return early so we dont waste resources
		tb.Score += tb.LunchBreaks(dayIndex)
		tb.Score += tb.LegalityOfTheDay(dayIndex)

		if tb.Score < -150000 {
			return false
		}

		//optional rules
		tb.Score += tb.GradeSubjectsInDay(dayIndex)
		tb.Score += tb.RoomChangePoints(dayIndex)
		tb.Score += tb.GradeIfSubjectOccursMultipleTimes(dayIndex)
		tb.Score += tb.CheckIfStartsWithProfileSubjects(dayIndex)
		tb.Score += tb.CheckIfEndsWithProfileSubjects(dayIndex)
		tb.Score += tb.IsFirstClassHighestFloor(dayIndex)
		tb.Score += tb.WellBeingPoints()
	}
	tb.Score += tb.IsFridayShort()

	return true
}

// RoomChangePoints grades the room changes in the day based on the distance between the rooms and the floors
func (tb *Table) RoomChangePoints(dayIndex int) int {
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

// LunchBreaks grades the lunch breaks in the day
func (tb *Table) LunchBreaks(dayIndex int) int {
	for i := dayIndex + 4; i < dayIndex+9; i++ {
		if tb.TimeTable[i].Name == "" {
			return 100
		}
	}
	return -100000
}

// LegalityOfTheDay grades the legality of the day based on the number of classes that day
func (tb *Table) LegalityOfTheDay(dayIndex int) int {
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
		return -250
	default:
		return -100000
	}
}

// IsPracticalSubjectConnecting iterates through the day and if it finds a practical subject it checks whether there is an adjacent practical subject in the same day or not.
func (tb *Table) IsPracticalSubjectConnecting() bool {
	for day := 0; day < 5; day++ {
		for period := 0; period < 10; period++ {
			currentSubject := tb.TimeTable[day*10+period]
			if currentSubject.IsPractical {
				if period < 9 {
					rightSubject := tb.TimeTable[day*10+period+1]
					if rightSubject.IsPractical && rightSubject.Name == currentSubject.Name {
						continue
					}
				}
				if period > 0 {
					leftSubject := tb.TimeTable[day*10+period-1]
					if leftSubject.IsPractical && leftSubject.Name == currentSubject.Name {
						continue
					}
				}
				return false
			}
		}
	}
	return true
}

// GradeSubjectsInDay grades the subjects in the day based on the position of the subject in the day
func (tb *Table) GradeSubjectsInDay(dayIndex int) int {
	finalScore := 0
	for i := dayIndex; i <= dayIndex+9; i++ {
		if tb.TimeTable[i].Name != "" {
			if i-dayIndex < 6 {
				finalScore += 100
			} else {
				finalScore -= 100
			}
		} else {
			if i-dayIndex < 6 {
				finalScore -= 200
			} else {
				finalScore += 100
			}
		}
	}
	return finalScore
}

// GradeIfSubjectOccursMultipleTimes grades the non-practical subjects in the day based on the number of times the subject occurs in the day
func (tb *Table) GradeIfSubjectOccursMultipleTimes(dayIndex int) int {
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

// CheckIfStartsWithProfileSubjects grades the day based on whether starts with a profile subject or not
func (tb *Table) CheckIfStartsWithProfileSubjects(dayIndex int) int {
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

// CheckIfEndsWithProfileSubjects grades the day based on whether ends with a profile subject or not
func (tb *Table) CheckIfEndsWithProfileSubjects(dayIndex int) int {
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

// IsFirstClassHighestFloor grades the day based on whether the first class is on the highest floor or not
func (tb *Table) IsFirstClassHighestFloor(dayIndex int) int {
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

// IsFridayShort grades the day based on whether the friday is short or not
func (tb *Table) IsFridayShort() int {
	classes := 0
	for i := 40; i < 50; i++ {
		if tb.TimeTable[i].Name != "" {
			classes++
		}
	}
	if classes < 6 {
		return 100
	} else {
		return -100
	}
}

// WellBeingPoints grades the day based on the well-being of the students
func (tb *Table) WellBeingPoints() int {
	finalScore := 0
	dayIndex := 0
	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		finalScore += tb.DislikedTeachers(dayIndex)
	}
	finalScore += tb.DislikedClassRoomPoint()
	finalScore += tb.LikedTeachers()

	return finalScore

}

// DislikedClassRoomPoint grades classes taking place in room 29
func (tb *Table) DislikedClassRoomPoint() int {
	finalScore := 0
	for i := 0; i < 50; i++ {
		if tb.TimeTable[i].Room == "29" {
			finalScore += -100
		}
	}
	return finalScore
}

// DislikedTeachers grades classes taking place with disliked teachers
func (tb *Table) DislikedTeachers(dayIndex int) int {
	finalScore := 0
	hatedTeachers := []string{"Vc", "Jz"}
	for i := dayIndex; i < dayIndex+9; i++ {
		for _, teacher := range hatedTeachers {
			if tb.TimeTable[i].Teacher == teacher {
				finalScore -= 100
			}
		}
	}
	return finalScore
}

// LikedTeachers grades classes being consecutive with liked teachers
func (tb *Table) LikedTeachers() int {
	likedTeacher := []string{"Re", "Ma", "Mo", "Sv", "Na"}
	finalScore := 0
	dayIndex := 0

	for _, teacher := range likedTeacher {
		occurrences := map[int]bool{}
		for i := 0; i < 5; i++ {
			dayIndex = i * 10
			for j := dayIndex; j <= dayIndex+9; j++ {
				if tb.TimeTable[j].Teacher == teacher {
					occurrences[dayIndex] = true
					if dayIndex > 0 && occurrences[dayIndex-10] == true {
						finalScore += 500
					} else {
					}
				}
			}
		}
	}
	return finalScore
}
