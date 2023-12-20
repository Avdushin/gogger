package main

import (
	"gogger/logger"
)

func main() {
	// Создаем новый логгер с уровнем WARNING и записью в файл "custom.log"
	log := logger.New(logger.WARNING)

	// Пример использования логгера с пользовательским форматом
	log.SetLogFormat("Custom Format: %s")
	log.Warning("Пример лога с пользовательским форматом")
}
