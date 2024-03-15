package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func InitializeLogger() {
	// Output to stdout instead of the default stderr
	log.Out = os.Stdout

	// Set log level
	log.SetLevel(logrus.InfoLevel)

	// Optionally, you can customize log format
	log.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

}
func GetLogger() *logrus.Logger {
	return log
}

// LogWithFields logs a message with additional fields
func LogWithFields(level logrus.Level, message string, fields map[string]interface{}) {
	log.WithFields(fields).Log(level, message)
}
