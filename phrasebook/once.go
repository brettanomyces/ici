package main

import (
	"fmt"
	"sync"
)

type LazyInit struct {
	once  sync.Once
	value int
}

func (s *LazyInit) Value() int {
	s.init()
	return s.value
}

func (s *LazyInit) init() {
	// this will run once and only once, it is thread safe
	s.once.Do(func() { s.value = 42 })
}

func (s *LazyInit) SetValue(v int) {
	s.value = v
}

func main() {
	var l LazyInit
	// Here we acess value directly, this would not be able to hapen outside
	// this package
	fmt.Printf("%d\n", l.value)
	// Here we access value throught the Value() method. init will be run.
	fmt.Printf("%d\n", l.Value())
	l.SetValue(12)
	// again we access value throught the value() method. This time thought
	// init will be called but not do anything.
	fmt.Printf("%d\n", l.Value())
}
