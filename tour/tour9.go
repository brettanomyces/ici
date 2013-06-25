package main

import "fmt"

// This func utilises named parameters, notice the return statement has no
// arguments. A return statement with no arguments returns the current values
// of the named parameters. However, arguments can still be given in which case
// it will return the given arguments.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return 4, 5
}

func main() {
	fmt.Println(split(17))
}
