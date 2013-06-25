package newmath

import "testing"

func TestSqrt(t *testing.T) {
	const in, out = 4, 2
	if x := Sqrt(in); x != out {
		t.Errorf("Sqrt(%v) = %v, expected %v", in, x, out)
	}
}

func TestZero(t *testing.T) {
	const zero = 0
	if x := Sqrt(zero); x != zero {
		t.Errorf("Sqrt(%v) - %v, expected %v", zero, x, zero)
	}
}
