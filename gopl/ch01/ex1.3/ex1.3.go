// Exercise 1.3 : Experience to measure the difference in running time between
// potentially inefficient versions and the one that uses strings.Join.

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var args = os.Args[1:]

func main() {
	echo1(args)

	fmt.Println("---")

	echo3(args)
}

func echo1(args []string) {

	start := time.Now()
	fmt.Println("Start: ", start)

	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Elapsed: ", elapsed)
}

func echo3(args []string) {

	start := time.Now()
	fmt.Println("Start: ", start)

	fmt.Println(strings.Join(args, " "))

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Elapsed: ", elapsed)
}
