// Tis package contains the utility functions
// Helper - a function that printing a Warnings message
// logger - a function that setups the logger
package utils

import (
	"github.com/Avdushin/gogger/lib"
	"github.com/Avdushin/gogger/logger"
)

func Logger() *lib.Logger {
	// Запись в файл logs/logs.log
	log := logger.InitLogger("logs/logs.log")

	// Пример использования логгера
	log.Info("Пример лога уровня INFO")
	log.Debug("Пример лога уровня DEBUG")
	log.Warning("Пример лога уровня WARNING")
	log.Error("Пример лога уровня ERROR")
	log.Print("Just Print INFO message")

	return log
}
