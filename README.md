# Go Programming

- Go language is very fast. Machine language are CPU instructions and are very simple and straight forward and it runs directly on processors. Assembly languages are similar to machine language but there are English mnemonics. Go is a high level language. In compiled code or interpreted code, there is compile or interpretation step that needs to happen. Go language also has garbage collection for automatic memory management.
- It has simpler objects. Go is weakly object-oriented language. In Go, we don't have class, but we have `struct` with associated methods. In Go, we don't have constructor, inheritance or generics.
- It has efficient concurrency built into language. Concurrency is the management of multiple tasks at the same time. This requires management of task execution, so we need communication between tasks. `Goroutines` represent concurrent tasks. `Channels` are used to communicate between tasks. `Select` enables task synchronization.

## Installing Go

Download go tarball. `go1.12.6.linux-amd64.targ.gz`.

```shell
sudo tar -C /usr/local -xzf go1.12.6.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go version
```

The go code is organized into:

src - for source code files
pkg - containing packages (libraries)
bin - contains executables

Three must be one package called `main`. The codes can be grouped into different packages and can be imported in another source code file. When we build the main package, it creates executable program. Main package needs to have a `main()` function.

```go
package main
import "fmt" // importing a package
func main() {
  fmt.printf("Hello, world\n")
}
```

When we import a package, GoTool has to search directories specified by GOROOT and GOPATH environment variables. If we want to install some package in another directory, we have to change these two environment variables to find those extra packages. When we download go language, we get GoTool with it.

- `go build` - compiles the program. We can also provide bunch of packages as arguments. It creates executable for the main package. It will have `.exe` suffix for executable in Windows.
- `go doc` - prints documentation for a package.
- `go fmt` - format source code files
- `go get` - downloads packages and installs them
- `go list` - lists installed packages
- `go run` - compiles go files and runs the executable
- `go test` - runs tests using files ending in "_test.go"

## Packages and Modules

Packages are used to organize Go code. They can be imported from the Go package registry. Package should perform one single task for example, drawing graphics, string conversion, formatting string, handling HTTP requests. To use a package we have to use `import` keyword followed by the name of  the package. We can import multiple packages like this. In Go we can use namespace to organize packages.

```golang
import (
    "fmt",
    "strconv"
)

import "namespace/package_name"

import (
    . "fmt" // use fmt package functions directly rather than using package name and then function name
    stc "strconv" // use strconv package functions using stc.Atoi
)
```

Modules on the other hand, are collection of packages. They are created by having `go.mod` file in the root directory of project. This file contains information about the project (dependencies, Go version and package info). All Go modules have `go.mod` file. Below is example `go.mod` file.

```golang
module example.com/demo

go 1.17

require(
    github.com/alexflint/go-arg v1.4.2
    github.com/fatih/color v1.13.0
)
```

## Variables and Data types

### DataTypes

Data type is a way that programs can interpret the binary numbers. Numbers, letters, words are different data types, for example, "hello", true, 1
If we specify the type and use wrong data type, go gives compile error. All primitive types are numeric.
They include below.

| Data Type | Minimum Value | Maximum Value | Byte size |
|:---------:|:------------:|:-------------:|:---------:|
| int8 | -128 | 127 | 1 |
| int16 | -32768 | 32767 | 2 |
| int/ int32 | -2^16 | 2^16 - 1| 4 |
| int64 | 2^32 | 2^32 - 1 | 8 |
| uint8 | 0 | 255 | 1 |
| byte | 0 | 255 | 1 |
| uint16 | 0 | 2^16 - 1 | 2 |
| uint/uint32 | 0 | 2^32 - 1 | 4 |
| uint64 | 0 | 2^64 - 1 | 8 |
| uintptr | 0 | <pointer_size> | 32/64 bits |
| float32 | | | 4 |
|  float64 | | | 8 |
| complex64 | | | 4 + 4 |
| complex128 | | | 8 + 8 |
| bool | | | |

We can also make type aliases. For example, we can rename `int` to `Id` type.

```golang
type Id int
type Acceleration float32
type Gravitation Acceleration
Id(10)
Gravitation(9.81)
```

### Variable Declarations

Names must start with letter. They can have number, underscore. Every variables need name and type. So, they need declaration. [This code](src/02-variables.go) shows some of the methods for declaring and initializing variables.

```go
var x int # variable name type
var x,y int # declare two variables on same line
// Defining alias (alternate name for a type)
type Celsius float64
type IDnum int
var temp Celsius
var pid IDnum
var x int = 100 // Initialization
var x = 100 // Don't need to specify type of a variable
// Initialize after declaration
var x int // declare
x = 100 // initialize
var x int // default value is x =0
var x string // default value of string is x = ""
// Declaration and initilization in the same statement
x := 100 // type is inferred. This works only inside a function
var (
		str   string = "foo"
		index int    = 12
	)
```

If we are declaring a variable and not initializing with a value, the variable gets a default value, known as zero value. This depends on the data type of the variable. For bool, 'false', for int, `0`, for float `0.0` and for string it is empty string as shown in [this code](src/02-variables.go).

### Printing Variables and values

There are multiple [ways to print variables](src/printing-variables.go). By default, we can add comma to print multiple items.
- `Println` prints with new line
- `Print` prints without new line.
- `Printf` uses format specifier to print template strings. The documentation on format specifiers is [here](https://pkg.go.dev/fmt). Some of the common format specifiers are as below.

**Variable Scope** is part of the program where variables is accessible. Block defines the scope of variables. Inner blocks can access variables defined in the outer block but outer block cannot access variables inside the inner block.
Local variables are variables declared inside a function or block. Global variables are variables declared outside the functional block, declared at the top of the program and can be accessed from any part of the program. Take a look at [examples](src/04-variable-scope.go).

### Taking User Input

One of the ways is using `scanf` function from `fmt` package. This takes a format string and number of arguments. This list are the variables where we want to store the data. The `fmt.Scanf` [function return two values](https://pkg.go.dev/fmt#Scanf), count - number of arguments the function writes to successfully and the error - the error thrown during the execution of the function. [Example code](src/05-user-input.go)

### Converting Data types

Before converting to different data types, we need to know the current data type of a variable. For this, `%T` format specifier can be used.

```golang
var grades int = 12
var isCheck bool = true

fmt.Println(" type of grades: %T", grades)
```

There is also a built-in function `TypeOf` which can print the type of a variable. This exists in `reflect` package. Alternatively, we can use `%T` format specifier to get the type of the variable or value. [This code](src/06-type-conversion.go) shows examples of their use.

The process of converting one data type to another type is called type casting. For converting to different data type, we just have to wrap the variable with the type we want to convert to. 
Data types can be converted to different type but it can also result in loss of value or error. If we try to convert incompatible types from one type to another, it will give compilation error. There is [**strconv**](https://pkg.go.dev/strconv@go1.20.5) which can be used to convert integer to string or string to integer. There are also `Parse` functions for type conversion. **Constants** are variables whose value once initialized cannot be altered.

```golang
const <name> <datatype> = <value>
const NAME = "hello"
```

Constants are untyped unless they are explicitly given a type at declaration. Like above is untyped. This allows it to be more flexible. They are typed if we specify the data type at declaration. We cannot declare the constant but not initialize the value. The value cannot be assigned later or the zero value is not applied to it. Also the short hand variable assignment does not work for constants. Constants are usually defined as global scope.