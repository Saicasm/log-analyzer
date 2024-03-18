package factories

import (
	"github.com/cli/internal/controllers"
	"github.com/cli/internal/services"
)

func MakeLogController() *controllers.LogController {
	logService := services.NewLogService()
	return controllers.NewLogController(logService)
}
