package lib

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// @ New создает новый экземпляр Logger
// @ Опциональные параметры:
// @ logLevel - уровень логирования
// @ filePath - путь к файлу, в который записывать логи
func New(args ...interface{}) *Logger {
	l := &Logger{}

	for _, arg := range args {
		switch v := arg.(type) {
		case LogLevel:
			// ? Если передан уровень логирования, устанавливаем его
			l.logLevel = v
		case string:
			// В любом случае создаем логгер с указанным путем
			err := createLogFile(v)
			if err != nil {
				log.Fatalf("Ошибка при создании файла лога: %v", err)
			}

			file, err := os.OpenFile(v, os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				log.Fatalf("Ошибка при открытии файла лога: %v", err)
			}

			// Устанавливаем файл в качестве вывода для стандартного логгера
			l.fileWriter = file
		default:
			log.Fatalf("Неподдерживаемый тип аргумента: %T", v)
		}
	}

	return l
}

// @ createLogFile создает файл лога и его директорию, если они не существуют
func createLogFile(filePath string) error {
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Ошибка при создании директории: %v", err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Ошибка при создании файла лога: %v", err)
	}
	file.Close()

	return nil
}
