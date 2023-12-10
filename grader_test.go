package main

import (
	"testing"
)

func TestIsPracticalSubjectConnecting(t *testing.T) {
	tb := Table{}
	tb.createDefault()

	dayIndex := 0

	results := [5]bool{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.isPracticalSubjectConnecting(dayIndex)
	}

	expected := [5]bool{true, true, true, true, true}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("isPracticalSubjectConnecting(%v) = %v, want %v", dayIndex, results[i], expected[i])
		}
	}
}

func TestLunchBreaks(t *testing.T) {
	tb := Table{}
	tb.createDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.lunchBreaks(dayIndex)
	}

	expected := [5]int{100, 100, 100, 100, 100}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("lunchBreaks(%v) = %v, want %v", dayIndex, results[i], expected[i])
		}
	}
}

func TestLegalityOfTheDay(t *testing.T) {
	tb := Table{}
	tb.createDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.legalityOfTheDay(dayIndex)
	}

	expected := [5]int{-25, -25, 25, 0, 25}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("legalityOfTheDay(%v) = %v, want %v", dayIndex, results[i], expected[i])
		}
	}
}

func TestGradeSubjectsInDay(t *testing.T) {
	tb := Table{}
	tb.createDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.gradeSubjectsInDay(dayIndex)
	}

	expected := [5]int{400, 200, 400, 500, 500}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("gradeSubjectsInDay(%v) = %v, want %v", i*10, results[i], expected[i])
		}
	}
}

func TestRoomChangePoints(t *testing.T) {
	tb := Table{}
	tb.createDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.roomChangePoints(dayIndex)
	}

	expected := [5]int{400, -250, 150, 150, 250}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("roomChangePoints(%v) = %v, want %v", dayIndex, results[i], expected[i])
		}
	}

}

func TestGradeIfSubjectOccursMultipleTimes(t *testing.T) {
	tb := Table{}
	tb.createDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.gradeIfSubjectOccursMultipleTimes(dayIndex)
	}

	expected := [5]int{-100, 0, 0, 0, 0}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("gradeIfSubjectOccursMultipleTimes(%v) = %v, want %v", dayIndex, results[i], expected[i])
		}
	}

}

func TestIsWeekReasonable(t *testing.T) {
	tb := Table{}
	tb.createDefault()

	dayIndex := 0

	result := false

	result = tb.isWeekReasonable()

	if !result {
		t.Errorf("isWeekReasonable(%v) = %v, want %v", dayIndex, result, true)
	}
}

func TestCheckIfStartsWithProfileSubjects(t *testing.T) {
	tb := Table{}
	tb.createDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.checkIfStartsWithProfileSubjects(dayIndex)
	}

	expected := [5]int{0, -100, 0, -100, 0}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("checkIfStartsWithProfileSubjects(%v) = %v, want %v", i*10, results[i], expected[i])
		}
	}
}

func TestCheckIfEndsWithProfileSubjects(t *testing.T) {
	tb := Table{}
	tb.createDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.checkIfEndsWithProfileSubjects(dayIndex)
	}

	expected := [5]int{-100, -100, 0, 0, 0}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("checkIfEndsWithProfileSubjects(%v) = %v, want %v", i*10, results[i], expected[i])
		}
	}
}
