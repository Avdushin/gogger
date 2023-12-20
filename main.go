package main

import "gogger/logger"

func main() {
	// Создаем новый логгер с уровнем INFO и записью в файл "app.log"
	log := logger.New(logger.INFO)

	// Пример использования логгера
	log.Info("Пример лога уровня INFO")
	log.Debug("Пример лога уровня DEBUG") // Этот лог не будет записан из-за уровня логирования INFO
	log.Warning("Пример лога уровня WARNING")
	log.Error("Пример лога уровня ERROR")

	// Меняем уровень логирования на DEBUG
	log.SetLogLevel(logger.DEBUG)

	// Теперь лог DEBUG будет записан
	log.Debug("Теперь этот лог будет записан")

	log.Print("Пример лога с уровнем INFO по умолчанию")
	log.Printf("Форматированный лог с уровнем INFO по умолчанию: %s", "some value")

}
