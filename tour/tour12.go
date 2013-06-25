package main

import "fmt"

// inside a function := can be used in place of a var declaration. Outside of a
// function, every construct begins with a keyword so := cannot be used.
func main() {
	var x, y, z int = 1, 2, 3
	c, python, java := true, false, "no!"

	fmt.Println(x, y, z, c, python, java)
}
