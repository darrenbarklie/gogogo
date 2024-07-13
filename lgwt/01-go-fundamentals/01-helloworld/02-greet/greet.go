package main

import "fmt"

func Greet(name string) string {
	return "Greetings, " + name
}

func main() {
	fmt.Println(Greet("Daz"))
}
