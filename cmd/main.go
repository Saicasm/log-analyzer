package main

import (
	"flag"
	"github.com/cli/internal/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	logger.InitializeLogger()
	log := logger.GetLogger()
	fileName := flag.String("f", "", "The filename to read")
	date := flag.String("d", "", "The date to be parsed")

	flag.Parse()
	// Check if a filename was provided with the -f flag
	if *fileName == "" {
		log.Error("Please provide a filename using the -f flag.")
		return
	}
	//readFile(*fileName)

	if *date == "" {
		log.Error("Please provide the date using the -d flag")
		return
	}
	logger.LogWithFields(logrus.InfoLevel, "CLI Parameters", map[string]interface{}{
		"filename": *fileName,
		"date":     *date,
	})
	//fmt.Println(cookie)
}