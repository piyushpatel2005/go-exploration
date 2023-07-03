package main

import "fmt"

const PI float64 = 3.14

func main() {
	const name = "Fixed"
	// name = "Another name" // cannot assign to name (untyped string constant "Fixed")
	fmt.Println(name)

	// const pi
	// pi = 3.14
	// fmt.Println(pi)

	// const pi := 3.14
	// fmt.Println(pi)

	// use case for constants
	var radius = 2.0
	area := PI * radius * radius
	perimeter := 2 * PI * radius
	fmt.Printf("Area of circle: %.2f, Perimeter of circle: %.2f\n", area, perimeter)

}
