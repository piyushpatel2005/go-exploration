---
title: "Printing and Taking User Input"
date: 2023-07-27T22:42:28-04:00
draft: false
---

There are various occassions when we have to print values or variables to show to users. Similarly, there will be situations when we want to respond to user input in different ways. For this, we have to have a ability to take user input and flexibility to provide different output. This lesson teaches **how to take user input** and different **ways to format string output**.

## Taking User Input

### Scanf Function

In Golang, we can take user input from the terminal using `Scanf` function in `fmt` module. [Scanf function] takes format string along with variable number of arguments into which the input values will be stored. Interestingly, in Go, usually functions return multiple values. This is one such function where it returns two values: number of variables assigned and error if any during assignment.

```go{ filename="main.go" }
package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello", name)
}
```

```shell{ .show-prompt lineNos=false }
go run main.go
```

```output{ lineNos=false }
Enter your name: Jay
Hello Jay
```

In this case, we specified that `name` variable should be of type string using format specifier `%s`. Also, when assigning values, we use `&` before the variable name. This will become clearer in the Pointers lesson in this tutorial. Let's look at some of the most common format specifiers.

### Format Specifiers

The format specifier are used to define what kind of data is to be printed or scanned. These are derived from C language.

| Format Specifier | Description |
|:-----------------|:------------|
| `%v` | Value in default format. Works with most types. |
| `%+v` | This is useful for printing `struct` data type. This adds field names in output. |
| `%T` | Type of the value |
| `%t` | Boolean value |
| `%d` | Base 10 integer value |
| `%s` | String |
| `%f` | float |
| `%.2f` | Float with precision upto 2 digits |
| `%e` | Number in scientific notation, e.g. -1.234456e+78 |

For more information on **format specifier** refer to [beginning of the fmt package](https://pkg.go.dev/fmt).

Taking input for integer, boolean and float numbers.

```go{filename="main.go"}
package main

import "fmt"

func main() {
    var quantity int
    var price float32
    var isAvailable bool
    fmt.Scanf("%d %f %t", &quantity, &price, &isAvailable) // collecting three variables in single line
    fmt.Println(quantity)
    fmt.Println(price)
    fmt.Println(isAvailable)
}
```

In this example, we are taking three input variables in single line, each of them separated by space. When we hit new line, the values are assigned to each of those three variables in the format we have specified.

```output{lineNos=false}
20 12.99 false
20
12.99
false
```

Let's also verify the values returned by this `Scanf` function.

```go{filename = "main.go"}
package main

import "fmt"

func main() {
    var str string
	var num int
	count, err := fmt.Scanf("%s %d", &str, &num)
	fmt.Println("Modified:", count)
	fmt.Println("Error:", err)
}
```

If we provide 2 input in the same pattern as the format specifier, that is string followed by space and that followed by integer. It gives this output.

```shell{ .show-prompt lineNos=false }
go run main.go
Hello 2
```

```output{lineNos=false}
Modified: 2
Error: <nil>
```

However, if we do not provide integer value which is specified in the format specifier, it throws an error as shown below.

```shell{ .show-prompt lineNos=false }
go run main.go
Hello
```

In this case, it assigned only one variable `str` which is specified by `Modified` line and there is an error.

```output{lineNos=false}
Modified: 1
Error: newline in input does not match format
```

### Scanln Function

There is also `Scanln` function which stops scanning after new line. This function only takes values as string. So, if we take number as input, we will have to convert them to required data type.

```go{filename="main.go"}
package main

import "fmt"

func main() {
    var str string
	var num string
	count, err := fmt.Scanln(&str)
    count,err = fmt.Scanln(&num)
    fmt.Println(str, num)
}
```

```shell{ .show-prompt lineNos=false }
go run main.go
Hello
There
```

```output{lineNos=false}
Hello There
```

## Printing and Formatting String

Again, `fmt` module contains string formatting functions. Let's see few examples.

### Println Function

`Println` function takes variable number of arguments and prints them separated by space. This also returns number of bytes written and error if any.

```go
name := "Angela"
age := 20
fmt.Println("My name is", name, "and I am", age, "years old.")
```

```output{lineNos=false}
My name is Angela and I am 20 years old.
```

### Printf Function

Similar to `Scanf` function, we also have `Printf` function which has C-style printing functionality using format specifier. The format specifiers mentioned above works for `Printf` function as well. This makes writing long string easier to read as those do not need to be separated by comma.

```go
name := "Angela"
age := 20
fmt.Printf("My name is %s and I am %d years old.\n", name, age)
```

```output{lineNos=false}
My name is Angela and I am 20 years old.
```

### Escape Sequences

Something interesting in this line is `\n` at the end of the line. This is called **Escape Sequence** in programming. Some of the characters starting with `\` have special meaning. The character back slash (`\`) is used to escape next character and provide different meaning to next character as mentioned below.

- `\n` - new line
- `\t` - tab
- `\\` - prints single `\`
- `\"` - escape double quote inside double quoted string

## Conclusion

In this tutorial, we saw how to format string using various format specifiers. I showed two methods for taking user input from the console and two methods for printing string. We also briefly looked at escape sequences towards the end of this lesson.