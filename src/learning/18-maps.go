package main

import "fmt"

func main() {
	// Create empty map

	// var employees map[string]int
	// employees["Akshay"] = 90000 // throws error
	// fmt.Println(employees)

	// To be able to add key-value pairs, we have to initialize it.

	// numbers := map[string]int{"one": 1, "two": 2}
	// fmt.Println(numbers)
	// // To find the lenght of the map, we use len function
	// fmt.Println(len(numbers)) // 2

	// Alternatively, use make function to create map. This allows adding new k-v pairs

	// numbers := make(map[string]int)
	// fmt.Println(numbers)      // []
	// fmt.Println(len(numbers)) // 0
	// numbers["one"] = 1
	// fmt.Println(numbers)

	// Indexing for Map

	// numbers := map[string]int{"one": 1, "two": 2, "three": 3}
	// fmt.Println(numbers["one"])
	// value, found := numbers["three"]
	// fmt.Println(found, value)

	// value, found = numbers["four"]
	// fmt.Println(found, value)

	// // Add k-v pairs
	// numbers["four"] = 4
	// fmt.Println(numbers)
	// // If the key already exists and we try to assign, it will overwrite or update the value
	// numbers["four"] = 40
	// fmt.Println(numbers)
	// // delete a kay value pair using delete
	// delete(numbers, "four")
	// fmt.Println(numbers)

	// Looping through maps

	numbers := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, value := range numbers {
		fmt.Println(key, "has value", value)
	}

	// Clean up maps

	// for key, _ := range numbers {
	// 	delete(numbers, key)
	// }

	fmt.Println(numbers)

	numbers = make(map[string]int)
	fmt.Println(numbers)

	// Create map from array using index as key

	arr := [5]string{"hi", "hello", "there"}
	greetings := make(map[int]string, 5)
	for index, elem := range arr {
		greetings[index+1] = elem
	}
	fmt.Println(greetings)
}
