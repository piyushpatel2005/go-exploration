package main

import "fmt"

func main() {
	// We have seen assign operator which assigns left operand with value of right side.
	var a int
	a = 10

	// short hand operator
	b := 10

	a += 5
	fmt.Println("a += 5:", a)

	b -= 2
	fmt.Println("b -= 2:", b)

	b *= 2
	fmt.Println("b *= 2:", b)

	b /= 2
	fmt.Println("b /= 2:", b)

	a %= 2
	fmt.Println("a %= 2:", a)
}
