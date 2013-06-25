package main

import "fmt"

func main() {
	sum := 1
	// You can leave the pre and post statements empty, 
	// e.g.: for ; sum < 1000; {
	// gofmt will simply remove them though
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
