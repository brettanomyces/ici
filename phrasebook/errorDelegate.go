package main

import (
	"fmt"
	"math"
)

// This interface is an error delegate, it has one method which implements
// some recovery behaviour
type sqrtError interface {
	invalidArgument(int) (int, error)
}

func sqrt(i int, e sqrtError) (result float64, err error) {
	// e.invalidArgument(i) will keep being called until i >= 0, or another
	// error occurs
	for i < 0 {
		var err error
		i, err = e.invalidArgument(i)
		if err != nil {
			return 0, err
		}
	}
	return math.Sqrt(float64(i)), nil
}

type sqrtHandler struct{}

// sqrtHandler implements the sqrtError interface
func (_ sqrtHandler) invalidArgument(i int) (int, error) {
	fmt.Printf("%d is not valid, please enter another value\n", i)
	fmt.Scanf("%d", &i)
	return i, nil
}

func main() {
	fmt.Printf("Enter a number\n")
	var i int
	fmt.Scanf("%d", &i)
	root, err := sqrt(i, sqrtHandler{})
	if err == nil {
		fmt.Printf("sqrt(%d) = %f\n", i, root)
	} else {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
