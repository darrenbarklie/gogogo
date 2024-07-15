package helloworld

import "fmt"

const (
	english    = "English"
	vietnamese = "Vietnamese"
	french     = "French"

	helloPrefixEnglish    = "Hello, "
	helloPrefixVietnamese = "Chao, "
	helloPrefixFrench     = "Bonjour, "
)

func HelloGreet(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case vietnamese:
		prefix = helloPrefixVietnamese
	case french:
		prefix = helloPrefixFrench
	default:
		prefix = helloPrefixEnglish
	}
	return
}

func main() {
	fmt.Println(HelloGreet("Daz", ""))
}
