package logger

import (
	"fmt"

	"github.com/Avdushin/gogger/lib"
)

// @ Инициализация логгера
func InitLogger(args ...interface{}) *lib.Logger {
	var logFilePath string

	if len(args) > 0 {
		// Первый аргумент - строка, путь к файлу логов
		if v, ok := args[0].(string); ok {
			logFilePath = v
		}
	}

	Log = lib.New(args...)
	if logFilePath != "" {
		// Если путь к файлу не пустой, создаем файл для записи логов
		err := Log.SetLogFile(logFilePath)
		if err != nil {
			// Обработка ошибки, например, вывод в консоль
			fmt.Println("Ошибка при установке файла логов:", err)
		}
	}

	return Log
}
