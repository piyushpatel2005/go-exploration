package main

import "fmt"

func main() {
	// Perform bit by bit operations
	a := 4                                  // 0100
	b := 2                                  // 0010
	c := 5                                  // 0101
	fmt.Printf("%d & %d = %d\n", a, b, a&b) // 0000 = 0

	fmt.Printf("%d | %d = %d\n", a, b, a|b) // 1010 = 6

	fmt.Printf("%d ^ %d = %d\n", a, c, a^c) // 0001 = 1

	fmt.Printf("%d << 1 = %d\n", b, b<<1) // 0100 = 4
	fmt.Printf("%d << 2 = %d\n", b, b<<2) // 1000 = 8
	fmt.Printf("%d << 1 = %d\n", c, c<<1) // 1010 = 10

	fmt.Printf("%d >> 1 = %d\n", c, c>>1) // 0010 = 2
}
