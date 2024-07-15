package helloworld

import "fmt"

const helloPrefixEnglish = "Hello, "

func Greet(name string) string {
	return helloPrefixEnglish + name
}

func main() {
	fmt.Println(Greet("Daz"))
}
