package main

import (
	"gogger/logger"
)

func main() {
	// Активировать запись в файлы по уровням
	log := logger.New("logs/logs.log")
	log.CreateDailyLogFile()

	// Пример использования логгера
	log.Info("Пример лога уровня INFO")
	log.Debug("Пример лога уровня DEBUG")
	log.Warning("Пример лога уровня WARNING")
	log.Error("Пример лога уровня ERROR")
	log.Print("lflflf")
}