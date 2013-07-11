package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
		file, err := os.Open("lineRead.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	lineReader := bufio.NewReaderSize(file, 20)
	lineNumber := 0

	// Readline returns 3 values, the slice,  a flag
	// indicationg whether the slice represents a partial line, and the 
	// error code.
	for line, isPrefix, e := lineReader.ReadLine(); e == nil; line, isPrefix, e = lineReader.ReadLine() {

		fmt.Printf("%.3d: ", lineNumber)
		lineNumber++
		os.Stdout.Write(line)
		if isPrefix {
			for {
				line, isPrefix, _ = lineReader.ReadLine()
				os.Stdout.Write(line)
				if !isPrefix {
					break
				}
			}
		}
		fmt.Printf("\n")
	}
}
