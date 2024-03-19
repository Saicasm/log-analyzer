package main

import (
	"fmt"
	"github.com/cli/internal/handlers"
	"github.com/cli/internal/logger"
	"github.com/sirupsen/logrus"
	"os"
)

func RunAsCli() {
	log := logger.GetLogger()
	if len(os.Args) < 2 {
		fmt.Println("Please provide at least one argument.")
		return
	}
	fileNames, dates := cliArgsSeparator(os.Args[1:])
	logger.LogWithFields(log, logrus.DebugLevel, "main", "CLI Parameters", map[string]interface{}{
		"filenames": fileNames,
		"dates":     dates,
	})
	// TODO: Do we need parallel processing handling multiple files and dates?
	if len(fileNames) > 1 || len(dates) > 1 {
	}

	cookies, err := handlers.GetMostActiveCookie(fileNames[0], dates[0])
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error Getting Active Cookie")
	}
	log.WithFields(logrus.Fields{
		"cookie": cookies,
	}).Debug("Cookie Value")
	if len(cookies) > 1 {
		for _, cookie := range cookies {
			fmt.Println(cookie)
		}

	} else {
		fmt.Println(cookies[0])
	}
}

// Adding the filename and dates to a slice since the future requirement might change into multiple filenames or dates
func cliArgsSeparator(args []string) ([]string, []string) {
	// Since the requirement doesn't have -f flag for filenames, removed its implementation
	var fileNames, dates []string
	foundDateFlag := false
	for _, arg := range args {
		if arg == "-d" {
			foundDateFlag = true
			continue
		}
		if !foundDateFlag {
			fileNames = append(fileNames, arg)
		} else {
			dates = append(dates, arg)
		}
	}
	return fileNames, dates
}
