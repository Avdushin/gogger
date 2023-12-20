package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// @ Уровни логирования
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
)

// @ Logger - основная структура логгера
type Logger struct {
	mu         sync.Mutex
	fileWriter io.Writer
	logLevel   LogLevel
	logFormat  string
}

// @ CreateDailyLogFile создает файл лога в отдельной папке с именем текущей даты
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
	l.fileWriter = file

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

// @ log записывает лог в консоль и/или файл в зависимости от настроек
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

	// Вывод в консоль с использованием стандартного log
	log.Println(fullMsg)

	if l.fileWriter != nil {
		// Добавляем временную метку и записываем в файл
		fullMsg = fmt.Sprintf("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), fullMsg)
		fmt.Fprintln(l.fileWriter, fullMsg)
	}
}

func (l *Logger) Write(p []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.fileWriter.Write(p)
}
