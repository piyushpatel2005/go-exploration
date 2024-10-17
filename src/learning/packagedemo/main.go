package main

import (
	"fmt"
	"os/user"
)

func main() {
	fmt.Println("Addition:", Add(10, 20))
	fmt.Println("Subtraction:", Subtract(20, 10))
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Username)
}
