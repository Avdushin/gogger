package main

import (
	"gogger/logger"
)

func main() {
	// Создаем новый логгер с директорией, в которой будут храниться файлы логов
	log := logger.New("logs/logs.log")

	// Пример использования логгера
	log.Info("Пример лога уровня INFO")
	log.Debug("Пример лога уровня DEBUG")
	log.Warning("Пример лога уровня WARNING")
	log.Error("Пример лога уровня ERROR")
}
