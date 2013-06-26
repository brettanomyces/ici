package main

import (
	"fmt"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	y := complex128(0)
	z := complex128(1)

	for y != z {
		y = z
		z = z - (cmplx.Pow(z, 3)-x)/(3*cmplx.Pow(z, 2))
	}
	return z
}

func main() {
	fmt.Println(Cbrt(complex128(1i)))
}
