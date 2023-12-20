package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// @ CreateDailyLogFile создает файл лога в отдельной папке с именем текущей даты
func (l *Logger) CreateDailyLogFile() (string, error) {
	//? Определяем текущую дату
	now := time.Now()
	dateFolder := now.Format("2006-01-02")

	//? Создаем папку для текущей даты
	dir := filepath.Join("logs", dateFolder)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("Ошибка при создании директории: %v", err)
	}

	//? Создаем имя файла лога внутри папки с текущей датой
	fileName := fmt.Sprintf("%s.log", now.Format("2006-01-02-150405"))

	//? Полный путь к файлу
	filePath := filepath.Join(dir, fileName)

	//? Создаем сам файл
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("Ошибка при создании файла лога: %v", err)
	}

	//? Сохраняем путь к файлу в структуре Logger
	l.fileWriter = file

	return filePath, nil
}
