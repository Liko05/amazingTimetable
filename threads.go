package main

import log "github.com/sirupsen/logrus"

func generateCalendars(hashes *threadSafeListOfHashes, shouldFinish chan bool, timeTables chan Table) {
	defaultTimeTable := Table{}
	defaultTimeTable.createDefault()
	for {
		defaultTimeTable = defaultTimeTable.generateNewTimeTable()
		log.Debug(defaultTimeTable.prettyPrint())
		if !defaultTimeTable.checkIfHashAlreadyExists(hashes) {
			hashes.add(defaultTimeTable.hash())
			//timeTables <- defaultTimeTable
			log.Debug("Sent new time table for checking")
		} else {
			log.Info("Found duplicate hash")
			shouldFinish <- true
		}
		log.Debug("Generating again")
	}
}
