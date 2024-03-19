package handlers

import (
	"encoding/csv"
	"github.com/cli/internal/constants"
	"github.com/cli/internal/logger"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func GetMostActiveCookie(logFile string, specifiedDate string) ([]string, error) {
	log := logger.GetLogger()
	emptyStrings := make([]string, 0)

	file, err := os.Open(logFile)
	if err != nil {
		return emptyStrings, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return emptyStrings, err
	}

	activeCookies := make(map[string]int)

	targetDateChecker, err := time.Parse(constants.InputDateLayout, specifiedDate)
	if err != nil {
		log.WithFields(logrus.Fields{
			"Error": err,
		}).Error("Cannot Parse Date")
		return emptyStrings, err
	}
	log.WithFields(logrus.Fields{
		"Target Date": targetDateChecker,
	}).Debug("Date checker")

	for _, record := range records[1:] { // skipping header

		log.WithFields(logrus.Fields{
			"record": record[0],
		}).Debug("Record Value")
		cookie := record[0]
		timestamp, err := time.Parse(time.RFC3339, record[1])
		if err != nil {
			return emptyStrings, err
		}
		if timestamp.Format(constants.InputDateLayout) == specifiedDate {
			activeCookies[cookie]++
		}
	}

	if len(activeCookies) == 0 {
		return emptyStrings, nil
	}

	var maxCount int
	var mostActiveCookie []string

	for cookie, count := range activeCookies {
		if count > maxCount {
			mostActiveCookie = nil
			maxCount = count
		}
		if count == maxCount {
			mostActiveCookie = append(mostActiveCookie, cookie)
		}
	}
	log.WithFields(logrus.Fields{
		"Most active cookies": mostActiveCookie,
	}).Debug("Record Value")
	return mostActiveCookie, nil
}
