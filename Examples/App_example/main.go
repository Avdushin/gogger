package main

import (
	"github.com/Avdushin/gogger/Examples/App_example/utils"
)

func main() {
	// Init logger
	log := utils.Logger()
	// Print Debug log
	log.Debug("Test message")
	// Call Helper function from utils
	utils.Helper("Emenem")

	// Output:
	// 2023/12/21 03:19:03 [INFO] Пример лога уровня INFO
	// 2023/12/21 03:19:03 [DEBUG] Пример лога уровня DEBUG
	// 2023/12/21 03:19:03 [WARNING] Пример лога уровня WARNING
	// 2023/12/21 03:19:03 [ERROR] Пример лога уровня ERROR
	// 2023/12/21 03:19:03 [INFO] Just Print INFO message
	// 2023/12/21 03:19:03 [DEBUG] Test message
	// 2023/12/21 03:19:03 [WARNING] Emenem
}
