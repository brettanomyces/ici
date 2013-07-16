package main

import (
	"fmt"
	"net/http"
)

type webCounter struct {
	count chan int
}

func NewCounter() *webCounter {
	counter := new(webCounter)
	counter.count = make(chan int, 1)
	// Because the channel size is only 1 the following function blocks
	// until something recieves the value currently in the channel
	go func() {
		for i := 1; ; i++ {
			counter.count <- i
		}
	}()
	return counter
}

func (w *webCounter) ServeHTTP(r http.ResponseWriter, rq *http.Request) {
	if rq.URL.Path != "/" {
		r.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprintf(r, "you are visitor %d", <-w.count)
}

func main() {
	// Every time a new request is recieved the ServeHTTP() method will be
	// called in a new goroutine. All the gorutines share the same channel.
	err := http.ListenAndServe(":8000", NewCounter())
	if err != nil {
		fmt.Printf("Server failed: ", err.Error())
	}
}
