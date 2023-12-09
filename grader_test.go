package main

import "testing"

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
