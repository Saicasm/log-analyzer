package main

import (
	"github.com/cli/internal/constants"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	godotenv.Load()
}
func main() {
	mode := os.Getenv(constants.LOG_MODE)
	if mode == constants.CLI {
		RunAsCli()
	} else if mode == constants.SERVER {
		RunAsServer()
	} else {
		RunAsCli()
		RunAsServer()
	}
}
