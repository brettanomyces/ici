package main

import (
	"fmt"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	oldZ := 1.0
	newZ := 1.0
	delta := 0.000001
	for {
		newZ = oldZ - ((oldZ*oldZ - x) / (2 * oldZ))
		if newZ-oldZ < delta && newZ-oldZ > -delta {
			return newZ, nil
		} else {
			oldZ = newZ
		}
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %g", e)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
