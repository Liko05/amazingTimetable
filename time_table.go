package main

import (
	"crypto/md5"
	"encoding/hex"
	log "github.com/sirupsen/logrus"
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
	TimeTable map[int]map[int]Subject
	Hash      string
}

func (tb *Table) createDefault() {
	tb.TimeTable = make(map[int]map[int]Subject)
	for i := 0; i < 5; i++ {
		tb.TimeTable[i] = make(map[int]Subject)
	}
	tb.TimeTable[0][0] = Subject{
		Name:        "TP",
		Teacher:     "No",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[0][1] = Subject{
		Name:        "AM",
		Teacher:     "Rk",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[0][2] = Subject{
		Name:        "A",
		Teacher:     "Jz",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[0][3] = Subject{
		Name:        "M",
		Teacher:     "Hr",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[0][4] = Subject{
		Name:        "M",
		Teacher:     "Hr",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[0][5] = Subject{
		Name:        "PIS",
		Teacher:     "Bc",
		Room:        "23",
		IsPractical: false,
	}

	tb.TimeTable[0][6] = Subject{
		Name:        "PAUSE",
		Teacher:     "",
		Room:        "",
		IsPractical: false,
	}
	tb.TimeTable[0][7] = Subject{
		Name:        "DS",
		Teacher:     "Vc",
		Room:        "19c",
		IsPractical: true,
	}
	tb.TimeTable[0][8] = Subject{
		Name:        "DS",
		Teacher:     "Vc",
		Room:        "19c",
		IsPractical: true,
	}
	tb.TimeTable[1][0] = Subject{
		Name:        "M",
		Teacher:     "Hr",
		Room:        "23",
		IsPractical: false,
	}

	tb.TimeTable[1][1] = Subject{
		Name:        "A",
		Teacher:     "Jz",
		Room:        "29",
		IsPractical: false,
	}
	tb.TimeTable[1][2] = Subject{
		Name:        "CIT",
		Teacher:     "Sv",
		Room:        "17b",
		IsPractical: true,
	}
	tb.TimeTable[1][3] = Subject{
		Name:        "CIT",
		Teacher:     "Sv",
		Room:        "17b",
		IsPractical: true,
	}
	tb.TimeTable[1][4] = Subject{
		Name:        "PAUSE",
		Teacher:     "",
		Room:        "",
		IsPractical: false,
	}
	tb.TimeTable[1][5] = Subject{
		Name:        "C",
		Teacher:     "Mr",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[1][6] = Subject{
		Name:        "PV",
		Teacher:     "Re",
		Room:        "18a",
		IsPractical: true,
	}

	tb.TimeTable[1][7] = Subject{
		Name:        "PV",
		Teacher:     "Re",
		Room:        "18a",
		IsPractical: true,
	}
	tb.TimeTable[1][8] = Subject{
		Name:        "TV",
		Teacher:     "Lc",
		Room:        "TV",
		IsPractical: false,
	}

	tb.TimeTable[2][0] = Subject{
		Name:        "PAUSE",
		Teacher:     "",
		Room:        "",
		IsPractical: false,
	}
	tb.TimeTable[2][1] = Subject{
		Name:        "A",
		Teacher:     "Jz",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[2][2] = Subject{
		Name:        "C",
		Teacher:     "Mr",
		Room:        "23",
		IsPractical: false,
	}

	tb.TimeTable[2][3] = Subject{
		Name:        "WA",
		Teacher:     "Ad",
		Room:        "23",
		IsPractical: false,
	}

	tb.TimeTable[2][4] = Subject{
		Name:        "DS",
		Teacher:     "Vc",
		Room:        "23",
		IsPractical: false,
	}

	tb.TimeTable[2][5] = Subject{
		Name:        "PSS",
		Teacher:     "Mo",
		Room:        "8a",
		IsPractical: true,
	}

	tb.TimeTable[2][6] = Subject{
		Name:        "PSS",
		Teacher:     "Mo",
		Room:        "8a",
		IsPractical: true,
	}

	tb.TimeTable[3][0] = Subject{
		Name:        "WA",
		Teacher:     "Na",
		Room:        "19b",
		IsPractical: true,
	}

	tb.TimeTable[3][1] = Subject{
		Name:        "WA",
		Teacher:     "Na",
		Room:        "19b",
		IsPractical: true,
	}

	tb.TimeTable[3][2] = Subject{
		Name:        "M",
		Teacher:     "Hr",
		Room:        "23",
		IsPractical: false,
	}

	tb.TimeTable[3][3] = Subject{
		Name:        "A",
		Teacher:     "Jz",
		Room:        "23",
		IsPractical: false,
	}

	tb.TimeTable[3][4] = Subject{
		Name:        "PV",
		Teacher:     "Re",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[3][5] = Subject{
		Name:        "PIS",
		Teacher:     "Bc",
		Room:        "19a",
		IsPractical: true,
	}
	tb.TimeTable[3][6] = Subject{
		Name:        "PIS",
		Teacher:     "Bc",
		Room:        "19a",
		IsPractical: true,
	}
	tb.TimeTable[4][0] = Subject{
		Name:        "PAUSE",
		Teacher:     "",
		Room:        "",
		IsPractical: false,
	}
	tb.TimeTable[4][1] = Subject{
		Name:        "C",
		Teacher:     "Mr",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[4][2] = Subject{
		Name:        "PSS",
		Teacher:     "Ms",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[4][3] = Subject{
		Name:        "AM",
		Teacher:     "Rk",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[4][4] = Subject{
		Name:        "PIS",
		Teacher:     "Bc",
		Room:        "23",
		IsPractical: false,
	}
	tb.TimeTable[4][5] = Subject{
		Name:        "TV",
		Teacher:     "Lc",
		Room:        "TV",
		IsPractical: false,
	}

	tb.Hash = ""
}

func (tb *Table) retrieveArrayOfSubjectsWithoutPauses() []Subject {
	var subjects []Subject
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			currentSubject := tb.TimeTable[i][j]
			if currentSubject.Name != "PAUSE" && currentSubject.Name != "" {
				subjects = append(subjects, currentSubject)
			}
		}
	}
	return subjects
}

func (tb *Table) prettyPrint() string {
	var str string
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			str += tb.TimeTable[i][j].Name + " " + strconv.FormatBool(tb.TimeTable[i][j].IsPractical) + " | "
		}
		str += "\n"
	}
	return str
}

func (tb *Table) hash() string {
	var hash string
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			hash += tb.TimeTable[i][j].Name
		}
	}
	log.Debug("Hash: " + hash)
	mfhash := md5.Sum([]byte(hash))
	tb.Hash = hex.EncodeToString(mfhash[:])
	return tb.Hash
}

func (tb *Table) generateNewTimeTable(subjs []Subject) Table {
	subjects := subjs
	newTimeTable := make(map[int]map[int]Subject)
	for i := 0; i < 5; i++ {
		newTimeTable[i] = make(map[int]Subject)
	}
	log.Debug("Subjects: " + strconv.Itoa(len(subjects)))

	for i := 0; i < 5; i++ {
		for j := 0; j < 6; j++ {
			currentSubjectIndex := rand.Intn(len(subjects))
			currentSubjectPick := subjects[currentSubjectIndex]
			log.Debug("Current Subject pick: " + currentSubjectPick.Name)
			if currentSubjectPick.IsPractical {
				newTimeTable[i][j] = currentSubjectPick
				newTimeTable[i][j+1] = currentSubjectPick
				j++
				subjects = tb.removeSubjectRemainingOccurences(subjects, currentSubjectPick)
			} else {
				newTimeTable[i][j] = currentSubjectPick
				subjects = append(subjects[:currentSubjectIndex], subjects[currentSubjectIndex+1:]...)
			}
		}
	}

	dayIndexes := []int{0, 1, 2, 3, 4}
	log.Debug("Subjects remaining: " + strconv.Itoa(len(subjects)))

	for i := 0; i < 2; i++ {
		randomDayIndex := rand.Intn(len(dayIndexes))
		randomDay := dayIndexes[randomDayIndex]
		dayIndexes = append(dayIndexes[:randomDayIndex], dayIndexes[randomDayIndex+1:]...)
		if len(subjects) == 0 {
			break
		}
		newTimeTable[randomDay][6] = Subject{
			Name:        "PAUSE",
			Teacher:     "",
			Room:        "",
			IsPractical: false,
		}

		for j := 0; j < 2; j++ {
			if len(subjects) == 0 {
				break
			}
			currentSubjectIndex := rand.Intn(len(subjects))
			currentSubject := subjects[currentSubjectIndex]
			newTimeTable[randomDay][7+j] = currentSubject
			if currentSubject.IsPractical {
				newTimeTable[randomDay][8+j] = currentSubject
				j++
				subjects = tb.removeSubjectRemainingOccurences(subjects, currentSubject)
			} else {
				subjects = append(subjects[:currentSubjectIndex], subjects[currentSubjectIndex+1:]...)
			}
		}

	}

	tb.TimeTable = newTimeTable

	return *tb
}

func (tb *Table) checkIfHashAlreadyExists(hashes *ThreadSafeListOfHashes) bool {
	hash := tb.hash()
	if hashes.contains(hash) {
		log.Debug("Hash already exists: " + hash)
		return true
	} else {
		hashes.add(hash)
		return false
	}
}

func (tb *Table) removeSubjectRemainingOccurences(subjects []Subject, subjectToRemove Subject) []Subject {
	var newSubjects []Subject
	for _, s := range subjects {
		if s != subjectToRemove {
			newSubjects = append(newSubjects, s)
		}
	}
	return newSubjects
}

func (tb *Table) isEmpty() bool {
	return tb.TimeTable[0][0].Name == ""
}
