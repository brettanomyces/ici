package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	y := 1.0
	z := 1.0
	n := 0
	diff := 0.0000001
	for {
		y = z - ((z*z - x) / (2 * z))
		if y-z > -diff && y-z < diff {
			fmt.Println("Iterations:", n)
			return y
		} else {
			z = y
		}
		n++
	}
}

func main() {
	n := float64(100)
	fmt.Println(Sqrt(n))
	fmt.Println(math.Abs(Sqrt(n) - math.Sqrt(n)))
}
