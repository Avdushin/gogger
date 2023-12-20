package lib

import "github.com/fatih/color"

// @ Цвет для каждого уровня логирования
var LogLevelColor = map[LogLevel]*color.Color{
	DEBUG:   color.New(color.FgBlue),
	INFO:    color.New(color.FgGreen),
	WARNING: color.New(color.FgYellow),
	ERROR:   color.New(color.FgRed),
}
