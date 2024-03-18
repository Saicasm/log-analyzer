package controllers

import (
	"fmt"
	"github.com/cli/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LogController struct {
	logService services.LogService
}

func NewLogController(LogService services.LogService) *LogController {
	return &LogController{logService: LogService}
}

// CURD controllers
func (logController *LogController) GetLogData(ctx *gin.Context) {
	fmt.Println("log")
	//TODO: Get vars from the requests
	//result, err := logController.logService.getParsedCookie(other vars)

	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get most active cookie"})
	//}

	ctx.JSON(http.StatusOK, "Sample get request")
}
