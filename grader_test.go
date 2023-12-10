package main

import (
	"testing"
)

func TestIsPracticalSubjectConnecting(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	dayIndex := 0

	results := [5]bool{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.IsPracticalSubjectConnecting(dayIndex)
	}

	expected := [5]bool{true, true, true, true, true}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("IsPracticalSubjectConnecting(%v) = %v, want %v", dayIndex, results[i], expected[i])
		}
	}
}

func TestLunchBreaks(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.LunchBreaks(dayIndex)
	}

	expected := [5]int{100, 100, 100, 100, 100}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("LunchBreaks(%v) = %v, want %v", dayIndex, results[i], expected[i])
		}
	}
}

func TestLegalityOfTheDay(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.LegalityOfTheDay(dayIndex)
	}

	expected := [5]int{-25, -25, 25, 0, 25}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("LegalityOfTheDay(%v) = %v, want %v", dayIndex, results[i], expected[i])
		}
	}
}

func TestGradeSubjectsInDay(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.GradeSubjectsInDay(dayIndex)
	}

	expected := [5]int{400, 200, 400, 500, 500}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("GradeSubjectsInDay(%v) = %v, want %v", i*10, results[i], expected[i])
		}
	}
}

func TestRoomChangePoints(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.RoomChangePoints(dayIndex)
	}

	expected := [5]int{400, -250, 150, 150, 250}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("RoomChangePoints(%v) = %v, want %v", dayIndex, results[i], expected[i])
		}
	}

}

func TestGradeIfSubjectOccursMultipleTimes(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.GradeIfSubjectOccursMultipleTimes(dayIndex)
	}

	expected := [5]int{-100, 0, 0, 0, 0}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("GradeIfSubjectOccursMultipleTimes(%v) = %v, want %v", dayIndex, results[i], expected[i])
		}
	}

}

func TestIsWeekReasonable(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	dayIndex := 0

	result := false

	result = tb.IsWeekReasonable()

	if !result {
		t.Errorf("IsWeekReasonable(%v) = %v, want %v", dayIndex, result, true)
	}
}

func TestCheckIfStartsWithProfileSubjects(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.CheckIfStartsWithProfileSubjects(dayIndex)
	}

	expected := [5]int{0, -100, 0, -100, 0}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("CheckIfStartsWithProfileSubjects(%v) = %v, want %v", i*10, results[i], expected[i])
		}
	}
}

func TestCheckIfEndsWithProfileSubjects(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.CheckIfEndsWithProfileSubjects(dayIndex)
	}

	expected := [5]int{-100, -100, 0, 0, 0}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("CheckIfEndsWithProfileSubjects(%v) = %v, want %v", i*10, results[i], expected[i])
		}
	}
}

func TestIsFirstClassHighestFloor(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.IsFirstClassHighestFloor(dayIndex)
	}

	expected := [5]int{-250, -250, -250, 0, -250}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("IsFirstClassHighestFloor(%v) = %v, want %v", i*10, results[i], expected[i])
		}
	}
}

func TestIsFridayShort(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	result := 0

	result = tb.IsFridayShort()

	if result == -100 {
		t.Errorf("IsFridayShort() = %v, want %v", result, 0)
	}
}

func TestDislikedClassRoomPoint(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	result := tb.DislikedClassRoomPoint()

	if result != -100 {
		t.Errorf("DislikedClassRoomPoint() = %v, want %v", result, 0)
	}

}

func TestDislikedTeachers(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	dayIndex := 0

	results := [5]int{}

	for i := 0; i < 5; i++ {
		dayIndex = i * 10
		results[i] = tb.DislikedTeachers(dayIndex)
	}

	expected := [5]int{-300, -100, -200, -100, 0}

	for i := 0; i < 5; i++ {
		if results[i] != expected[i] {
			t.Errorf("DislikedTeachers(%v) = %v, want %v", i*10, results[i], expected[i])
		}
	}
}

func TestLikedTeachers(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	result := tb.LikedTeachers()

	if result != 0 {
		t.Errorf("LikedTeachers() = %v, want %v", result, 0)
	}
}

func TestTable_WellBeingPoints(t *testing.T) {
	tb := Table{}
	tb.CreateDefault()

	result := 0

	result = tb.WellBeingPoints()

	if result != -800 {
		t.Errorf("WellBeingPoints() = %v, want %v", result, -800)
	}
}
