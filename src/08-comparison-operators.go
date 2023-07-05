package main

import "fmt"

func main() {
	a, b := 10, 10
	fmt.Printf("%d == %d: %t\n", a, b, a == b)
	fmt.Printf("%d != %d: %t\n", a, b, a != b)

	a = 20
	fmt.Printf("%d > %d: %t\n", a, b, a > b)
	fmt.Printf("%d < %d: %t\n", a, b, a < b)

	fmt.Printf("%d >= %d: %t\n", a, b, a >= b)
	fmt.Printf("%d <= %d: %t\n", a, b, a <= b)

	a = 10
	fmt.Printf("%d > %d: %t\n", a, b, a > b)
	fmt.Printf("%d < %d: %t\n", a, b, a < b)

	fmt.Printf("%d >= %d: %t\n", a, b, a >= b)
	fmt.Printf("%d <= %d: %t\n", a, b, a <= b)

}
