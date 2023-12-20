package logger

import (
	"github.com/Avdushin/gogger/logger"
)

// ? Log - глобальная переменная логгера
var Log *logger.Logger

// @ Инициализация логгера
func InitLogger() *logger.Logger {
	//@ Create logger
	Log = logger.New()
	//@ Create Daily Log files
	// Log.CreateDailyLogFile()

	return Log
}

// Debug записывает лог уровня DEBUG
func Debug(format string, args ...interface{}) {
	Log.Debug(format, args...)
}

// Info записывает лог уровня INFO
func Info(format string, args ...interface{}) {
	Log.Info(format, args...)
}

// Warning записывает лог уровня WARNING
func Warning(format string, args ...interface{}) {
	Log.Warning(format, args...)
}

// Error записывает лог уровня ERROR
func Error(format string, args ...interface{}) {
	Log.Error(format, args...)
}
