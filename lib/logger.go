// The lib package contains core functionality of the Gogger
package lib

import (
	"fmt"
	"io"
	"log"
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

	// Вывод в консоль с использованием цветов
	if color, ok := LogLevelColor[level]; ok {
		fullMsg = color.Sprint(fullMsg)
	}
	log.Println(fullMsg)

	if l.fileWriter != nil {
		fullMsg = fmt.Sprintf("%s %s", time.Now().Format("2006-01-02 15:04:05"), fullMsg)
		fmt.Fprintln(l.fileWriter, fullMsg)
	}
}
