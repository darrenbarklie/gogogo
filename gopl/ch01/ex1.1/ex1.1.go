// Exercise 1.1 : Modify echo to also print name of
// command that invoked it

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(os.Args[0])
	fmt.Println(s)
}
