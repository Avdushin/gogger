# GOGGER - simple logger for GO projects

## Installation

```
go get -u github.com/Avdushin/gogger/logger
```

## Examples

### Init logger

```go
package main

import (
	"github.com/Avdushin/gogger/logger"
)

func main() {
	// Запись в файл logs/logs.log
	log := logger.New("logs/logs.log")
	log.CreateDailyLogFile()

	// Пример использования логгера
	log.Info("Пример лога уровня INFO")
	log.Debug("Пример лога уровня DEBUG")
	log.Warning("Пример лога уровня WARNING")
	log.Error("Пример лога уровня ERROR")
	log.Print("Just Print INFO message")
}

```

## Use Case

```bash
./EXAMPLES/APP_EXAMPLE     
│   main.go
│   
└───helpers
    └───logger
            logger.go
```

#### main.go

```go
package main

import (
	"github.com/Avdushin/gogger/Examples/App_example/helpers/logger"
)

func main() {
	logger.Info("App started...")
	logger.Debug("App started...")
	logger.Warning("App started...")
	logger.Error("App started...")
}

func init() {
	logger.InitLogger()
}

```


####

```go
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
```