package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["answer"] = 42
	fmt.Println("The value:", m["answer"])

	m["answer"] = 48
	fmt.Println("The value:", m["answer"])

	delete(m, "answer")
	fmt.Println("The value:", m["answer"])

	// If the value is not present the the return value for v is the zero 
	// value for the map's element type, in this case 0
	v, ok := m["answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
