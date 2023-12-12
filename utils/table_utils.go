package utils

import (
	table "amazingTimetable/table"
	"strconv"
)

// TableToString returns a string representation of the table Definition of the string representation can be found in the Implementation notes
func TableToString(t table.Table) string {
	var result string
	var subjects = getMapOfSubjects()
	var teachers = getMapOfTeachers()
	var rooms = getMapOfRooms()
	for i := 0; i < len(t.TimeTable); i++ {
		currentSubject := t.TimeTable[i]
		isPractical := currentSubject.Name > 100
		if isPractical {
			currentSubject.Name -= 100
		}
		result += "Subject: " + subjects[currentSubject.Name] + " Teacher: " + teachers[currentSubject.Teacher] + " Room: " + rooms[currentSubject.Room] + " Floor: " + strconv.Itoa(int(currentSubject.Floor)) + " Is Practical: " + strconv.FormatBool(isPractical) + "\n"
	}
	return result
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
