package main

import (
	"github.com/cli/internal/constants"
	"github.com/cli/internal/logger"
	"github.com/cli/internal/router"
	"github.com/gin-gonic/gin"
)

func RunAsServer() {
	log := logger.GetLogger()
	r := gin.Default()
	router.InitRoutes(r)
	port := constants.ServerPort
	log.Info("Server is running on port", port)
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
