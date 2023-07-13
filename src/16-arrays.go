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

	colors := [...]string{"orange", "red", "blue", "black"} // compiler picks length automatically
	fmt.Println(colors)

	// Length and capacity of Array - both same for array
	fmt.Printf("Length of %v array: %d\n", numbers, len(numbers))
	fmt.Printf("Length of %s array: %d\n", names, len(names))

	// Accessing elements at an index
	fmt.Println(names[1])
	fmt.Println(numbers[3])

	// Re-assign values
	names[1] = "Jane"
	fmt.Println(names)

	// looping through array
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Another for each style loop using range keyword
	fmt.Println("For each loop")
	for index, elem := range numbers {
		fmt.Printf("Element at %d: %d\n", index, elem)
	}

	fmt.Println("Ignore index using _")
	for _, elem := range numbers {
		fmt.Println(elem)
	}

	// Multi-dimensional arrays
	arr := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(arr[2])

	fmt.Println(arr[2][0])
	fmt.Println(arr[2][1])

	// Looping two-dimensional array
	fmt.Println("Looping multi-dimensional array")
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[i][j])
		}
	}
}
