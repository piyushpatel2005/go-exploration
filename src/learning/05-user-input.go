package main

import "fmt"

func main() {
	// var name string
	// fmt.Print("Enter your name: ")
	// fmt.Scanf("%s", &name)
	// fmt.Println("Hello", name)

	// var user string
	// var age int
	// fmt.Print("Enter your name and age: ")
	// fmt.Scanf("%s %d", &user, &age)
	// fmt.Println("Hey,", user, "you are", age, "years old.")

	var str string
	var num int
	count, err := fmt.Scanf("%s %d", &str, &num)
	fmt.Println("Modified:", count)
	fmt.Println("Error:", err)
	// pass correct values string and number
	// pass both  string and you will see error.
}
