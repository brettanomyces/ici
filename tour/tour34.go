package main

import "fmt"

func main() {
	pow := make([]int, 10)

	// Here we ignore the value given by range, this could be written as:
	// for i, _ := range pow {...}
	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	// Here we ignore the 'i' given by range by assigning it to _
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
