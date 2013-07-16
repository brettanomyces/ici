package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type webCounter struct {
	count    chan int
	template *template.Template
}

func NewCounter() *webCounter {
	counter := new(webCounter)
	counter.count = make(chan int, 1)
	go func() {
		for i := 1; ; i++ {
			counter.count <- i
		}
	}()

	counter.template, _ = template.ParseFiles("counter.html")
	return counter
}

func (w *webCounter) ServeHTTP(r http.ResponseWriter, rq *http.Request) {
	if rq.URL.Path != "/" {
		r.WriteHeader(http.StatusNotFound)
		return
	}
	// the second argument in the Execute() call is a structure with one
	// field for every value referenced in the template. We define an 
	// anonymous structure and then initialize it with the current counter
	// value.
	w.template.Execute(r, struct{ Counter int }{<-w.count})
}

func main() {
	// Every time a new request is recieved the ServeHTTP() method will be
	// called in a new goroutine. All the gorutines share the same channel.
	err := http.ListenAndServe(":8000", NewCounter())
	if err != nil {
		fmt.Printf("Server failed: ", err.Error())
	}
}
