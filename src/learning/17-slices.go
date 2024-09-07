package main

import "fmt"

func main() {
	// Create slice
	// numbers := []int{1, 2, 3, 4} // no need to specify size
	// fmt.Println(numbers)

	// var empty_slice []int = make([]int, 3, 4)
	// fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", empty_slice, len(empty_slice), cap(empty_slice)) // length and capacity different

	// Create slice from array

	// numbers := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	// slice1 := numbers[0:5]
	// fmt.Printf("Length of array: %d, Capacity of Array: %d\n", len(numbers), cap(numbers)) // len = capacity for array
	// fmt.Printf("Length of slice1: %d, Capacity of slice1: %d\n", len(slice1), cap(slice1))
	// fmt.Println(slice1)
	// slice2 := numbers[:4] // Without start index, it's 0 by default
	// fmt.Println(slice2)
	// slice3 := numbers[6:] // end index is length of the array by default, even though we slice from index 6, slice index starts from 0
	// fmt.Println(slice3)
	// fmt.Println(slice3[0])
	// var mini_slice1 = slice1[1:3] // We can also create slice from another slice
	// fmt.Println(mini_slice1)

	// Slices are basically referring to same data as original source array

	// numbers := [5]int{1, 2, 3, 4, 5}
	// slice := numbers[1:]
	// fmt.Println(slice)
	// slice[2] = 2
	// fmt.Println(slice)
	// fmt.Println(numbers)

	// append new elements

	// array := [5]int{1, 2, 3, 4, 5}
	// slice := array[:3]
	// fmt.Printf("%v has length %d and capcity %d\n", slice, len(slice), cap(slice))

	// slice = append(slice, 40, 50, 60)
	// fmt.Printf("%v has length %d and capcity %d\n", slice, len(slice), cap(slice))

	// append slice to another slice and Delete elements of a slice

	// numbers := [4]int{1, 2, 3, 4}
	// numbers2 := [4]int{5, 6, 7, 8}
	// slice1 := numbers[:]
	// slice2 := numbers2[:2]
	// output_slice := append(slice1, slice2...)
	// fmt.Println(output_slice)
	// // Delete element from a slice
	// output_slice1 := output_slice[:2]
	// output_slice2 := output_slice[3:]
	// deleted_element_slice := append(output_slice1, output_slice2...)
	// fmt.Println(deleted_element_slice)

	// Copy slice to another slice

	var src = []int{1, 2, 3, 4, 5}
	var dest = make([]int, 4)
	count := copy(dest, src)
	fmt.Printf("Destination slice: %v, number of elements copied: %d\n", dest, count)

	// Looping over slice, same as arrays
	for i := 0; i < len(dest); i++ {
		fmt.Println(dest[i])
	}
	fmt.Println("========================")
	for _, elem := range src {
		fmt.Println(elem)
	}
}
