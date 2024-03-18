package main

import (
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}
func main() {
	// TODO:commenting the code as only CLI mode is needed now
	//mode := os.Getenv(constants.LOG_MODE)
	//if mode == constants.CLI {
	//	RunAsCli()
	//} else if mode == constants.SERVER {
	//	RunAsServer()
	//} else {
	//	RunAsCli()
	//	RunAsServer()
	//}
	RunAsCli()
}
