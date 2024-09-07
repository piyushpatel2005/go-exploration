# Error Handling

In Go, there are several kinds of errors such as syntax error, compilation errors, etc. In order to make your code reliable and robust, you need to know how to handle different types of errors and how to show more explanatory message to the end users.

## Types of Errors

Syntax errors are the mistakes a developer makes in writing the code. It is basically not following the grammar of a programming language. These kinds of errors stop the code from compiling and will need to be fixed before running the program. There are logical errors in which case the logic of the program is not right. These are situations like assigning the wrong value to a variable or checking for wrong state or missing a check of `nil`, etc. The third set of errors which is something we know that may occur in certain situations will need particular handling. These kinds of errors include dividing a number by zero, reading a file that doesn't exist.

Most of the times in Go if a function throws an error, it will be part of second return type. For example in `strconv` package, the functions which can raise an error have return type with second value as `error`.

```go
func ParseFloat(s string, bitSize int) (float64, error)
func ParseInt(s string, base int, bitSize int) (i int64, err error)
func QuotedPrefix(s string) (string, error)
```

**Example:**

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "10"
	value, error := strconv.Atoi(str)
	fmt.Println(value, error)
}
```

**Output:**

```output{ .show-prompt lineNos=false }
go run main.go
10 <nil>
```

## Providing Custom Error Messages

We can check if a function has thrown an error by checking second returned value. If the second returned value is `nil`, that means there was no error else we can provide custom error message using `Println` statement. If we want to check the exact error type, we can use `reflect` package as we have seen before.

**Example:**

```go
package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func main() {
	var str string = "10"
	value, error := strconv.Atoi(str)
	fmt.Println(value, error)

	str = "10.5"
	value, error = strconv.Atoi(str)
	if error != nil {
        fmt.Println("Cannot convert", str, "to a number.")
		fmt.Println("Error:", error)
		fmt.Println("Error type:", reflect.TypeOf(error))
	} else {
		fmt.Println(value)
	}
}
```

This code builds on previous code and checks if there is an error. If an error occurs, we print the error and to check the error type we can use `reflect.TypeOf` function.

**Output:**

```output{ lineNos=false }
10 <nil>
Cannot convert 10.5 to a number.
Error: strconv.Atoi: parsing "10.5": invalid syntax
Error type: *strconv.NumError
```

## Error Interface

In Go, the error type is an interface with only one method.

```go
type error interface {
    Error() string
}
```

Any type that implements this interface is called an error. In your application, you can define custom error types by simply implementing the `error` interface.

**Example:**

```go
package main

import (
	"fmt"
	// "errors"
)

type Voter struct {
	name string
	age int
}

type InvalidAgeError struct {
	age int
}

func (e *InvalidAgeError) Error() string {
	return fmt.Sprintf("Invalid age: %d", e.age)
}

func(v *Voter) CanVote() (bool, error) {
	if (v.age < 0 || v.age > 120) {
		return false, &InvalidAgeError{age: v.age}
	}
	return v.age >= 18, nil
}

func main() {
	voter := Voter{name: "John", age: 20}
	canVote, err := voter.CanVote()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("can vote:", canVote)

	voter.age = 121
	canVote, err = voter.CanVote()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("can vote:",canVote)
}
```

**Output:**

```output{ lineNos=false }
can vote: true
Invalid age: 121
can vote: false
```

If we want to create custom error messages, we can easily create one using `errors.New` method. This method takes a `string` which is the error message. Below example shows its usage.

**Example:**

```go
package main

import (
	"fmt"
	"errors"
	"strconv"
)

type Voter struct {
	name string
	age int
}

func(v *Voter) CanVote() (bool, error) {
	if (v.age < 0 || v.age > 120) {
		return false, errors.New("Invalid age: " + strconv.Itoa(v.age))
	}
	return v.age >= 18, nil
}

func main() {
	voter := Voter{name: "John", age: 20}
	canVote, err := voter.CanVote()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("can vote:", canVote)

	voter.age = 121
	canVote, err = voter.CanVote()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("can vote:",canVote)
}
```

**Output:**

```output{ lineNos=false }
an vote: true
Invalid age: 121
can vote: false
```

## Errorf function

In Go, `fmt` package has a function `Errorf` which can be helpful when using data from the method to create custom errors. Although we could use `errors.New`, but it's somewhat limiting in using data from the method. It looks unclear and we have to import additional packages to convert other number types into string data type. With `Errorf` we can use `Printf` style format specifiers. The signature for this function looks like this.

```go
func Errorf(format string, a ...any) error
```

This function simply formats using the format specifier and returns a string that satisfies `error` interface.

```go
package main

import (
	"fmt"
)

type Voter struct {
	name string
	age int
}

func(v *Voter) CanVote() (bool, error) {
	if (v.age < 0 || v.age > 120) {
		return false, fmt.Errorf("Invalid age: %d", v.age)
	}
	return v.age >= 18, nil
}

func main() {
	voter := Voter{name: "John", age: 20}
	canVote, err := voter.CanVote()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("can vote:", canVote)

	voter.age = 121
	canVote, err = voter.CanVote()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("can vote:",canVote)
}
```

Above code is exactly the same except that it uses `fmt.Errorf` to return an `error`. The output is also same as previous output.