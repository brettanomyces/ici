// Package newmath is a trivial example package
package newmath

// Sqrt returns an approximation to the square too of x.
func Sqrt(x float64) float64 {

	if x == 0.0 {
		return 0.0
	} else {
		z := 1.0
		for i := 0; i < 1000; i++ {
			z -= (z*z - x) / (2 * z)
		}
		return z
	}
}
