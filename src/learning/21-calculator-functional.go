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
		result = calc(a, b, add)
	} else if operation == "-" {
		result = calc(a, b, func(a int, b int) int { return a - b }) // anonymous function
	} else if operation == "*" {
		result = calc(a, b, multiply)
	} else if operation == "/" {
		result = calc(a, b, divide)
	} else if operation == "%" {
		result = calc(a, b, modulo)
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
	fmt.Println("'-' => subtraction")
	fmt.Println("'*' => multiplication")
	fmt.Println("'/' => division")
	fmt.Println("'%' => modulo")
	fmt.Println("'?' => Help Menu")
}

func calc(a int, b int, calculation func(int, int) int) int {
	return calculation(a, b)
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
