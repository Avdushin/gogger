package main

import "github.com/Avdushin/gogger/logger"

func main() {
	// Инициализация логгера
	log := logger.InitLogger()
	// log.CreateDailyLogFile()

	l := "lol"
	// Пример использования логгера
	log.Info("Пример лога уровня INFO")
	log.Debug("Пример лога уровня DEBUG")
	log.Warning("Пример лога уровня WARNING")
	log.Error("Пример лога уровня ERROR")
	log.Print("Just print the message")
	log.Printf("Just print the message %s", l)
}
