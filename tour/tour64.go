package main

import "fmt"

func main() {
	// here we make a buffered channel with buffer lenght 2
	c := make(chan int, 2)
	c <- 1
	c <- 2
	//c <- 3
	fmt.Println(<-c)
	fmt.Println(<-c)
}
