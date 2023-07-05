package main

import "fmt"

func main() {
	i := 10

	fmt.Printf("(i < 100) && (i > 5): %t\n", (i < 100) && (i > 5))     // true
	fmt.Printf("(i < 100) && (i > 500): %t\n", (i < 100) && (i > 500)) // false

	fmt.Printf("(i > 100) && (i > 500): %t\n", (i < 100) && (i > 5))   // false
	fmt.Printf("(i < 100) && (i > 500): %t\n", (i < 100) && (i > 500)) // true

	// unary operator - reverses logical output
	fmt.Printf("!true: %t\n", !true)
	fmt.Printf("!false: %t\n", !false)
	fmt.Printf("!(i > 100): %t\n", !(i > 100))
}
