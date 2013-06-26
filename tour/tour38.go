package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	// Notice we don't need to include 'Vertex' within.
	// Notice also that these comments don't stop this program from compiling
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -123.08408},
}

func main() {
	fmt.Println(m)
}
