package main

import "code.google.com/p/go-tour/pic"

// Note that this image will not display in the command line
func Pic(dx, dy int) [][]uint8 {
	x := make([][]uint8, dx)

	for i := range x {
		x[i] = make([]uint8, dy)

		for j := range x[i] {
			x[i][j] = uint8(i) * uint8(j)
		}
	}

	return x
}

func main() {
	pic.Show(Pic)
}
