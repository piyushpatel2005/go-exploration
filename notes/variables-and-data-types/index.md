# Variabls and Data Types

Let's take a look at variables and data types available in Go programming language. We will cover primitive data types and complex data types briefly and see different options for variable declaration.

## Variable Declarations

There are several ways to declare and initialize variables in Go. Go has several built-in types. The general syntax for declaring a variable looks like below.

```go
var <variable_name> <data_type>
```

Above syntax only declares the variables and do not assign any specific values. For example, in below code we declare variables of type `int`, `string`, `float32` and `bool`. We will learn about this data types later in this lesson.

```go
var id int
var name string
var price float32
var isActive bool
```

When we declare the variable without initializing, the variables get what is called **zero value** in Go. These are like *default value*s in other programming languages. So, for example, in above code, the value of `id` will be `0` even though we have not explicitly assigned the value `0` to it. The value of variable `name` will be empty string and similarly `price` will have `0.00` and `isActive` will be `false`. 

**Example:**
```go
package main

import "fmt"

func main() {    
    var id int
    var name string
    var price float32
    var isActive bool

    fmt.Println(id)
    fmt.Println(name)
    fmt.Println(price)
    fmt.Println(isActive)
}
```

**Output:**
```shell{ .show-prompt lineNos=false }
go run main.go
```

```output{ lineNos=false }
0

0
false
```

## Variable Initialization

We can also provide specific value to a variable when we declare a variable. This is called initialization. We can do that like below where we use assignment operator (`=`) to assign value on the right side to the variable.

```go
var id int = 5
var name string = "Jack"
```

We can also declare multiple variables in single statement as long as their data types are same. For example, in below code, on line 11, we declare and initialize two variables `a` and `b` with values 20 and 30. We can also print two variables together by separating those variables with comma (`,`) in `fmt.Println` function as shown in line 18.

**Example:**
```go{ filename="main.go" }
package main

import "fmt"

func main() {    
    var id int = 5 // declare and initialize int variable
    var name string = "Jack" // string variable
    var price float32 = 99.99 // float
    var isActive bool = true // boolean
    var largeInt int64 = 23434523 // int64 data type
    var a, b int = 20, 30 // declare and initialize multiple variables

    fmt.Println(id)
    fmt.Println(name)
    fmt.Println(price)
    fmt.Println(isActive)
    fmt.Println(largeInt)
    fmt.Println(a, b) // print multiple variables
}
```
```shell{ .show-prompt lineNos=false }
go run main.go
```

**Output:**
```{ lineNos=false}
5
Jack
99.99
true
23434523
20 30
```

### Shorthand assignment

There is another shorthand notation for declaring and initializing values to the variable using `:=` operator. This is used only for initialization of the variable and once initialized, we have to use assignment operator (`=`) to modify variable's value. This is generally the preferred method for declaring and initializing variables in Go community.

```go
a := 10
b, c := 20, 30
a = 100
```

In this case, we have not specified the data type of the variable `a`, `b` and `c`. However, based on the values being assigned on the right side, the compiler infers the data types of those variables as `int`.

**Example:**
```go{ filename="main.go" }
package main

import "fmt"

func main() {    
    a := 10
    b, c := 20, 30
    fmt.Println(a, b, c)
    a = 100
    fmt.Println(a)
}
```

**Output:**

```bash{ .show-prompt lineNos=false }
go run main.go
```
```{ lineNos=false }
10 20 30
100
```

If we try to re-assign these variables with `string` data type, the compiler throws an error. This is the first encouter with error so try to read the error and understand. Here, it says "Hello" which is string cannot be used as `int` in assignment because of static type checking.

```go{ filename="main.go" }
package main

import "fmt"

func main() {    
    a := 10
    fmt.Println(a)
    a = "Hello" // This line fails
}
```

```bash{ .show-prompt lineNos=false }
go run main.go
```

```{lineNos=false}
# command-line-arguments
./main.go:8:9: cannot use "Hello" (untyped string constant) as int value in assignment
```

## Constants

Constants are like variables however once initialized, the value cannot be modified. These are declared using `const` keyword. We cannot declare a constant without initializing a value, i.e. the value cannot be assigned later or zero value does not apply to constants. Also, shorthand variable assignment does not apply to constants. They will usually have a global scope. The basic syntax for constants look like this.

```go
const <constant_name> <datatype> = <value>
```

For example, 

```go
const PI float32 = 3.14159
```

## Data Types

As we saw in examples above that we have data types for primitive types like integer, boolean, float and we also have string data types. We also have some collection data types which help store group of data. This is how most of the programs use computer memory. These collection data types are created with these primitive types. Below table shows the data type and their range.

| Data Type | Description |
|:---------:|:------------|
| int8 | 8-bit signed integer (-128 to 127) |
| int16 | 16-bit signed integer (-32768 to 32767) |
| int/ int32 | 32-bit signed integer  |
| int64 | 64-bit signed integer |
| uint8 | 8-bit unsigned integer (0 to 255) |
| byte | 8-bit unsigned integer, synonymous to uint8|
| uint16 | 16-bit unsigned integer |
| uint/uint32 | 32-bit unsigned integer|
| uint64 | 64-bit unsigned integer |
| uintptr | pointer of type unsigned integer |
| float32 | 32-bit floating pointer number |
| float64 | 64-bit floating point number|
| complex64 | complex number with 32-bit real and imaginary parts |
| complex128 | complex number with 64-bit real and imaginary parts |
| bool | boolean |
| rune | same as int32 |

There are also other container data types which we will discuss in lot more detail in this tutorial series. These are arrays, slices, maps and structs. Below table gives brief idea about those, but we will see their examples throughout the tutorials.

| Data Type | Description |
|:----------|:------------|
| array | These are fixed length container variable which can hold data of the same type |
| slice | These are similar to Arrays but can grow as we add more data values. |
| map | These are key-value pairs. These are called map, dictionary or associative arrays in other languages |
| struct | These are like structure in C and helps like classes in other object oriented programming languages |


### Define Custom Data types

We can also define our own types using `type` keyword. This is usually used with another data types called `struct`s.

```go
type UserId int
type Username string
```

Here, `UserId` is user-defined type which will have value of type `int`. However, this makes lot more readable code when used for `User` collection to store user specific information.
