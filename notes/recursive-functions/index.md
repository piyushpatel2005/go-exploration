# Recursive Functions

Recursion is one of the techniques used to solve complex problems. This is useful in solving problems where the solution depends on solutions to smaller instances of the same problem.

## What are Recursive Functions?
Recursive functions are functions where it calls itself. The function calls itself until it reaches the base case. This is used to recursive problems where the problem can be solved by smaller part of the same problem. This technique is part of divide and conquer strategy for solving problems.

Recursive function will have two cases. One of them is called a base case. This is where the recursive function terminates or stops calling itself. If we miss writing bse case, this will result in infinitely recursive calls to the function. This will eventually result in StackOverflow error. On the other hand, there will be recursive case where the function calls itself. The recursive case is used to break down complex inputs into simpler ones.

## Recursive Factorial Function 

Let's write very simple recursive function to find factorial of a number. Factorial of a number is calculated using below equation.

$$ n! = n * (n - 1)! * (n - 2)! * ... * 1 $$

```go
package main

import "fmt"

func main() {
	fmt.Println("Factorial of 4:", factorial(4))
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}
```

In this case, we first define the base case. For this we define base case when number is `1`, we have to return 1 else we can call factorial of one lower number multiplied by the number itself.

**Output:**

```output{ lineNos=false }
Factorial of 4: 24
```

## Recursive Fibonacci Function

Another very common example of using recursion is with Fibonacci numbers. Fibonacci numbers are sequence of numbers which are sum of previous two fibonacci numbers. They start with 0 and 1. Fibonacci sequence looks like `0 1 1 2 3 5 8 13 21 34 55`

**Example:**

```go
package main

import "fmt"

func main() {
	fmt.Println(fibonacci(10))
}

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
```

**Output:**

```output{ lineNos=false }
55
```
