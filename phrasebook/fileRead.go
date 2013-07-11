package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("fileRead.go")
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	buffer := make([]byte, 100)
	// buffer is 100 bytes so 100 bytes will be read at a time (if there
	// are 100 bytes remaining to be read.)
	for n, e := file.Read(buffer); e == nil; n, e = file.Read(buffer) {
		if n > 0 {
			os.Stdout.Write(buffer[0:n])
		}
	}
}
