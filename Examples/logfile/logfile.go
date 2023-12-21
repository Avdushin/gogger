package main

import "github.com/Avdushin/gogger/logger"

func main() {
	// Запись в файл logs/logs.log
	log := logger.InitLogger("logs/logs.log")

	// Пример использования логгера
	log.Info("Пример лога уровня INFO")
	log.Debug("Пример лога уровня DEBUG")
	log.Warning("Пример лога уровня WARNING")
	log.Error("Пример лога уровня ERROR")
	log.Print("Just Print INFO message")

	// Output:
	// 2023/12/21 03:08:59 [INFO] Пример лога уровня INFO
	// 2023/12/21 03:08:59 [DEBUG] Пример лога уровня DEBUG
	// 2023/12/21 03:08:59 [WARNING] Пример лога уровня WARNING
	// 2023/12/21 03:08:59 [ERROR] Пример лога уровня ERROR
	// 2023/12/21 03:08:59 [INFO] Just Print INFO message
	//
	// Files Tree:
	// .
	// ├── dayly_logs.go
	// └── logs
	//     └── logs.log

}
