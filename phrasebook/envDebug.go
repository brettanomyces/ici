package main

import (
	"fmt"
)

func debugLog(level int, msg string, args ...interface{}) {
	if debugLevel > level {
		fmt.Printf(msg, args...)
	}
}

func main() {
	debugLevel, _ = strconv.Atoi(os.Getenv("DEBUG"))
	debugLog(3, "Starting\n")
}
