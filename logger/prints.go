package logger

import "fmt"

// Print выводит лог с уровнем по умолчанию [INFO]
func (l *Logger) Print(args ...interface{}) {
	l.log(INFO, fmt.Sprint(args...))
}

// Printf выводит лог с уровнем по умолчанию [INFO] и форматирует строку
func (l *Logger) Printf(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}
