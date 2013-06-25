package main

import "fmt"

var x, y, z int = 1, 2, 3
var c, python, java = true, false, "no!"

// Notice the variables c, python and java simply take the type of their 
// initializer
func main() {
	fmt.Println(x, y, z, c, python, java)
}
