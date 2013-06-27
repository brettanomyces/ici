package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same deterimes wheter the trees t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for v1 := range ch1 {
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println()
	ch := make(chan int)

	// the channel can not be closed in main so stick it in an anonomous func?
	// This does mean they start traversing straight away...
	go func() {
		Walk(tree.New(1), ch)
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))

}
