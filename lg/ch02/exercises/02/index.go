package ch02exer02

import (
	"fmt"
	"reflect"
)

func main() {

	const value = 100

	var i = int(value)
	var f = float64(value)

	fmt.Printf("%v\n", i)
	fmt.Printf("%v\n", reflect.TypeOf(i))
	fmt.Printf("%v\n", f)
	fmt.Printf("%v\n", reflect.TypeOf(f))

	fmt.Println("Done!")
}
