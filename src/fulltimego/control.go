package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	names := []string{"John", "Paul", "George", "Ringo"}
	for i := range names {
		fmt.Println(names[i])
	}
}
