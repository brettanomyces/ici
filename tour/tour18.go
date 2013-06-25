package main

import "fmt"

func main() {
	sum := 1
	// Go has no while statement. for is used instead.
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
