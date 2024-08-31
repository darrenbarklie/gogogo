package ch02exer03

import (
	"fmt"
)

func main() {

	var (
		b      byte
		smallI int32
		bigI   uint64
	)

	b = 255
	smallI = 2_147_483_647
	bigI = 18_446_744_073_709_551_615

	fmt.Printf("%v\n%v\n%v\n", b, smallI, bigI)

	b += 1
	smallI += 1
	bigI += 1

	fmt.Printf("%v\n%v\n%v\n", b, smallI, bigI)

	fmt.Println("Done!")

	// 255
	// 2147483647
	// 18446744073709551615
	// 0
	// -2147483648
	// 0
	// Done!
}
