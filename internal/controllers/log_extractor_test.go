package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockLogService is a mock implementation of services.LogService
type MockLogService struct{}

func (m *MockLogService) GetLogData(vars string) (string, error) {
	return "Mocked data", nil
}

func TestGetLogData(t *testing.T) {
	// Setup Gin router
	r := gin.Default()

	// Initialize mock log service
	mockLogService := &MockLogService{}

	// Initialize LogController with mock log service
	logController := NewLogController(mockLogService)

	// Define route handler
	r.GET("/logs", logController.GetLogData)

	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "api/v1/logs/path", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	recorder := httptest.NewRecorder()

	// Serve the HTTP request to the recorder
	r.ServeHTTP(recorder, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Check the response body
	assert.Equal(t, "Sample get request", recorder.Body.String())
}
