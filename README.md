# Go Programming

- Developed and supported by Google, Version 1 released in 2012
- Go language is very fast. Machine language are CPU instructions and are very simple and straight forward and it runs directly on processors. Assembly languages are similar to machine language but there are English mnemonics. Go is a high level language. In compiled code or interpreted code, there is compile or interpretation step that needs to happen. Go language also has garbage collection for automatic memory management.
- It has simpler objects. Go is weakly object-oriented language. In Go, we don't have class, but we have `struct` with associated methods. In Go, we don't have constructor, inheritance or generics.
- It has efficient concurrency built into language. Concurrency is the management of multiple tasks at the same time. This requires management of task execution, so we need communication between tasks. `Goroutines` represent concurrent tasks. `Channels` are used to communicate between tasks. `Select` enables task synchronization.

Companies like Google, AT&T, Facebook, Netflix and many more use Go. Tools like Docker, Kubernetes, Terraform are written in Go so the future of Go is very bright.

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

There must be one package called `main`. The codes can be grouped into different packages and can be imported in another source code file. All the files in the same folder must have same package name. `package` statement is required at the beginning of any Go code. Unused imports and unused code results in compilation errors. When we build the main package, it creates executable program. Main package needs to have a `main()` function which is entry point for the go code. There is no semi-colon. We can put semicolon but it will be removed by Go compiler automatically. `go run main.go` compiles and runs the code whereas `go build` will run faster but requires build step.

```shell
go mod init some-mod
go build
ls
./main
```

```go
package main
import "fmt" // importing a package
func main() {
  fmt.printf("Hello, world\n")
}
```

Every go project requires a `mod` file which can be created using `go mod init <module_name>`. This creates a module file with module name and minimum version of Go SDK.

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

```go
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

```go
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
var x int // variable name type, int is interesting because it will assign memory 4 byte or 8 byte depending on architecture 32bit or 64 bit. So, it will be int32 on 32-bit architecture and int64 for 64-bit architecture.
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
// When string is defined, it allocates two words in memory. The first word represents a pointer and the second word represents number of bytes of a string with zero value as 0.
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

```go
var grades int = 12
var isCheck bool = true

fmt.Println(" type of grades: %T", grades)
```

There is also a built-in function `TypeOf` which can print the type of a variable. This exists in `reflect` package. Alternatively, we can use `%T` format specifier to get the type of the variable or value. [This code](src/06-type-conversion.go) shows examples of their use.

The process of converting one data type to another type is called type casting. For converting to different data type, we just have to wrap the variable with the type we want to convert to. 
Data types can be converted to different type but it can also result in loss of value or error. If we try to convert incompatible types from one type to another, it will give compilation error. There is [**strconv**](https://pkg.go.dev/strconv@go1.20.5) which can be used to convert integer to string or string to integer. There are also `Parse` functions for type conversion. **Constants** are variables whose value once initialized cannot be altered.

```go
const <name> <datatype> = <value>
const NAME = "hello"
```

Constants are untyped unless they are explicitly given a type at declaration. Like above is untyped. This allows it to be more flexible. They are typed if we specify the data type at declaration. We cannot declare the constant but not initialize the value. The value cannot be assigned later or the zero value is not applied to it. Also the short hand variable assignment does not work for constants. Constants are usually defined as global scope.

## Operators

Operators specify the operation we want to perform on operands. Operands are the elements on which operation is being performed. Operators can be classified as [comparison operators](src/08-comparison-operators.go) (==, !=, >, <, >=, <=), [logical operators](src/10-logical-operators.go) (&&, ||, !), [arithmetic operators](src/09-arithmetic-operators.go) (+, -, *, /, %, ++, --), [assignment](11-assignment-operators.go) (=, :=, +=) and [bitwise operators](12-bitwise-operators.go) (&, |, >>, <<, ^).

## Control Flow

The program execution can be varied by using control flow structures. Usually the execution flow is from start to finish, but with some conditionals, we can alter the execution flow based on the output of the condition check. `if...else` structure comes in play for such conditions. If the condition specified in the `if` expression is true then that block is executed else `else` block gets executed. The parenthesis around if expression are optional. We can also have nested `if...else` blocks. We can also specify multiple conditions using `if...else if...else` control flow. If one of the condition is true, the control will come out of all `if..else` blocks. The [example code](src/13-if-else.go) demonstrates this. The `else` block has to be on the same line where `if` block ends with closing curly brace else it will throw an error.

```go
if ( expression ) {
    ... ...
} else if condition_2 {
    ... ...
} else {
    ... ...
}
```

Similar to `if...else` block, `switch...case` provides multi-way control flow. The syntax looks like below. Unlike some of the other programming language such as C or Java, in Go, the switch statement is not fall through by default. If we need it to behave like Java, we can use `fallthrough` keyword in the case block. We can specify this keyword in some of the case blocks and can skip some of the blocks when not needed. If specified, it will fall through successive cases.

```go
switch expression {
    case value1:
        // execute block
        fallthrough
    case value2:
        // block 2
    default:
        // default block if no match found
}
```

We can also omit expression and use the variables in each of the `case` blocks to evaluate a conditional expression. The actual variable being checked is also accessible for reading in the case block as shown in [example code](src/14-switch-case.go).

```go
switch {
    case condition1:
        // statements
    case condition2:
        // statements 2
    default:
        // default statements
}
```

### For loop

If we want to perform certain action multiple times, we can do that using loops. For example to print Hello, we can type three types or better, we can run a loop. Loops execute until a certain condition is false.

```go
for initialize; check_condition; post_iteration {
    ... ...
}

for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

We can also skip initialization step and post iteration statement. If we skip conditional check, we may get infinite loop which is very bad situation. To break out the loop, we can also use `break` or to continue execution on a certain condition without performing an action, we can use `continue` statement. [Example code](src/15-for-loop-break-continue.go) shows usage of for loop, break statement and continue statement.

## Collection Data types

### Arrays

[Arrays](src/16-arrays.go) are a collection of similar data type elements located at contiguous memory location. Each of these elements are placed next to each other. They can be of integers or strings. Most of the times, we will be using collection of data than individual data. For example, to store information about employee salary, we would need a collection or an array. An array has a pointer at the beginning of the array and also has length property which represents number of elements in an array and a capacity which is the size of the array or the number of elements it can hold. If we know the starting memory address and the data type of an array, we can find memory address of each of the elements in an array. By default, uninitialized array contains zero value of the data type.

```go
var <name> <size> <datatype>
var salaries [5] int
var names [5] string
var names [5]string = [5]string{"John", "Jenny", "Zoe"} // the number of elements on the right should be less than or equal to capacity on the left side. If less, other elements get zero value.
names := [5]string{"John", "Jenny", "Zoe"}  // shorthand declaration
names := [...]string{"John", "Jenny", "Zoe"} // do not need to specify length of the array, compiler calculates it.
```

The length of the array can be found using `len()` function. To get individual elements, we can use array index. Indices start at 0. So, to access element at index 1, we can use `names[1]`. We can also change the value at any index position using assignment operators. Index must be between 0 and `len(arr) - 1`. If we go above or below this index, we can compilation error (`out of bounds`).
For iterating through an array, we can use `for` loop. The iterating variable is usually called `i` for index. There is another method to loop through an array without using index variable in which we don't have to worry about off by 1 error in normal for loop. This is using `range` keyword.

```go
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}

for index, element := range arr {
    fmt.Printf("Index: %d, value: %s\n", index, element)
}
```

If we don't need `index` of elements, we can replace that with `_` and it will be ignored.

```go
for _, elem := range arr {
    fmt.Println(elem)
}
```

Multi-dimensional array is like an array of array with multiple dimensions like in Mathematics. To access individual elements, we have to use multiple indices.

```go
arr := [2][2]int{{2,4}, {1,3}}
arr[1][1] // 3
```

### Slice

Slice is a continuous segment of an array. It provides numbered sequence of elements from original array. Unlike arrays, these can be of variable typed and are more flexible. Elements can be added or removed easily. Slices contain three major componenets. Pointer to the first element of the slice. The length of the slice. Length of the slice are the number of elements the slice contains. The capacity of slice is calculated from the first element of the slice to the last element of the underlying array. So, capacity will always be greater than or equal to length of the slice. Declaration of a slice is similar to array except that it does not need the size of the slice. First an array is created and then the slice reference is returned.

To find, length we use `len` and for capacity, we can use `cap` functions.

```go
<slice_variable_name> := []<data_type>{.... value .....}
numbers := []int {1, 2, 3, 4}
fmt.Println(numbers)

// Another way to create a slice is to use make function
slice := make ([] <data_type>, length, capacity)
slice := make([] int, 3, 5)
fmt.Println(slice) // will have zero value for int
fmt.Printf("Length %d, Capacity: %d\n", len(slice), cap(slice))
```

To create a slice from an array, we can use `array[start_index: end_index]`. If you're coming from Python background, then this may feel very familiar as Python list also has this kind of slicing. In this case, the element at `end_index` is not included in the slice.

```go
numbers := [8]int {0,1,2,3,4,5,6,7}
slice1 := numbers[0:5]
fmt.Printf("Length of array: %d, Capacity of Array: %d\n", len(numbers), cap(numbers)) // len = capacity for array
fmt.Printf("Length of slice1: %d, Capacity of slice1: %d\n", len(slice1), cap(slice1))
fmt.Println(slice1)
slice2 := numbers[:5] // Without start index, it's 0 by default
fmt.Println(slice2)
slice3 := numbers[6:] // end index is length of the array by default
fmt.Println(slice3)

// We can also create slice from another slice
var mini_slice1 = slice1[1:3]
fmt.Println(mini_slice1)
```

Slice is basically using the same data from the original array. It only has different pointer at the beginning of the slice. If we modify the data in slice, the original array data also gets modified. Slice is not a deep copy of the original array.

```go
numbers := [5]int{1, 2, 3, 4, 5}
slice := numbers[1:]
fmt.Println(slice)
slice[2] = 2
fmt.Println(slice)
fmt.Println(numbers)
```

We can append new elements to a slice using [built-in function `append`](https://pkg.go.dev/builtin#append). From the documentation, it looks like first argument is a slice and the next arguments are list of values of the same type. If underlying array has enough capacity, it will be added and returned. If it does not have enough capacity, a new array will be allocated and returned. The new slice will have twice the capacity of the original.

```go
array := [5]int{1,2,3,4,5}
slice := array[:3]
fmt.Printf("%v has length %d and capcity %d\n", slice, len(slice), cap(slice))

slice = append(slice, 40, 50, 60)
fmt.Printf("%v has length %d and capcity %d\n", slice, len(slice), cap(slice))
```

We can append a slice to another slice using `append(slice, slice2...)` using three dots (`...`). This is variadic function which is acceptable for situations where function takes variable number of arguments. We can delete an element from a slice by splitting the slice at that index position and then appending two slices.

```go
numbers := [4]int{1,2,3,4}
numbers2 := [4]int{5,6,7,8}
slice1 := numbers[:]
slice2 := numbers2[:2]
output_slice := append(slice1, slice2...)
fmt.Println(output_slice)

output_slice1 := output_slice[:2]
output_slice2 := output_slice[3:]
deleted_element_slice := append(output_slice1, output_slice2...)
fmt.Println(deleted_element_slice)
```

There is also [copy function](https://pkg.go.dev/builtin#copy) which copies all elements from the source slice into destination slice and it returns the number of elements copied. Here, both slices need to be of the same type.

```go
new_slice := copy(destination_slice, src_slice)

src = []int{1,2,3,4,5}
dest = make([]int, 4)
count := copy(dest, src)
fmt.Printf("Destination slice: %s, number of elements copied: %d\n", dest, count)
```

Looping through slice is same as looping through array. 

### Map

Map is unordered collection of key/value pairs. They are usually associated with key. So, they are called associative arrays or dictionary in other languages. They are implemented using hash tables and due to that they provide constant time lookup. So, add, get and delete operations are very fast with maps.

```go
var <name> map[<key_data_type>]<value_data_type> //declare a map using key type and data type
var employees map[string]int // employee salary map. This creates nil map and we cannot assign key/values to nil map.

var employees map[string]int // creates empty map[]
employees["Akshay"] = 90000 // throws error
fmt.Println(employees)
```

To create a map with key/value pairs we also have to initialize it with key-value pairs.

```go
<variable_name> := map[<key_type>]<value_type>{<kv-pairs>}
numbers := map[string]int{"one": 1, "two": 2}
fmt.Println(numbers)
// To find the lenght of the map, we use len function
fmt.Println(len(numbers)) // 2
```

We can also declare and initialize a map with `make` function.

```go
numbers := make(map[string]int)
fmt.Println(numbers) // []
len(numbers) // 0
```

We can get values for a specific key using its keys. When getting values, we get value and a boolean flag if the key exists or not. If key doesn't exist, the first value will be zero value

```go
numbers := map[string]int{"one": 1, "two": 2, "three": 3}
fmt.Println(numbers["one"])
value, found := numbers["three"]
fmt.Println(found, value)
value, found := numbers["four"]
fmt.Println(found, value)

// To add new key just use assignemnt using new key-value pairs
numbers["four"] = 4
fmt.Println(numbers)
// If the key already exists and we try to assign, it will overwrite or update the value
numbers["four"] = 40
fmt.Println(numbers)
// delete a kay value pair using delete
delete (numbers, "four")
fmt.Println(numbers)
```

Looping through map requires key and value pairs and we can again use the `range` keyword for it.

```go
for key, value := range numbers {
    fmt.Println(key, "has value", value)
}

// delete map using 
for key, value := range numbers {
    delete(numbers, key)
}
fmt.Println(numbers)
// Alternatively, we can re-initialize with make function to clear numbers map
numbers = make(map[string]int)
fmt.Println(numbers)
```

## Functions

Functions are code which carry out a specific job. They are self-contained. They help us divide a program into smaller manageable, organized blocks of code. It takes an input and returns an output. This is reusable and can be used multiple times. This also avoids repeatitive code. It also simplifies the coding because we don't have to know the underlying code or how it works, we just use it. so, it provides higher level of abstraction. It's important to understand what is parameter. Parameters are the names listed in the function definition. Function arguments are the actual values passed to the function when calling a function. If we try to pass more arguments, it will give compilation error. Function name must begin with a letter. It follows same naming convention as a variable.

```go
func <func_name>(<params>) <return_type> { // function signature
    // statements, function body
    return <return_type>
}

func add(a int, b int) int {
    sum := a + b
    return sum
}
sum := add(2, 3)
fmt.Println(sum)
```

[Function](src/19-functions.go) return value must match the function signature. Go allows us to return multiple values like Python. 

```go
func division(a int, b int) (int, int) {
}
```

Go also allows named return values. When we use this feature, we do not need to explicitly specify which local values we want to return. It will by default return the values for the named variables defined in the return type of the function declaration. There is also no need to declare the local variables as they are automatically declared. This improves code readability.

```go
func division(a int, b int) (quotient int, remainder int) {
    quotient = a / b // no need to declare the variable, directly initialize it.
    remainder = a % b
    return // no need to say what we need to return
}
```

Something that's interesting with Go is that it provides blank identifiers which helps us ignore one of the values which we don't need. Let's say we want to get only the remainder, so in such cases, we can ignore the quotient part using blank identifier where we call this function. If we don't do this, then Go gives compilation error.

```go
quotient = division(5,2) // error
_, remainder = division(5,2)
fmt.Println(remainder)
```

**Variadic functions** are functions which accept variable number of arguments of the same type. For declaring variadic function, the type of the final parameter is preceded by elipsis (...). The example of such functions include `fmt.Println` or even `append` function we have seen which takes multiple arguments of the same type. The only limitation of this functions is that this parameter must be at the end of the function. So, we cannot have a function like this. The variadic argument will be stored in a slice.

```go
func something(a int, b ...int, c float32) int {}
```

```go
func sum(numbers ...int) (sum int) {
    sum = 0
    for _, value := range numbers {
        sum += value
    }
    return
}
```

Another example, where it takes multiple arguments with one variadic argument

```go
func greetings(greeting string, guests ...string) {
	for _, guest := range guests {
		fmt.Println(greeting, guest)
	}
}
```

**Anonymous Functions** are functions without any specific name assigned to it. They are usually used once only that's why assigning a name may not be so much useful. They look and behave same as normal functions, in that they can input arguments and return some output.

```go {linenos=true}
output := func (a int, b int) int {
    return a + b
}
fmt.Printf("%T\n", output) // func(int, int) int
fmt.Println("%d\n", output(2,3)) // 5

result := func (a int, b int) int {
    return a + b
}( 10, 20)
fmt.Printf("%T\n", result) // int
fmt.Println("%d\n", result) // 30
```

**Higher order functions** are functions which receive a function as input or return a function as output. This is very useful in functional programming. This makes for readable code.

```go {linenos=true}
func calc(a int, b int, calculation func(int, int) int) int {
	return calculation(a, b)
}

func add(a int, b int) int {
	return a + b
}
calc (5, 2, add)
```

If I want to add another operation that works on two inputs, I just have to define another method like `add` and pass that. This is something I can also pass as an anonymous function and it will still work fine. So, this helps in code reusability.

**Recursive functions** are functions where it calls itself. The function calls itself until it reaches the base case. This is used to recursive problems where the problem can be solved by smaller part of the same problem. For example, factorial, greatest common divisor and even sum we can think as a recursive case. Let's see [couple of examples](src/19-functions.go).

```go
func factorial (n int) int {
    if n == 1 {
        return 1
    }
    return n * factorial (n - 1)
}
```

**Defer** statement is a mechanism to delay the execution of a function until surrounding function returns. Even though this kind of function's arguments are evaluated immediately, the function itself is not called until the surrounding functions are finished executing. If we have multiple `defer` statements, their execution is evaluated in LIFO order. This is used to ensure files are closed once their use is over.

```go
fmt.Println("Hi")
defer fmt.Println("How are you?")
fmt.Println("Hey there")
```

## Pointers

When a variable is declared, a certain amount of memory is allocated based on the variable type. This memory allocation is during the program execution. So, every time we execute the code, we may get different memory location. Pointer is something that holds memory address of a variable. They do not only hold memory addresses but also point to where the memory is allocated and allows us to modify the value stored in that memory location. The same thing happens for functions as well, they are stored in memory.

- address of operator (`&`) gives the address of a variable. This is preceded before the name of the variable. This is also called address of operator.
- dereference (`*`) oeprator is used to get the value at a given address. This is also preceded before the address.

```
var a int = 50
&a = 0x1234
*a = 50
```

The function `new` can be used to declare any variable. By default, this returns the pointer to that variable and it sets that variable to zero value.

```go
ptr := new(int)
fmt.Println(*ptr) // 0
```

```go
func main () {
    var a int = 50
    fmt.Printf("%T %v \n", &a, &a) // *int 0x3241abc
    fmt.Printf("%T %v\n", *(&a), *(&a)) // int 50
}
```

For declaring a pointer variable, we use below syntax.

```go
var <pointer_name> *<data_type>
var ptr_id *int
var ptr_name *string

// initialize pointer using &
ptr_id = &id
ptr_name = &address

var <pointer_name> = &<variable_name> // data type inferred according to variable
<pointer_name> := &<variable_name>
```

```go
var a *int
var str *string
fmt.Println(a) // <nil>
fmt.Println(str) // <nil>

intData := 8
a = &intData
fmt.Println(a)
fmt.Println(*a)
```

Now that we are aware of pointers, we can understand argument passing by value. The parameter is copied into another location of the memory. That copy is accessible only within the scope of the function and when we modify that variable, it will modify only copy and will not change the original variable value.

On the other hand, when a function parameter is passed by reference, the actual memory address is passed into the function call as a pointer. So, if the function modifies the values, the value in the calling function changes. Slices, maps and other complex data types are passed by reference. So, if we modify the slice in the function, it will modify the original slice as well as original array.

## Structure

It's a user-defined data type which groups several data elements for related entity. This allows to store multiple values of different data types under single variable name. For example, a person may have name, address, citizenship, security number, etc. So, we can create a structure for Person and it can include all these information under single variable. To initialize a structure we can use below syntax. Struct assigns contiguous blocks of memory for each of its members. So, if we have a struct with 2 fields of int16, it will assign 4 bytes plus a padding byte.

```go
type <Name> struct {
    field1 <data_type>
    field2 <data_type>
    ... ...
}

type Point struct {
    x int
    y int
}
type Person struct {
    firstName string
    lastName string
    age string
    address string
    phones []string
}
// initialization
var <variable_name> <struct_name>
var pt Point

fmt.Printf("%+v", pt)
```

`new` keyword allocates memory for each of the fields of the struct, assigns them zero values and returns the pointer to the struct. This is more uncommon because we want to provide initial values to fields.

```go
person := new(Person) // new field 
fmt.Printf("%+v", person)
```

```go
<variable_name> := <struct_name> {
    field1: value1,
    field2: value2,
    ... ...
}
```

```go
pt := Point {
    x: 1,
    y: 2,
}
fmt.Printf("%+v", pt)
```

It is also possible to omit the field name and the values will be assigned in the sequence of field names. This is less common due to code readability.

```go
pt := Point(1, 2)
fmt.Printf("%+v", pt)

// Accessing fields using dot operator
fmt.Ptinf("%d, %d\n", pt.x, pt.y)
pt.x = 2
fmt.Printf("%d, %d\n", pt.x, pt.y)
// If you try to access a field that's not part of the struct, that will error out
fmt.Println(pt.z) // error
```

Passing structures as parameters is similar to using any other variable

```go
type Point struct {
    x int
    y int
}

func calculateDistance(pt Point) {
    return math.Sqrt(math.Pow(pt.x, 2) + math.Pow(pt.y, 2))
}
```

We can also pass by reference

```go
type Person struct {
    firstName string
    lastName string
    fullName string
}

func getFullName(p *Person) {
    (*p).fullName = p.firstName + " " + p.lastName
}

func main() {
    p := Person {
        firstName = "First",
        lastName = "Last",
        fullName = ""
    }
    fmt.Printf("%+v\n", p)
    getFullName(&p)
    fmt.Printf("%+v\n", p)
}
```

Comparing structs is possible using equality operator. This checks the fields for comparision

```go
p1 := Point(1,2)
p2 := Point(1,1)
p3 := Point (1,2)
fmt.Println(p1 == p2) // type and fields should be same
fmt.Println(p1 == p3)
```

In Go, we have methods which are defined differently. A method is something that augments a function with an extra parameter that specifies the receiver. This argument is added right after `func` keyword and it accepts only one argument. This is called a receiver. A method is a function that has a defined receiver. This ensures that any instance of `Point` struct will have `calcDistance` available with them.

```go
func (p Point) calcDistance() float64 {

}

func (p *Person) populateFullName() {

}
```

This method cannot be same name as one of the fields of the struct.
**Method sets** are a set of methods available to a data type and are used to encapsulate functionality to specific data type.

```go
type Employee struct {
    name string
    role string
    department string
    salary int
}

func (e *Employee) calculateBonus() float64 {
    return 0.10 * e.salary
}

func (e *Employee) calculateGrossSalary() float64 {
    if e.department == "Sales" || e.department == "Marketing" {
        return e.salary + e.calculateBonus() + 0.10 * e.salary
    } else {
        return e.salary + e.calculateBonus()
    }
}
```

**Interfaces** specifies method sets and is used to create modularity in Golang. It is like a blueprint for a method set and it is defined so that it can be implemented by other types like Structures. They only provide method signatures for each method of a method set. These are then implemented by other types. They cannot even have any variable declarations. So, keep in mind, interfaces do not implement methods.

```go
type <interface_name> interface {
    // method signatures
}

type Account interface {
    withdrawMoney(int amount) bool
    depositMoney(int amount) bool
    getBalance() float64
}
```

Interfaces are defined so that other types can implement them. In order to implement an interface, a type has to implement all of interface's methods. There is no `implements` keyword like Java or C#. These are implemented implicitly and does not use specific keyword to implement an interface. This is how we can implement features like polymorphism from other object oriented programming languages.

```go
type Shape interface {
    area() float64
    perimeter() float64
}

type Square struct {
    side float64
}

func (s Square) area() float64 {
    return s.side * s.side
}

func (s Square) perimeter() float64 {
    return 4 * s.side
}

type Circle struct {
    radius float64
}

const PI := 3.14
func (c Circle) area() float64 {
    return PI * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
    return 2 * PI * c.radius
}
```

Here, we can say `Circle` and `Square` types implicitly implements `Shape` interface. This allows us to call a method on both those types without knowing their types.

```go
func printDetails(s Shape) {
    fmt.Printf("Area = %.2f, Perimeter = %.2f\n", s.area(), s.perimeter)    
}

func main() {
    square := Square{side: 2}
    circle := Circle{radius: 2}
    printDetails(circle)
    printDetails(square)
}
```
