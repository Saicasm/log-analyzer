package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"time"

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
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		path := c.Request.URL.Path

		fmt.Printf("%s | %3d | %13v | %15s | %-7s %#v\n",
			end.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
