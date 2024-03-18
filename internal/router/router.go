package router

import (
	"github.com/cli/internal/factories"
	"github.com/cli/internal/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(logger.LoggingMiddleware())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Replace with your allowed origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))
	logController := factories.MakeLogController()
	v1 := r.Group("/api/v1")
	{
		logs := v1.Group("/log")
		{
			logs.GET("/path", logController.GetLogData)
			// Define other routes`
		}
	}
}
