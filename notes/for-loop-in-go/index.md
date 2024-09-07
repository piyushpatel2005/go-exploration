# For Loop

In this lesson, we learn about the `for` loop syntax for Golang and how to write for loop in Go.

Go is very simple and tries to minimize developer errors. For this, Go supports only one kind of loops and that is using `for` keyword. Go does not have `while` or `do...while` loop.

## Syntax of For loop

Go syntax for the `for` loop looks like below.

```go
for initialize; check_condition; post_iteration {
    ... ...
}

for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

Here, we don't need parentheses like we didn't need for `if` condition checks. For loop contains three main items. 
1. Initilization of loop variable: This is where the loop variable gets initialized with start value.
2. Condition: This is the condition which is checked after each iteration. If this condition is `true`, loop gets executed and when this condition becomes `false`, the loop terminates.
3. Post Iteration statements: This is where we modify the loop variable so that this doesn't result in infinite loop.

Let's see an example of for loop.

## For Loop Example

```go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
```

**Output:**
```output{lineNos=false }
0
1
2
3
4
```


We can also skip any of these statements. For example, in below code, this same `for` loop is working like a `while` loop which produces same result as above.

```go
package main

import "fmt"

func main() {
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}
```

If we omit certain statements, that will likely cause **infinite loop**. Below is an example of infinite loop which prints `0` forever.

```go
i := 0
for {
    fmt.Println(i)
}
```

## Continue statement

If we want to skip one of the iteration, we can use the `continue` keyword with a conditional and depending on the result of conditional evaluation, that iteration will be skipped. For example, below code skips iteration when `i` has odd value and hence it prints only even numbers.

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
```

**Output:**
```output{ lineNos=false }
2
4
6
8
10
```

## Break Statement

Similarly, if we want loop to break out pre-maturely when certain condition holds `true`, we can use `break` keyword.
In below code example, the loop `break`s when we encounter first number which is divisible by `5`.

```go
package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%5 == 0 {
			fmt.Println("Breaking out of loop when i is", i)
			break
		}
		fmt.Println(i)
	}
}
```

**Output:**
```output{ lineNos=false }
1
2
3
4
Breaking out of loop when i is 5
```