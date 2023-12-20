package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fatih/color"
)

// Уровни логирования
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

// Logger - основная структура логгера
type Logger struct {
	mu         sync.Mutex
	logLevel   LogLevel
	filePath   string
	fileLogger *log.Logger
	logFormat  string
}

// CreateDailyLogFile создает файл лога в отдельной папке с именем текущей даты
func (l *Logger) CreateDailyLogFile() (string, error) {
	// Определяем текущую дату
	now := time.Now()
	dateFolder := now.Format("2006-01-02")

	// Создаем папку для текущей даты
	dir := filepath.Join("logs", dateFolder)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("Ошибка при создании директории: %v", err)
	}

	// Создаем имя файла лога внутри папки с текущей датой
	fileName := fmt.Sprintf("%s.log", now.Format("2006-01-02-150405"))

	// Полный путь к файлу
	filePath := filepath.Join(dir, fileName)

	// Создаем сам файл
	file, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("Ошибка при создании файла лога: %v", err)
	}

	// Сохраняем путь к файлу в структуре Logger
	l.filePath = filePath
	l.fileLogger = log.New(file, "", log.Ldate|log.Ltime)

	return filePath, nil
}

// @ New создает новый экземпляр Logger
// @ Опциональные параметры:
// @ logLevel - уровень логирования
// @ filePath - путь к файлу, в который записывать логи
func New(args ...interface{}) *Logger {
	l := &Logger{}

	for _, arg := range args {
		switch v := arg.(type) {
		case LogLevel:
			// Если передан уровень логирования, устанавливаем его
			l.logLevel = v
		case string:
			// Если передан путь к файлу, создаем файл лога
			if v == "" {
				// Если путь не задан, создаем логи по дням
				v, _ = l.CreateDailyLogFile()
			}

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
			l.fileLogger = log.New(file, "", log.Ldate|log.Ltime)

			l.filePath = v
		default:
			log.Fatalf("Неподдерживаемый тип аргумента: %T", v)
		}
	}

	return l
}

// createLogFile создает файл лога и его директорию, если они не существуют
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

// SetLogLevel устанавливает уровень логирования
func (l *Logger) SetLogLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logLevel = level
}

// SetLogFormat устанавливает формат лога
func (l *Logger) SetLogFormat(format string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logFormat = format
}

// log записывает лог в консоль и/или файл в зависимости от настроек
func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	if level < l.logLevel {
		return
	}

	msg := fmt.Sprintf(format, args...)
	fullMsg := fmt.Sprintf("[%s] %s", l.logLevelToString(level), msg)

	// Добавляем формат лога, если он задан
	if l.logFormat != "" {
		fullMsg = fmt.Sprintf(l.logFormat, fullMsg)
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.fileLogger != nil {
		l.fileLogger.Println(fullMsg)
	}

	// Выбор цвета в зависимости от уровня логирования
	switch level {
	case INFO:
		color.Set(color.FgWhite)
	case DEBUG:
		color.Set(color.FgGreen)
	case WARNING:
		color.Set(color.FgYellow)
	case ERROR:
		color.Set(color.FgRed)
	}

	defer color.Unset()

	log.Println(fullMsg)
}

// Debug записывает лог уровня DEBUG
func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

// Info записывает лог уровня INFO
func (l *Logger) Info(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

// Warning записывает лог уровня WARNING
func (l *Logger) Warning(format string, args ...interface{}) {
	l.log(WARNING, format, args...)
}

// Error записывает лог уровня ERROR
func (l *Logger) Error(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

// logLevelToString преобразует уровень логирования в строку
func (l *Logger) logLevelToString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return "INFO"
	}
}

func (l *Logger) Write(p []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.fileLogger.Print(string(p))
	return len(p), nil
}
