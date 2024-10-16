package utils

import (
	"fmt"
	"os"
)

func DebugLog(format string, a ...interface{}) {
	if os.Args[len(os.Args)-1] == "debug" {
		fmt.Printf(format, a...)
	}
}
