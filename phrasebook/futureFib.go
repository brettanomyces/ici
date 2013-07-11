package main

import (
	"fmt"
	"sync"
	"strconv"
)

type futureInt64 struct {
	ch      chan int64
	v       int64
	collect sync.Once
}

func (f *futureInt64) String() string {
	f.collect.Do(func() { f.v = <-f.ch })
	return strconv.FormatInt(f.v, 10)
}

func fib(n int64) (int64, int64) {
	if n < 2 {
		return 1, 1
	}
	f1, f2 := fib(n - 1)
	return f2, f1 + f2

}

func Fib(n int64) fmt.Stringer {
	var ch futureInt64
	ch.ch = make(chan int64)
	go func() {
		_, f := fib(n)
		ch.ch <- f
	}()
	return &ch
}

func main() {
	// f is a fmt.Stringer, its String() method blocks until there is an
	// int64 in the channel.
	f := Fib(100)
	fmt.Printf("the 100th Fibonacci number is: ")
	fmt.Printf("%v\n", f)
}
