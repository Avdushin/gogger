package main

import (
	"github.com/Avdushin/gogger/Examples/App_example/helpers/logger"
)

func main() {
	logger.Info("App started...")
	logger.Debug("App started...")
	logger.Warning("App started...")
	logger.Error("App started...")
}

func init() {
	logger.InitLogger()
}
