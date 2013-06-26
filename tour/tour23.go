package main

import (
	"fmt"
	"math"
)

func Sqrt(x, delta float64) float64 {
	fmt.Println("My implementation")
	oldZ := 1.0
	newZ := 1.0
	for {
		newZ = oldZ - ((oldZ*oldZ - x) / (2 * oldZ))
		if newZ-oldZ > -delta && newZ-oldZ < delta {
			return newZ
		} else {
			oldZ = newZ
		}
	}
}

func main() {
	fmt.Println(Sqrt(100, 0.001))
	fmt.Println(math.Sqrt(100))
}
