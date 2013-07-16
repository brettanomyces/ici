package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	components := []string{"a", "path", "..", "with", "relative", "elements"}
	// "path/../" will be removed from the path since it is irrelevant.
	// the ... after components indicates that componets should be passed 
	// as the variadic arguments set, rather than a single argument.
	path := path.Join(components...)
	fmt.Printf("Path: %s\n", path)
	decomposed := filepath.SplitList(path)
	for _, dir := range decomposed {
		fmt.Printf("%s%c", dir, filepath.Separator)
	}
	fmt.Printf("\n")
}
