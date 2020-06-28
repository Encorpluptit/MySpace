package debug

import (
	"fmt"
	"log"
	"runtime"
)

// Debug prints a debug information to the log with file and line.
func Debug(format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	info := fmt.Sprintf(format, a...)

	log.Printf("[cgl] debug %s:%d %v", file, line, info)
}
