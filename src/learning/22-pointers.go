package main

import "fmt"

func main() {

	// first look

	// var a int  50
	// fmt.Printf("%T %v \n", &a, &a) // *int 0x3241abc
	// fmt.Printf("%T %v\n", *(&a), *(&a)) // int 50=

	// initializing pointers

	var a *int
	var ptr *string
	fmt.Println(a)   // <nil>
	fmt.Println(ptr) // <nil>

	intData := 8
	a = &intData
	fmt.Println(a)
	fmt.Println(*a)

	new_ptr := new(int) // sets to zero value
	fmt.Println("new_ptr", *new_ptr)

	// verify the pointer values, they point to same memory address

	greeting := "Hello"
	ptr = &greeting
	fmt.Println(ptr, *ptr)

	var ptr2 *string = &greeting
	fmt.Println(ptr2, *ptr2)

	ptr3 := &greeting
	fmt.Println(ptr3, *ptr3)
	// Now if I modify this value, then it will change for original variable
	*ptr3 = "Hi"
	fmt.Println(greeting, *ptr, *ptr2, *ptr3)

	// Passing by value. It does not alter the original variable, int, float, string etc.
	number := 10
	fmt.Println("number =", number)
	increment(number)
	fmt.Println("number =", number) // still remains the same
	input := "hello"
	fmt.Println("input =", input)
	concatenate(input)
	fmt.Println("input =", input)

	// Passing by Reference

	concatenate_ptr(&input)
	fmt.Println("input =", input)

	arr := [...]int{1, 2, 3, 4}
	fmt.Println("arr =", arr)
	slice := arr[:3]
	fmt.Println("slice =", slice)
	update_slice(slice)
	fmt.Println("slice =", slice)
	fmt.Println("arr =", arr)

	// The same thing happens for maps
	language := make(map[string]string)
	language["USA"] = "English"
	language["Spain"] = "Spanish"
	language["France"] = "French"
	language["Canada"] = "English"
	fmt.Println(language)
	update_map(language)
	fmt.Println(language)
}

func increment(number int) {
	number += 1 // this modifies the copy of parameter
}

func concatenate(input string) {
	input += " there"
}

func concatenate_ptr(input *string) {
	*input += " there"
}

func update_slice(slice []int) {
	slice[1] = 200
}

func update_map(m map[string]string) {
	m["Canada"] = "French"
}
