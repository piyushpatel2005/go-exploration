---
title: "Variadic Functions in Go"
date: 2023-08-11T15:40:28Z
draft: false
---

In this lesson, let's learn about variadic functions. These are functions which take variable number of arguments.

<!--more-->

**Variadic functions** are functions which accept variable number of arguments of the same type. 

- These functions accept variable number of arguments.
- Note that the arguments must be of the same data type. 
- For declaring variadic function, the type parameter is preceded by elipsis (`...`). The variadic argument will be stored in a slice.

We have seen few examples of this already like `fmt.Println` or `append` function we have seen which takes multiple arguments of the same type. One limitation of this type of functions is that this parameter must be at the end of the function. So, we cannot have a function like this because the parameter which can take multiple values is defined in the middle. 

```go
func incorrect(a int, b ...int, c float32) int {}
```
However, below function declaration will be acceptable.

```go
func correct(a int, c float32, b ...int) int {}
```

Let's define a `sum` function which can sum variable number of arguments.

**Example:**

```go
package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3, 4, 5))
}

func sum(numbers ...int) (sum int) { // named return value
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return
}
```

As we can see, we can pass any number of arguments to the `sum` function.
**Output:**

```output{ lineNos=false }
6
0
15
```

Let's define another function which can take variable number of guests and greet them with a message.

```go
package main

import "fmt"

func main() {
	greetings("Good morning", "John", "Jenny", "Jessica")
	greetings("Good bye", "John", "Jenny", "Jessica")
}

func greetings(greeting string, guests ...string) {
	for _, guest := range guests {
		fmt.Println(greeting, guest)
	}
}
```

**Output:**

```output{ lineNos=false }
Good morning John
Good morning Jenny
Good morning Jessica
Good bye John
Good bye Jenny
Good bye Jessica
```



