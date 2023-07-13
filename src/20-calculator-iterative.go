package main

import "fmt"

func main() {
	help()
	fmt.Println("Enter two numbers separated by space to perform operation.")
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Println("Enter operation: ")
	var operation string
	fmt.Scanf("%s", &operation)
	var result int
	if operation == "+" {
		result = add(a, b)
	} else if operation == "-" {
		result = subtract(a, b)
	} else if operation == "*" {
		result = multiply(a, b)
	} else if operation == "/" {
		result = multiply(a, b)
	} else if operation == "%" {
		result = modulo(a, b)
	} else if operation == "?" {
		help()
	} else {
		fmt.Println("Invalid operation. Please try again.")
		return
	}
	fmt.Printf("Results of %d %s %d = %d\n", a, operation, b, result)
}

func help() {
	fmt.Println("Go Calculator App")
	fmt.Println("=================")
	fmt.Println("'+' => addition")
	fmt.Println("'-' => addition")
	fmt.Println("'*' => addition")
	fmt.Println("'/' => addition")
	fmt.Println("'%' => addition")
	fmt.Println("'?' => Help Menu")
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}

func modulo(a int, b int) int {
	return a % b
}
