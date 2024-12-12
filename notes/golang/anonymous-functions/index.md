# Anonymous Functions

Anonymous Functions are functions without any specific name assigned to it. They are usually used once only that's why assigning a name may not be so much useful. 

**Anonymous functions** look and behave same as normal functions, in that they can input arguments and return some output. Anonymous functins are also called *function literal*. This is one of the fundamental feature of any functional programming language. Go does support functional programming with built-in features.

## Defining Anonymous Functions

These are defined using below syntax.

```go
func function_name(parameters) (return_type) {
    // statements
}
```

At first time, this syntax may look awkward but you will get used to this as you use them more often.

**Example:**

```go
package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello")
	}()
}
```

Here, we defined a function inside our `main` function on line `6`. However, this function does not have any name. Also interesting thing is that after the definition of the function, we are right away calling this function using pair of parentheses on line `8`. 

**Output:**

```output{ lineNos=false }
Hello
```

## Anonymous Function with Parameters

Anonymous functions can take parameters which can be used inside the function.

```go
package main

import "fmt"

func main() {
	sum := func(a int, b int) int {
		return a + b
	}(10, 20)
	fmt.Printf("%d\n", sum) // 5
}
```

Here, we defined a function which takes two arguments and returns the sum of those two numbers.

**Output:**

```output{ lineNos=false }
30
```

Interestingly, we can also assign anonymous function to a variable. The data type of that variable will be same as the signature of the function.

**Example:**

```go
package main

import "fmt"

func main() {
	output := func(a int, b int) int {
		return a + b
	}
	fmt.Printf("%T\n", output)       // func(int, int) int
	fmt.Printf("%d\n", output(2, 3)) // 5
}
```

In this case, we assign the function to a variable `output`. We can treat this variable as a function.

**Output:**
```output{ lineNos=false }
func(int, int) int
5
```

We can also return an anonymous function as a value from another function.

**Example:**

```go
package main

import "fmt"

func main() {
	sum_function := adder(2, 3)
	fmt.Println(sum_function())
	sum_function = adder(10, 10)
	fmt.Println(sum_function())
}

func adder(a int, b int) func() int {
	sum_fn := func() int {
		return a + b
	}
	return sum_fn
}
```

In this case, from `adder` function definition, we return another function `sum_fn` which does addition of two `int` values.

**Output:**

```output{ lineNos=false }
5
20
```