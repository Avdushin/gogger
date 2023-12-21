package utils_test

import (
	"github.com/Avdushin/gogger/Examples/App_example/utils"
)

func ExampleHelper() {
	utils.Helper("emenem")

	// Output:
	// 2023/12/21 03:19:03 [WARNING] Emenem
}

func ExampleLoggger() {
	utils.Logger()

	// Output:
	// 2023/12/21 03:28:42 [INFO] Пример лога уровня INFO
	// 2023/12/21 03:28:42 [DEBUG] Пример лога уровня DEBUG
	// 2023/12/21 03:28:42 [WARNING] Пример лога уровня WARNING
	// 2023/12/21 03:28:42 [ERROR] Пример лога уровня ERROR
	// 2023/12/21 03:28:42 [INFO] Just Print INFO message
	// 2023/12/21 03:28:42 [DEBUG] Test message
}
