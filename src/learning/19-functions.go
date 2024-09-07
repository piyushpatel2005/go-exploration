package main

import "fmt"

func main() {
	greeting := greet()
	fmt.Println(greeting)

	fmt.Println(hello("Jack"))

	a, b := 2, 3
	fmt.Printf("%d + %d = %d\n", a, b, add(a, b))

	quotient, remainder := division(5, 2)
	fmt.Println(quotient, remainder)
	// Blank identifier - we don't need to declare variable if we are not using them
	// quotient = division(5,2) // gives error
	quotient, _ = division(5, 2)
	fmt.Println(quotient)
	_, remainder = division(5, 2)
	fmt.Println(remainder)

	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3, 4, 5))
	greetings("Good morning", "John", "Jenny", "Jessica")
	greetings("Good bye", "John", "Jenny", "Jessica")

	// Anonymous function or Lambda

	fn := func(a int, b int) int {
		return a + b
	}
	fmt.Printf("%T\n", fn)       // func(int, int) int
	fmt.Printf("%d\n", fn(2, 3)) // 5

	result := func(a int, b int) int {
		return a * b
	}(10, 20)
	fmt.Printf("%T\n", result) // int
	fmt.Printf("%d\n", result) // 200

	// Higher order functions

	// Recursive Function
	fmt.Println(factorial(5))
	fmt.Println(factorial(-5))

	fmt.Println(fibonacci(10))

	// defer statement execution is delayed until surrounding functions return
	// println and print functions are in built-int but they are not guaranteed to stay in the language, so noone uses them.
	fmt.Println("hello")
	defer fmt.Println("how are you?") // defer statement runs after the next function call
	a2 := 1
	fmt.Println(a2)
	fmt.Println("another")
	fmt.Println("Hey there!")
	defer fmt.Println("Another defer")
	fmt.Println("Final")

}

func greet() string {
	return "Hello there!"
}

func hello(name string) string {
	return "Hello " + name + "!"
}

func add(a int, b int) int {
	sum := a + b
	return sum
}

// Function can return multiple values

// func division(a int, b int) (int, int) {
// 	division := a / b
// 	remainder := a % b
// 	return division, remainder
// }

// Functions can also have named return values.
func division(a int, b int) (quotient int, remainder int) {
	quotient = a / b // no need to declare the variable, directly initialize it.
	remainder = a % b
	return // no need to say what we need to return
}

// Variadic functions

func sum(numbers ...int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return
}

func greetings(greeting string, guests ...string) {
	for _, guest := range guests {
		fmt.Println(greeting, guest)
	}
}

func factorial(n int) int {
	if n < 0 {
		fmt.Println("Invalid number")
		return -1
	}
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 0 1 1 2 3 5 8 13 21 34 55
func fibonacci(n int) int {
	if n < 0 {
		fmt.Println("Invalid number")
		return -1
	}
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
