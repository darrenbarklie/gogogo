// Exercise 1.2 : Modify echo to print the index and value of each
// argument, one per line

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}
