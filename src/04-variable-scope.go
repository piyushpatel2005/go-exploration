package main

import (
	"fmt"
)

func main() {
	name := "Louis"
	{ // defines new block
		age := 28
		fmt.Println(age)
		fmt.Println(name)
	}
	fmt.Println(name)
	// fmt.Println(age)
}
