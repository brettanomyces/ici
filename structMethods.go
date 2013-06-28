package main

import "fmt"

type changee struct {
	val string
}

// This change will recieve a POINTER to the changee regardless of whether
// it is called on a '*changee' or a 'changee'. Changes made in this func
// are visible to the caller.
func (ptrToChangee *changee) change1() {
	ptrToChangee.val += " has been changed"
}

// This change will recieve a COPY of the changee regardless of whether it is
// called on a '*changee' or a 'changee'. Changes made will not be visible
// to the caller.
func (copyOfChangee changee) change2() {
	copyOfChangee.val += " has been changed"
}

// Note that a method with the same name cannot be defined for both 'changee'
// and '*changee'

func main() {
	s := changee{"s"}
	s.change1()
	fmt.Println(s)

	s2 := &(changee{"s2"})
	s2.change1()
	fmt.Println(*s2)

	s3 := changee{"s3"}
	s3.change2()
	fmt.Println(s3)

	s4 := &(changee{"s4"})
	s4.change2()
	fmt.Println(*s4)
}
