package main

import (
	"github.com/Avdushin/gogger/logger"
)

func main() {
	// Создаем новый логгер с уровнем WARNING и записью в файл "custom.log"
	log := logger.InitLogger()

	// Пример использования логгера с пользовательским форматом
	log.SetLogFormat("Custom Format: %s")
	log.Warning("Пример лога с пользовательским форматом")
}
