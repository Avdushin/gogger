// Tis package contains the utility functions
// Helper - a function that printing a Warnings message
// logger - a function that setups the logger
package utils

import "github.com/Avdushin/gogger/logger"

func Helper(name string) {
	logger.Warning(name)
}
