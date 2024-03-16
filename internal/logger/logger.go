package logger

import (
	"os"

	"github.com/sirupsen/logrus"
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

// LogWithFields logs a message with additional fields using the logger instance
func LogWithFields(log *logrus.Logger, level logrus.Level, category string, message string, fields map[string]interface{}) {
	entry := log.WithFields(fields)
	entry.Data["category"] = category
	entry.Log(level, message)
}
