![Gogger logo](./imgs/logo1.png)

# GOGGER - simple logger for GO projects

## Installation

```
go get -u github.com/Avdushin/gogger/logger
go get -u github.com/Avdushin/gogger/lib
```

## Examples

### Init logger

```go
package main

import "github.com/Avdushin/gogger/logger"

func main() {
	// Инициализация логгера
	log := logger.InitLogger()
	// log.CreateDailyLogFile()

	l := "lol"
	// Пример использования логгера
	log.Info("Пример лога уровня INFO")
	log.Debug("Пример лога уровня DEBUG")
	log.Warning("Пример лога уровня WARNING")
	log.Error("Пример лога уровня ERROR")
	log.Print("Just print the message")
	log.Printf("Just print the message %s", l)
}

```

### Write all logs to file

```go
package main

import "github.com/Avdushin/gogger/logger"

func main() {
	// Запись в файл logs/logs.log
	log := logger.InitLogger("logs/logs.log")

	// Пример использования логгера
	log.Info("Пример лога уровня INFO")
	log.Debug("Пример лога уровня DEBUG")
	log.Warning("Пример лога уровня WARNING")
	log.Error("Пример лога уровня ERROR")
	log.Print("Just Print INFO message")
}
```

### Writing logs to files by day

```go
package main

import (
	"github.com/Avdushin/gogger/logger"
)

func main() {
	// Создаем новый логгер
	log := logger.InitLogger()
	// Лог файлы будут создаваться автоматически в директории "./logs/сегодняя-дата/сегодняя-дата-время/"
	log.CreateDailyLogFile()

	// Пример использования логгера
	log.Info("Пример лога уровня INFO")
	log.Debug("Пример лога уровня DEBUG")
	log.Warning("Пример лога уровня WARNING")
	log.Error("Пример лога уровня ERROR")

	log.Print("Пример лога с уровнем INFO по умолчанию")
	log.Printf("Форматированный лог с уровнем INFO по умолчанию: %s", "some value")
}
```

#### Output
```bash
.
├── dayly_logs.go
└── logs
    └── 2023-12-21
        └── 2023-12-21-002356.log
```

### Custom log formats

```go
package main

import (
	"github.com/Avdushin/gogger/logger"
)

func main() {
	// Создаем новый логгер с уровнем WARNING и записью в файл "custom.log"
	log := logger.InitLogger()

	// Пример использования логгера с пользовательским форматом
	log.SetLogFormat("Custom Format: %s")
	log.Warning("Пример лога с пользовательским форматом")
}
```

### Use Case Examples

|Example | code | repo |
|-|-|-|
|App Example| [watch example](#app-example)|[watch code]("./Examples/App_example") |
|Simple App Example | [watch example](#simple-app-example) |[watch code]("./Examples/SimpleApp")|

---


## App Example

### App Tree
```bash
App_example
├── main.go
└── utils
    ├── helper.go
    └── logger.go
```

### main.go
```go
package main

import (
	"github.com/Avdushin/gogger/Examples/App_example/utils"
)

func main() {
	log := utils.Logger()
	log.Debug("Test message")
	utils.Helper("Emenem")
}
```

### utils/logger.go
```go
package utils

import (
	"github.com/Avdushin/gogger/lib"
	"github.com/Avdushin/gogger/logger"
)

func Logger() *lib.Logger {
	// Запись в файл logs/logs.log
	log := logger.InitLogger("logs/logs.log")

	// Пример использования логгера
	log.Info("Пример лога уровня INFO")
	log.Debug("Пример лога уровня DEBUG")
	log.Warning("Пример лога уровня WARNING")
	log.Error("Пример лога уровня ERROR")
	log.Print("Just Print INFO message")

	return log
}
```

### utils/helper.go
```go
package utils

import "github.com/Avdushin/gogger/logger"

func Helper(name string) {
	logger.Warning(name)
}
```

#### Output
```bash
$ go run main.go 
2023/12/21 00:18:08 [INFO] Пример лога уровня INFO
2023/12/21 00:18:08 [DEBUG] Пример лога уровня DEBUG    
2023/12/21 00:18:08 [WARNING] Пример лога уровня WARNING
2023/12/21 00:18:08 [ERROR] Пример лога уровня ERROR    
2023/12/21 00:18:08 [INFO] Just Print INFO message      
2023/12/21 00:18:08 [DEBUG] Test message
2023/12/21 00:18:08 [WARNING] Emenem
```

### App Output Tree
```bash
App_example
├── logs
│   └── logs.log
├── main.go
└── utils
    ├── helper.go
    └── logger.go
```

---

## Simple App Example

### App Tree
```bash
SimpleApp
├── main.go
└── utils
    └── helper.go
```


```go
package main

import "github.com/Avdushin/gogger/logger"

var (
	libName = "gogger"
)

func main() {
	// Запись в файл logs/logs.log
	log := logger.InitLogger("logs/logs.log")

	// Пример использования логгера
	log.Info("Пример лога уровня INFO")
	log.Debug("Пример лога уровня DEBUG")
	log.Warning("Пример лога уровня WARNING")
	log.Error("Пример лога уровня ERROR")
	log.Print("Just Print INFO message")

	log.Debug("You using %s logger for logging messages in this project!", libName)
}
```

#### utils/helper.go

```go
package utils

import "github.com/Avdushin/gogger/logger"

func Helper(name string) {
	logger.Debug(name)
}
```

#### Output
```bash
$ go run main.go
2023/12/20 23:55:30 [INFO] Пример лога уровня INFO
2023/12/20 23:55:30 [DEBUG] Пример лога уровня DEBUG
2023/12/20 23:55:30 [WARNING] Пример лога уровня WARNING
2023/12/20 23:55:30 [ERROR] Пример лога уровня ERROR
2023/12/20 23:55:30 [INFO] Just Print INFO message
2023/12/20 23:55:30 [DEBUG] You using gogger logger for logging messages in this project!
```

### App Output Tree
```bash
SimpleApp
├── logs
│   └── logs.log
├── main.go
└── utils
    └── helper.go
```

<div align="center">
   <a href="https://github.com/Avdushin">© 2023 Avdushin</a>
</div>