package main

import (
	"github.com/Avdushin/gogger/logger"
)

func main() {
	// Создаем новый логгер
	log := logger.InitLogger()
	// Лог файлы будут создаваться автоматически в директории "./logs/сегодняя-дата/сегодняя-дата-время/"
	log.CreateDailyLogFile()

	// Пример использования логгера
	log.Info("Пример лога уровня INFO")
	log.Debug("Пример лога уровня DEBUG")
	log.Warning("Пример лога уровня WARNING")
	log.Error("Пример лога уровня ERROR")

	log.Print("Пример лога с уровнем INFO по умолчанию")
	log.Printf("Форматированный лог с уровнем INFO по умолчанию: %s", "some value")
}
