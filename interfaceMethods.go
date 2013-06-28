package main

import "fmt"

type changee struct {
	val string
}

type changer interface {
	change()
}

// This method will recieve a pointer to changee regarless of if the method is
// called on a '*changee' or a 'changee'. However, only '*changee' implements
// the changer interface. This is just a regular function, it is not tied to 
// the interface.
func (ptrToChangee *changee) change() {
	ptrToChangee.val += " has been changed"
}

// Note change() cannot be declared for both '*changee' and 'changee'

// a '*changee' can be assigned to x but a 'changee' cannot
var x changer

func main() {
	c := changee{"c"}
	c.change()
	fmt.Println(c)

	// The following is ILLEGAL because 'changee' does not implement the
	// changer interface
	//x = changee{"x"}
	//fmt.Println(x)

	// The following is legal becase '*changee' implements the changer
	// interface
	x = &(changee{"x"})
	fmt.Println(x)
}
