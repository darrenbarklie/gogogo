package main

import "fmt"

func main() {
	fmt.Println("Start")

	a1 := addTwoNumbers(2, 3)

	a2i := []int{2, 3, 4, 5, 6}
	a2 := addManyNumbers(a2i)

	fmt.Printf("A1: %v\n", a1)
	fmt.Printf("A2: %v\n", a2)
}

func addTwoNumbers(num1, num2 int) int {
	return num1 + num2
}

func addManyNumbers(nx []int) int {
	var r int
	for i := 0; i < len(nx); i++ {
		r += nx[i]
	}
	return r
}
