package main

import (
	"math/rand"
	"strconv"
)

type subject struct {
	name        string
	teacher     string
	room        string
	isPractical bool
}

var time_table [5][10]subject
var subjectsU []subject

func initTimeTable() {
	time_table = [5][10]subject{}
	createDefault()
	subjectsU = retrieveArrayOfSubjectsWithouPauses()

}

func retrieveSubjects() []subject {
	return subjectsU
}

func createDefault() {
	time_table[0][0] = subject{
		name:        "TP",
		teacher:     "No",
		room:        "23",
		isPractical: false,
	}
	time_table[0][1] = subject{
		name:        "AM",
		teacher:     "Rk",
		room:        "23",
		isPractical: false,
	}
	time_table[0][2] = subject{
		name:        "A",
		teacher:     "Jz",
		room:        "23",
		isPractical: false,
	}
	time_table[0][3] = subject{
		name:        "M",
		teacher:     "Hr",
		room:        "23",
		isPractical: false,
	}
	time_table[0][4] = subject{
		name:        "M",
		teacher:     "Hr",
		room:        "23",
		isPractical: false,
	}
	time_table[0][5] = subject{
		name:        "PIS",
		teacher:     "Bc",
		room:        "23",
		isPractical: false,
	}

	time_table[0][6] = subject{
		name:        "PAUSE",
		teacher:     "",
		room:        "",
		isPractical: false,
	}
	time_table[0][7] = subject{
		name:        "DS",
		teacher:     "Vc",
		room:        "19c",
		isPractical: true,
	}
	time_table[0][8] = subject{
		name:        "DS",
		teacher:     "Vc",
		room:        "19c",
		isPractical: true,
	}
	time_table[1][0] = subject{
		name:        "M",
		teacher:     "Hr",
		room:        "23",
		isPractical: false,
	}

	time_table[1][1] = subject{
		name:        "A",
		teacher:     "Jz",
		room:        "29",
		isPractical: false,
	}
	time_table[1][2] = subject{
		name:        "CIT",
		teacher:     "Sv",
		room:        "17b",
		isPractical: true,
	}
	time_table[1][3] = subject{
		name:        "CIT",
		teacher:     "Sv",
		room:        "17b",
		isPractical: true,
	}
	time_table[1][4] = subject{
		name:        "PAUSE",
		teacher:     "",
		room:        "",
		isPractical: false,
	}
	time_table[1][5] = subject{
		name:        "C",
		teacher:     "Mr",
		room:        "23",
		isPractical: false,
	}
	time_table[1][6] = subject{
		name:        "PV",
		teacher:     "Re",
		room:        "18a",
		isPractical: true,
	}

	time_table[1][7] = subject{
		name:        "PV",
		teacher:     "Re",
		room:        "18a",
		isPractical: true,
	}
	time_table[1][8] = subject{
		name:        "TV",
		teacher:     "Lc",
		room:        "TV",
		isPractical: false,
	}

	time_table[2][0] = subject{
		name:        "PAUSE",
		teacher:     "",
		room:        "",
		isPractical: false,
	}
	time_table[2][1] = subject{
		name:        "A",
		teacher:     "Jz",
		room:        "23",
		isPractical: false,
	}
	time_table[2][2] = subject{
		name:        "C",
		teacher:     "Mr",
		room:        "23",
		isPractical: false,
	}

	time_table[2][3] = subject{
		name:        "WA",
		teacher:     "Ad",
		room:        "23",
		isPractical: false,
	}

	time_table[2][4] = subject{
		name:        "DS",
		teacher:     "Vc",
		room:        "23",
		isPractical: false,
	}

	time_table[2][5] = subject{
		name:        "PSS",
		teacher:     "Mo",
		room:        "8a",
		isPractical: true,
	}

	time_table[2][6] = subject{
		name:        "PSS",
		teacher:     "Mo",
		room:        "8a",
		isPractical: true,
	}

	time_table[3][0] = subject{
		name:        "WA",
		teacher:     "Na",
		room:        "19b",
		isPractical: true,
	}

	time_table[3][1] = subject{
		name:        "WA",
		teacher:     "Na",
		room:        "19b",
		isPractical: true,
	}

	time_table[3][2] = subject{
		name:        "M",
		teacher:     "Hr",
		room:        "23",
		isPractical: false,
	}

	time_table[3][3] = subject{
		name:        "A",
		teacher:     "Jz",
		room:        "23",
		isPractical: false,
	}

	time_table[3][4] = subject{
		name:        "PV",
		teacher:     "Re",
		room:        "23",
		isPractical: false,
	}
	time_table[3][5] = subject{
		name:        "PIS",
		teacher:     "Bc",
		room:        "19a",
		isPractical: true,
	}
	time_table[3][6] = subject{
		name:        "PIS",
		teacher:     "Bc",
		room:        "19a",
		isPractical: true,
	}
	time_table[4][0] = subject{
		name:        "PAUSE",
		teacher:     "",
		room:        "",
		isPractical: false,
	}
	time_table[4][1] = subject{
		name:        "C",
		teacher:     "Mr",
		room:        "23",
		isPractical: false,
	}
	time_table[4][2] = subject{
		name:        "PSS",
		teacher:     "Ms",
		room:        "23",
		isPractical: false,
	}
	time_table[4][3] = subject{
		name:        "AM",
		teacher:     "Rk",
		room:        "23",
		isPractical: false,
	}
	time_table[4][4] = subject{
		name:        "PIS",
		teacher:     "Bc",
		room:        "23",
		isPractical: false,
	}
	time_table[4][5] = subject{
		name:        "TV",
		teacher:     "Lc",
		room:        "TV",
		isPractical: false,
	}
}

func retrieveArrayOfSubjectsWithouPauses() []subject {
	var subjects []subject
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			currentSubject := time_table[i][j]
			if currentSubject.name != "PAUSE" && currentSubject.name != "" {
				subjects = append(subjects, currentSubject)
			}
		}
	}
	//shuffle the array

	return subjects
}

func prettyPrint() string {
	var str string
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			str += time_table[i][j].name + " " + strconv.FormatBool(time_table[i][j].isPractical) + " | "
		}
		str += "\n"
	}
	return str
}

func prettyPrintWithInput(timeTableNew *[5][10]subject) string {
	var str string
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			str += timeTableNew[i][j].name + " " + strconv.FormatBool(timeTableNew[i][j].isPractical) + " | "
		}
		str += "\n"
	}
	return str
}

func generateNewTimeTable(subjects []subject) [5][10]subject {
	newTimeTable := [5][10]subject{}

	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {
			currentSubjectIndex := rand.Intn(len(subjects))
			currentSubjectPick := subjects[currentSubjectIndex]
			if currentSubjectPick.isPractical {
				newTimeTable[i][j] = currentSubjectPick
				newTimeTable[i][j+1] = currentSubjectPick
				j++
				subjects = removeSubjectRemainingOccurences(subjects, currentSubjectPick)
			} else {
				newTimeTable[i][j] = currentSubjectPick
				subjects = append(subjects[:currentSubjectIndex], subjects[currentSubjectIndex+1:]...)
			}
		}
	}

	return newTimeTable
}

func removeSubjectRemainingOccurences(subjects []subject, subjectToRemove subject) []subject {
	for i := 0; i < len(subjects); i++ {
		if subjects[i].name == subjectToRemove.name && subjects[i].isPractical {
			subjects = append(subjects[:i], subjects[i+1:]...)
		}
	}
	return subjects
}
