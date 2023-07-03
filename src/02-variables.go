package main

import "fmt"

func main() {
	var name string = "Piyush"
	fmt.Println(name)

	// declare a variable and then initialize it.
	var user string
	user = "Joseph"
	fmt.Println(user)

	// variables of one type cannot be assigned another type of value
	// user = 11

	// initialize multiple variables in single line
	var a, b string = "a", "b" // of same data type
	fmt.Println(a, b)

	// initialize multiple variables of different datatypes
	var (
		str   string = "foo"
		index int    = 12
	)
	fmt.Println(str, index)

	// shorthand version of declaring a variable
	s := "hello" // type is inferred
	fmt.Println(s)
	s = "hi"
	fmt.Println(s)
	// s = 10

	// user := "Another" // once declared, we cannot use this syntax.

	var zeroValue int
	fmt.Println(zeroValue)

	// var unused string // unused variable throws compilation error
	// unused = "something"

}
