package lib

import (
	"fmt"
	"os"
)

// @ SetLogLevel устанавливает уровень логирования
func (l *Logger) SetLogLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logLevel = level
}

// @ SetLogFormat устанавливает формат лога
func (l *Logger) SetLogFormat(format string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logFormat = format
}

// @ SetLogFile устанавливает файл для записи логов
func (l *Logger) SetLogFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Ошибка при создании файла лога: %v", err)
	}

	l.fileWriter = file
	return nil
}
