# Higher Order Functions

Higher order functions are functions which receive a function as input or return a function as output. This is very useful in functional programming. This helps with code readability.

## Function as Parameter

Below is an example in which `calc` function takes another function as an argument. This creates a very interesting reading code. In this case, `calc` is called higher order function. I can also pass an anonymous function into this function as we are doing for the `subtraction` variable.

**Example:**

```go
package main

import "fmt"

func main() {
	addition := calc(5, 2, add)
	fmt.Println("Addition:", addition)

	subtraction := calc(5, 1, func(a int, b int) int {
		return a - b
	})
	println("Subtraction:", subtraction)
}

func calc(a int, b int, calculation func(int, int) int) int {
	return calculation(a, b)
}

func add(a int, b int) int {
	return a + b
}
```

In this case, `calc` function can take any function that takes two `int` arguments and returns an `int` value. This way we can perform various mathematical operation by passing different function into it.

**Output:**

```output{ lineNos=false }
Addition: 7
Subtraction: 4
```

## Function as a Return Value

We have already seen examples of functions which return function as output. 

```go
package main

import "fmt"

func main() {
	fmt.Println("Sum of 1 and 2 is", sum(1, 2)())
}

func sum(a int, b int) func() int {
	return func() int {
		return a + b
	}
}
```

In this case, the function call on line `6` looks slightly strange: `sum(1,2)()`. We are using parentheses twice here.

**Output:**

```output{ lineNos=false }
Sum of 1 and 2 is 3
```
