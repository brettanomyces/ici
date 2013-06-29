package main

import "fmt"
import "sync"

func main() {
	m := make(map[int]string)
	m[2] = "First Value"
	var mutex sync.Mutex

	cv := sync.NewCond(&mutex)
	updateCompleted := false
	go func() {
		cv.L.Lock()
		m[2] = "Second Value"
		updateCompleted = true
		cv.Signal()
		// when signal or broadcast is called the sleeping goroutines(s)
		// wake up and attempt to reaquire the lock
		cv.L.Unlock()
	}()

	cv.L.Lock()
	for !updateCompleted {
		cv.Wait()
		// when we call wait the lock is automatically released and the
		// goroutine goes to sleep.
	}

	v := m[2]
	cv.L.Unlock()
	fmt.Printf("%s\n", v)
}
