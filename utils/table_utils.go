// Package utils Contains helper functions that are used in other packages
package utils

import (
	table "amazingTimetable/table"
	"github.com/fatih/color"
	tb "github.com/rodaine/table"
)

// TableToString returns a string representation of the table Definition of the string representation can be found in the Implementation notes
func TableToString(t table.Table) {
	var subjects = getMapOfSubjects()
	var teachers = getMapOfTeachers()
	var rooms = getMapOfRooms()

	headerFmt := color.New(color.BgHiCyan, color.Bold).SprintfFunc()
	columnFmt := color.New(color.BgHiCyan, color.Bold).SprintfFunc()

	days := []string{"Mon", "Tue", "Wed", "Thu", "Fri"}

	tbl := tb.New("Day", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(5)

	for i := 0; i < 5; i++ {
		var subjectsInDay [10]string = [10]string{"", "", "", "", "", "", "", "", "", ""}
		for j := 0; j < 10; j++ {
			subjIndex := t.TimeTable[j+i*10].Name
			if subjIndex > 100 {
				subjIndex -= 100
			}
			if subjIndex == 0 {
				subjectsInDay[j] = ""
			} else {
				subjectsInDay[j] = subjects[subjIndex] + " : " + teachers[t.TimeTable[j+i*10].Teacher] + " : " + rooms[t.TimeTable[j+i*10].Room]
			}
		}
		tbl.AddRow(days[i], subjectsInDay[0], subjectsInDay[1], subjectsInDay[2], subjectsInDay[3], subjectsInDay[4], subjectsInDay[5], subjectsInDay[6], subjectsInDay[7], subjectsInDay[8], subjectsInDay[9])
	}

	tbl.Print()
}

// getMapOfSubjects returns a mapped subject ids to their names
func getMapOfSubjects() map[uint8]string {
	var result = make(map[uint8]string)
	result[1] = "TP"
	result[2] = "AM"
	result[3] = "A"
	result[4] = "M"
	result[5] = "PIS"
	result[6] = "DS"
	result[7] = "CIT"
	result[8] = "C"
	result[9] = "TV"
	result[10] = "WA"
	result[11] = "PSS"
	result[12] = "PV"
	return result
}

// getMapOfTeachers returns a mapped teacher ids to their names
func getMapOfTeachers() map[uint8]string {
	var result = make(map[uint8]string)
	result[1] = "No"
	result[2] = "Rk"
	result[3] = "Jz"
	result[4] = "Hr"
	result[5] = "Bc"
	result[6] = "Vc"
	result[7] = "Sv"
	result[8] = "Mr"
	result[9] = "Re"
	result[10] = "Lc"
	result[11] = "Ad"
	result[12] = "Mo"
	result[13] = "Na"
	result[14] = "Ms"
	return result

}

// getMapOfRooms returns a mapped room ids to their names
func getMapOfRooms() map[uint8]string {
	var result = make(map[uint8]string)
	result[1] = "23"
	result[2] = "19c"
	result[3] = "29"
	result[4] = "17b"
	result[5] = "18a"
	result[6] = "TV"
	result[7] = "8a"
	result[8] = "19b"
	result[9] = "19a"

	return result
}
