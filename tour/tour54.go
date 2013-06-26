package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// By implementing Error() a pointer to MyError implicitly implements the error
// interface
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// Which allows it to be returned as an error
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
