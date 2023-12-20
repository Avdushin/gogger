package lib

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
