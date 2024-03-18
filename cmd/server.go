package main

import (
	"github.com/cli/internal/constants"
	"github.com/cli/internal/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunAsServer() {
	log := logger.GetLogger()
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(logger.LoggingMiddleware())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} // Replace with your allowed origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config))
	// TODO: DB Initialisation and Controllers Init
	//logController := factories.MakeLogController(db)
	//v1 := r.Group("/api/v1")
	//{
	//	logs := v1.Group("/log")
	//	{
	//		logs.GET("/:path", logController)
	//		// Define other routes`
	//	}
	//}
	port := constants.ServerPort
	log.Info("Server is running on port", port)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
