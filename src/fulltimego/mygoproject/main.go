package main

import (
	"fmt"
	"mygoproject/types"
	"mygoproject/util"
)

func main() {
	number := getNumber()
	fmt.Println("Hello")
	fmt.Println("The number is", number)
	user := types.User{
		Username: "John",
		Age:      util.GetAge(),
	}
	fmt.Println("The user is", user)
}
