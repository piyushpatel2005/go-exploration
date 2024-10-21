# Data Type Conversion

Go is strongly typed language. That means it will not accept other types when variable is declared with different type. Data type conversion comes very useful in order to perform certain operation on the data.

## How to check data type in Go?

Before we get into modifying the data type, let's see how we can verify data types in Go. 

### Using Format Specifier

Go provides format specifier `%T` to get the data type of the variable. This can be useful for debugging certain errors.

```go{filename = "main.go"}
package main

import "fmt"

func main() {
	var grades int = 12
    var isHighest bool = true
    fmt.Printf("Data type of grades: %T\n", grades)
    fmt.Printf("Data type of isHighest: %T\n", isHighest)
}
```

```output{ lineNos=false }
Data type of grades: int
Data type of isHighest: bool
```

### Using `reflect.TypeOf` function

There is a built-in function for checking the type of the variable. This function **[TypeOf](https://pkg.go.dev/reflect#TypeOf)** is located in `reflect` module. So, we will have to import this module before using this function.

```go{ filename="main.go" }
package main

import (
    "fmt"
    "reflect"
)

func main() {
	var str string = "hello there"
	var amount float32 = 23.23

	fmt.Printf("Type: %v \n", reflect.TypeOf(str))
	fmt.Printf("Type: %v\n", reflect.TypeOf(amount))
}
```

```output{ lineNos=false }
Type: string 
Type: float32
```

Another important point to note with above code is that when we want to import multiple modules, we can put them in parentheses.

```go
import (
    "fmt"
    "reflect"
)
```

## Type conversion (Type casting)

The processing of converting a value of one type to another is called *type casting*. The easiest way to convert to different type is to wrap the variable with the type we want to convert to.

> Data types can be converted to different type, but it can also result in loss of some precision. When we convert wider type into narrow type, that will result in precision loss. For example, converting `float` to `int` would result in some precision loss.

```go{ filename="main.go" }
package main

import (
    "fmt"
)

func main() {
    var i int = 25
	var convertedInt float64 = float64(i)
	fmt.Printf("%.2f\n", convertedInt)

	var f float64 = 39.99
	var convertedFloat int = int(f) // truncates the number, removes the decimal point values
	fmt.Printf("%d\n", convertedFloat)
}
```

```output{ lineNos=false }
25.00
39
```

> **Note:** If we try to convert incompatible types from one type to another, it will result in compilation error.

Below code throws an error as Go compiler doesn't know how to convert word `hello` to integer.

```go
package main

import "fmt"

func main() {
	var s string = "hello"
	var strToInt int = int(s)
	fmt.Printf("%d\n", strToInt)
```

```output{ lineNos=false }
cannot convert s (variable of type string) to type int
```

For conversion between string and numeric data types, we have `strconv` module. This is used to convert `string` type into `integer` and vice a versa. This module provide a function `strconv.Itoa` to convert integer into string value and `strconv.Atoi` to convert string to integer values. This module also provides various functions to parse different types of values from string like `ParseInt`, `ParseBool`, `ParseFloat` etc. 

The function `Itoa` takes a number and returns string representation of the number. Here, we've used `%q` format specifier so that it prints string in double quotes to make it clear that this is string.

The function `Atoi` takes a string and return two values. First is converted number if successful, but if unsuccessful, second return value will have the error message for why it failed.

```go{ filename="main.go" }
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 10
	str := strconv.Itoa(num)
	fmt.Printf("%q\n", str)

	str = "200"
	value, err := strconv.Atoi(str)
	fmt.Printf("Value: %d, err: %v\n", value, err)
}
```


```output{ lineNos=false }
"10"
Value: 200, err: <nil>
```

## Conclusion

This lesson started with how we can check the data type of a variable in Golang. Go provides `%T` format specifier as well as `reflect.TypeOf` method from `reflect` module to do this. Then, we also looked at how we can convert data types between numeric values using simply the type we want to convert to as a function. For string, we had to use special functions from `strconv` module. These were `Atoi` for String to Integer and `Itoa` for Integer to String conversion. 