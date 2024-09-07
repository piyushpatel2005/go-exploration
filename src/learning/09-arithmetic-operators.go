package main

import "fmt"

func main() {

	a, b := 2, 5
	fmt.Printf("%d + %d: %d\n", a, b, a+b)

	// + operator also works as concatenation operator for strings
	s, t := "Hello ", "there"
	fmt.Printf("%s + %s: %s\n", s, t, s+t)

	fmt.Printf("%d - %d: %d\n", b, a, b-a)
	fmt.Printf("%d * %d: %d\n", a, b, a*b)
	fmt.Printf("%d / %d: %d\n", b, a, b/a) // returns quotient only
	fmt.Printf("%d mod %d: %d\n", b, a, b%a)

	// unary operators
	i := 10
	i++
	fmt.Printf("After i++ : %d\n", i)
	i--
	fmt.Printf("After i-- : %d\n", i)
}
