package main

import "fmt"

const (
	Red   = (1 << iota)
	Green = (2 << iota)
	Blue  = (3 << iota)
)

const (
	True  = iota
	False = iota
)

func main() {
	fmt.Println(Red)
	fmt.Println(Green)
	fmt.Println(Blue)
	fmt.Println(True)
	fmt.Println(False)
}
