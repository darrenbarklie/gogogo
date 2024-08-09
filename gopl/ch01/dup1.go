// Dup1 prints text of each line that occurs more than once
// in the standard input, proceeded by its count

// Usage: Run ./dup1, pass in strings on new lines
// Ctrl+D to submit EOF, when first character in new line

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
