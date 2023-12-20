package main

import (
	"github.com/Avdushin/gogger/Examples/App_example/utils"
)

func main() {
	log := utils.Logger()
	log.Debug("Test message")
	utils.Helper("Emenem")
}
