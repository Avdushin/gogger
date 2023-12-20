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
}
