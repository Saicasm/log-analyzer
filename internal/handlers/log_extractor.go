package handlers

import (
	"encoding/csv"
	"fmt"
	"github.com/cli/internal/constants"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func GetMostActiveCookie(log *logrus.Logger, logFile string, specifiedDate string) ([]string, error) {
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

	targetDate, err := time.Parse(constants.InputDateLayout, specifiedDate)
	fmt.Println("Target Date", targetDate)
	if err != nil {
		return emptyStrings, err
	}

	for _, record := range records[1:] { // skipping header

		log.WithFields(logrus.Fields{
			"record": record[0],
		}).Info("Record Value")
		cookie := record[0]
		timestamp, err := time.Parse(time.RFC3339, record[1])
		if err != nil {
			return emptyStrings, err
		}
		if timestamp.Format(constants.InputDateLayout) == specifiedDate {
			activeCookies[cookie]++
		}
	}

	log.WithFields(logrus.Fields{
		"active cookies": activeCookies,
	}).Info("Record Value")
	if len(activeCookies) == 0 {
		return emptyStrings, nil
	}

	var maxCount int
	var mostActiveCookie []string

	for cookie, count := range activeCookies {
		if count > maxCount {
			maxCount = count
			mostActiveCookie = append(mostActiveCookie, cookie)
		} else if count == maxCount {
			mostActiveCookie = append(mostActiveCookie, cookie)
		}
	}
	log.WithFields(logrus.Fields{
		"Most active cookies": mostActiveCookie,
	}).Info("Record Value")
	return mostActiveCookie, nil
}
