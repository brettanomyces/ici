package main

import (
	"fmt"
	"time"
)

func main() {
	abort := make(chan bool)
	count := make(chan int)

	go cancel(abort)
	go countdown(count)
	for {
		select {
		case i := <-count:
			if i == 0 {
				selfDestruct()
				return
			}
			fmt.Printf("%d seconds remaining\n", i)
		case a := <-abort:
			if a {
				fmt.Printf("Self destruct aborted\n")
			} else {
				selfDestruct()
			}
			return
		}
	}
}

func cancel(abort chan bool) {
	fmt.Printf("This program will self destruct, do you wish to cancel?\n")
	var r int
	fmt.Scanf("%c", &r)
	switch r {
	default:
		abort <- false
	case 'Y':
		abort <- true
	case 'y':
		abort <- true
	}
}

func countdown(count chan int) {
	for i := 10; i >= 0; i-- {
		count <- i
		time.Sleep(1000000000)
	}
}

func selfDestruct() {
	fmt.Printf("BOOM!\n")
}
