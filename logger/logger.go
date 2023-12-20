package logger

import (
	"fmt"
	"log"
	"os"
	"sync"

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
}

// NewLogger создает новый экземпляр Logger с заданным уровнем логирования и опциональным путем к файлу
func New(level LogLevel, filePath ...string) *Logger {
	l := &Logger{
		logLevel: level,
	}

	if len(filePath) > 0 && filePath[0] != "" {
		file, err := os.OpenFile(filePath[0], os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Ошибка при открытии файла лога: %v", err)
		}
		l.fileLogger = log.New(file, "", log.Ldate|log.Ltime)
	}

	return l
}

// SetLogLevel устанавливает уровень логирования
func (l *Logger) SetLogLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logLevel = level
}

// log записывает лог в консоль и/или файл в зависимости от настроек
func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	if level < l.logLevel {
		return
	}

	msg := fmt.Sprintf(format, args...)
	fullMsg := fmt.Sprintf("[%s] %s", levelToString(level), msg)

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

	// Восстановление цвета после вывода
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

// levelToString преобразует уровень логирования в строку
func levelToString(level LogLevel) string {
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

// Print выводит лог с уровнем по умолчанию (INFO)
func (l *Logger) Print(args ...interface{}) {
	l.log(INFO, fmt.Sprint(args...))
}

// Printf выводит лог с уровнем по умолчанию (INFO) и форматирует строку
func (l *Logger) Printf(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}
