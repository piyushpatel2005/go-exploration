---
title: "Functions in Golang"
date: 2023-08-10T06:03:06-04:00
draft: false
---

Functions help make code reusable and modular. They also help code readability and are very useful construct in any programming language. This lesson explains how we can define function and how to use them in our code.
<!--more-->

## Overview 
So far we have seen various data structures in Go. This lesson explains what functions are, their advantages and various examples on functions.

- Functions are code blocks which carry out a specific job. We have been using a built-in function `Println` defined in `fmt` package.
- They are self-contained. 
- They help us divide a program into smaller manageable, organized blocks of code. 
- It takes an input in the form of arguments and can return an output. 
- Functions are reusable and can be used multiple times in the program. This also avoids repeatitive code. 
- It also simplifies the coding because we don't have to know the underlying code of a function or how it works, we just use it. so, it provides higher level of abstraction. Just like we have been using `fmt.Println` function without learning how it is written.

## How to define a function in Go

In Golang, functions are defined using `func` keyword. The function can take parameters and can also optionally return a value from inside the function. General syntax for definition of a function looks like this.

```go
func <func_name>(<params>) <return_type> { // function signature
    // statements, function body
    return <return_type>
}
```

Function names follow similar convention as variable names.

- Function name must begin with a letter, and can have any letters and numbers in the name, but cannot start with a number.
- Function names are case-sensitive. So, name `greet` and `Greet` are not the same.
- Function name cannot contain spaces.
- If function name consists of multiple words, each word after the first should be capitalized like this. For example, fullName or FullName.
- Functions that will be exported outside the package, will have names starting with uppercase. If the function name starts with a lowercase letter, it won't be exported, but you can call this function within the same package. We will learn more about this later.

Function can also be defined without any parameter or return type. Let's define a function without parameters first. 

```go
func greet() {
    fmt.Println("Hello there")
}
```

In this case, we defined a function with name `greet`. Inside the pair of parentheses, we can pass parameters, but in this case this function takes no parameters. To use this function, we can call it using function name and pair of parentheses like `greet()`. Again `main()` function is the entry point so we have to call this function inside the `main` function.

**Example:**

```go
package main

import "fmt"

func main() {
	greet()
}

func greet() {
    fmt.Println("Hello there")
}
```

Also note above, we can define the function even after `main` function and it will work fine. With some other languages, we have to define the function first and then only we can use it.

**Output:**
```output{ lineNos=false }
Hello there
```

## Function Parameters and Return Value

It's important to understand what parameters are. They are sometimes confused with arguments. 
Parameters are the names listed in the function definition whereas Function arguments are the actual values passed to the function when calling a function. 
If we try to pass more arguments than the number of parameters expected in the function definition, it will give compilation error.
Below is an example of a function with parameters and return value. Note that the return type is specified after the pair of parentheses.

**Example:**

```go
package main

import "fmt"

func main() {
	sum := add(2, 3)
	fmt.Println(sum)
}

func add(a int, b int) int {
    sum := a + b
    return sum
}
```

In line `10`, we define a function which takes two `int` parameters `a` and `b`. This function also returns an `int` value as specified after parentheses. To return a value from the function, we use `return` keyword on line `12` which returns the sum of `a` and `b`. It is important that the value returned by the function must match the return type mentioned in the function definition.

**Output:**
```output{ lineNos=false }
5
```

### Multiple Return Values

In Go, functions can return more than one values. For example, `division` function below returns quotient and remainder of division operation

```go
package main

import "fmt"

func main() {
	quotient, remainder := division(8, 3)
	fmt.Println(quotient, remainder)
}

func division(a int, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}
```

**Output:**
```output{ lineNos=false }
2 2
```

### Named Return Values
Go also allows **named return values**. When we use this feature, we do not need to explicitly specify which local values we want to return. It will by default return the values for the named variables defined in the return type of the function declaration. There is also no need to declare the local variables as they are automatically declared. This improves code readability.

**Example:**
```go
package main

import "fmt"

func main() {
	quotient, remainder := division(8, 3)
	fmt.Println(quotient, remainder)
}

func division(a int, b int) (quotient int, remainder int) {
	quotient = a / b // no need to declare the variable, directly initialize it.
	remainder = a % b
	return // no need to say what we need to return
}
```

**Output**
```output{ lineNos=false }
2 2
```

Go also provides *blank identifiers* which helps us ignore one of the return values which we don't need. Let's say we want to get only the `remainder`, so in such cases, we can ignore the `quotient` part using blank identifier when we call this function. If we don't do this, then Go gives compilation error.

```go
_, remainder = division(5, 2)
fmt.Println(remainder)
```