package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogService(t *testing.T) {
	// Call the constructor function
	logService := NewLogService()

	// Assert that the returned LogService is not nil
	assert.NotNil(t, logService)
}
