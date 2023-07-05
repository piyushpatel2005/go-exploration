package main

import "fmt"

func main() {
	// Basic syntax

	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}
	fmt.Println(numbers)

	var numbers2 [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(numbers2)

	// Shorthand notation

	names := [3]string{"John", "Jenny", "Zoe"}
	fmt.Println(names)

	// Length and capacity of Array - both same for array
	fmt.Printf("Length of %v array: %d\n", numbers, len(numbers))
	fmt.Printf("Length of %s array: %d\n", names, len(names))
}
