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
tar -xzf go*.tar.gz
sudo cp -r go /usr/local/
# sudo mv go /usr/local/
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

Now that we are aware of pointers, we can understand argument [passing by value](./src/22-pointers.go). The parameter is copied into another location of the memory. That copy is accessible only within the scope of the function and when we modify that variable, it will modify only copy and will not change the original variable value.

On the other hand, when a function parameter is passed by reference, the actual memory address is passed into the function call as a pointer. So, if the function modifies the values, the value in the calling function changes. Slices, maps and other complex data types are passed by reference. So, if we modify the slice in the function, it will modify the original slice as well as original array.

## Structure

It's a user-defined data type which groups several data elements for related entity. This allows to store multiple values of different data types under single variable name. For example, a person may have name, address, citizenship, security number, etc. So, we can create a structure for Person and it can include all these information under single variable. To initialize a structure we can use below syntax. Struct assigns contiguous blocks of memory for each of its members. So, if we have a struct with 2 fields of int16, it will assign 4 bytes plus a padding byte.

[Examples](./src/23-structures.go)

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

## Concurrency

Let's first understand between sequential and concurrency. Sequential programming is simple. Commands are executed one after another in linear fashion. Then comes multi-tasking where multiple tasks are executed. This is still single-core processor with context switching. Concurrency is like running multiple tasks at the same time. Concurrent programming is not same as parallel programming.

```go
// sequential program
func main() {
    start := time.Now()
    for i:=1; i <= 100; i++ {
        calculateSquare(i)
    }
    elapsed := time.Since(start)
    fmt.Println("Elapsed time", elapsed)
}

func calculateSquare(i int) {
    time.Sleep(1 * time.Second)
    fmt.Println(i * i)
}
```

Above code is [sequential](./src/26-sequential-execution.go), so it will take around 100 seconds.

Go routines is the mechanism used in Go to run concurrent programs. This is considered as a lightweight thread that has its own independent execution. Go routines can be executed concurrently with other go-routines and it is the fundamental way to execute concurrent programs in Go. Go routines are managed by Go runtime. The syntax looks like below which applies to functions or methods. `go` keyword makes it run in a separate go routine.

```go
go func_name()
```

You can find an [example of go-routines](./src/27-go-routine.go) in code directory of this repo.

Main function in the main package is the **main go-routine**. All go-routines are started from [main go-routine](src/28-main-goroutine.go). it represents main program. When main go-routine exits, it assumes all go-routines have been exited. Go routines do not have parents or children. They all execute in parallel. All go routines exit when main go routine exits. To prove that there is no parent-child relationship, we can use below code. The output of this code will not be deterministic because both methods `process1` and `process2` execute in separate go routines.

```go
package main

import "fmt"
import "time"

func process1() {
	time.Sleep(1 * time.Second)
	fmt.Println("Running Process 1")
}

func process2() {
	go process1()
	time.Sleep(1 * time.Second)
	fmt.Println("Running Process 2")
}

func main() {
	go process2()
	time.Sleep(2 * time.Second)
}
```

**[Anonymous go-routine](src/29-anonymous-goroutine.go)** are anonymous functions. They don't have words. They are used only once. They can be called using go-routines. This routine behaves just like any other go-routine.

```go
go func() { }(args...)
```

```go
package main

func main() {
    go func() {
        fmt.Println("An anonymous function")
    }()
    time.Sleep(1 * time.Second)
}
```

### Go routine scheduler

When we launch go program, it will launch OS threads that is equal to logical CPUs available. They are completely managed by kernel or OS level. From creating, managing and terminating threads is managed by OS. We can find available number of logical processors using `runtime` package with a property `Numcpus()` methods. The logical threads are the number of physical cores available in your system multiplied by the number of threads that can run on each core (hardware threads). Go routines are lightweight application threads that run independently. The go routine has scheduler that will multiplex the go-routines on OS level threads in the go runtime. It schedules arbitrary number of go-routines onto an arbitrary number of OS threads(m x n multiplexing).

OS Scheduler manages OS threads for each logical core. Within Go runtime, each of these threads will have one queue associated with it called LRQ (Local Run quque). It consists of all go-routines to be executed in that thread. Go runtime scheduler will be doing scheduling and context switching belonging to LRQ. We also have GRQ (Global run queue). It consists of all go-routines which are not assigned any LRQ. Go scheduler will assign these go routines to any of the Local run queue of OS thread. This is how Go scheduler works and multiplexes go routines on OS threads. Properties of Go scheduler include:

- Cooperative scheduler:It means it's not pre-emptive. There is no time based pre-emption from the OS. In this OS, never interrupts a running process to initiate a context switch from one process to another. Processes must voluntarily yield control periodically or when logically blocked on a resource. Of course, there are some specific checkpoints which will ensure go-routine can yield execution to other go-routines. These are called context switches. For example, runtime calls the scheduler on function calls to check if new go routine needs to be scheduled. In this case, a context switch might happen. It is also possible that current go routine may continue. Examples when these kinds of context switching can happen include: function calls, network calls, garbage collection, channel operations, `go` keyword usage. The scheduler gets an option to context switch, but it does not mean it will always do so.

#### Comparing Go-routines and Threads

- go routines are cheaper. Just few KB in stack and stack can grow and shrink in size. With threads, stack size has to be specified and is fixed in size. OS threads generally start with 1MB. Go routines are cheap so we can start hundreds of thousands of go-routines.
- Go routines are multiplexed to a fewer number of OS threads. There may be only one thread which can handle 1000s of go routines.
- Scheduling of go routines is managed by Go runtime. So, it's faster. For threads, scheduling is done by OS runtime. So, context switching time for go-routines is much faster.
- Go routines communicate using channels. Channels do not cause race condition when using shared memory. This is a powerful construct built into the language. It can be thought of bytes using with Go routines communicate with each other.

### WaitGroups

During the usage of go-routines, we saw that main go-routine was terminating before the other go-routines completed or even began their execution. To wait for all go-routines to finish, we can use wait group. A wait group is a construct for synchronization that allows multiple go-routines to wait for each other. There is package `sync` for this and it acts like a counter that keeps information about spawn go routines and blocks execution in a structured manner until its internal counter becomes zero. That is all go-routines has finished. The syntax for creating a waitgroup looks like below.

```go
var wg sync.Waitgroup
```

We can call various methods on the wait group. `Add(int)` method keeps track of number of go-routines to wait for. This acts like a counter. `Wait()` method blocks the execution of the code until the counter reduces to zero. `Done()` method decreases the internal counter in the `Add` method by 1. Each go routine calls `Done()` method which reduces the counter by 1. [Example of waitgroup](src/30-waitgroups.go) can be seen with usage of it. It executes in about 1 seconds but the output is not in specific order because go routines are not deterministic. We pass the waitgroup reference for go-routines to use them.

### Channels

Channels are a way through which different go-routines communicates. It's like a data exchange programming idea which allows us to exchange data between different parts of go-routine. Channels make concurrent programming easier. Traditional threading models were communicating by sharing memory, however channels try to avoid that because that can result in nasty bugs. These memories are locked by threads and there is possiblity of deadlock and thread contention over that data. Go routines and channels allow access to data using channels. Channels share memory by communicating which ensures that at a given time, only one routine has access to the memory. The communication in the channels is bi-directional, that is you can send and receive values from the same channel. By default, channels send and receive until the other side is ready. This way go-routines synchronize without the use of locks or conditional variables. Each channels holds data of particular data type. Syntax for this looks like below. We use `chan` keyword to declare channel of specific type. For example, below is example of string data type channel.

```go
var c chan string
c := make(chan string)
```

Channels can be bi-directional or single-directional. Channel has several operations.
- send a value using `<-` channel send operator. The value of `v` must be assignable to channel `ch`.
- receive a value using `<-ch` operator.
- closing a channel using `close(ch)`
- querying buffer of a channel
- querying length of a channel using `len(ch)` returns integer denoting the length of a channel

### Unbuffered Channel
```go
ch <- v // send v value to channel ch
val := <-ch // receive value from a channel
close(ch) // close a channel
cap(ch) // returns integer denoting the buffer of a channel
len(ch) // find length of a channel 
```

The data sent in channel can be received only once in any of the go routines. By default, when we created channel, it creates unbuffered channel by default. This means it's going to block the execution of go routine sending the data until the other go routine has received previous value.
[Example of two channels for communication](src/31-channels.go) between go routines explains how this works.

The receive on a channel is blocked until there is another go routine sending data into that channel. A channel that needs a receiver as soon as a message is emitted to the channel is called Unbuffered channel. We do not need to declare the capacity for these kinds of channel because they cannot store any data. For unbuffered channel, the length is always zero.

### Buffered Channel

It will have capacity to hold data. In buffered channel, sending data into go-routine is blocked only if the buffere is full and receiving from a channel is blocked only when channel is empty.

```go
c := make (chan <data_type>, capacity)
c := make(chan int, 10)
```

To find the currently available data in the channel, we can use the `len()` function to get length of the channel. These are number of elements queued in the channel. Length of the channel will never be higher than the capacity of the channel. For an unbuffered channel, the length is always zero.
[Buffered Channel example](src/32-buffered-channel.go) shows how we can send 3 elements without receiving any data. If we add one more element into channel, it will exceed the capacity and this will cause deadlock situation. To avoid this situation, we have to call `receive` before sending 4th element. Similarly, if we remove all `send` operations into channel and directly call `receive` routine, it will cause deadlock because the channel is empty.

We need to close the channel to make sure no more data can be sent to the channel. This is done when we do not want to send any more data to the channel. To [close a channel](src/33-closing-channel.go), we can use `close()` function. While receiving, we can check if the channel is closed by assigning second variable when receiving. This will have boolean values. If `ok` has `true`, it means channel is open and if `ok` is false, channel is closed and we will not receive any more values.

```go
v, ok := <- ch
```

### Panic

In Go, panic is like an exception arising at runtime. Due to this execution of program is terminated. While working with channels, we may notice panic, [for example](src/34-panic.go), when sending to a channel once it has been closed or closing an already closed channel.

### Using `range` for Channel receive

We can use [`for...range` expression to receive data](src/35-channels-for-range.go) from the channel. In this case, we must ensure to close the channel. If channel is not closed, for range never finishes until channel is closed and it will result in panic situation.

### Select Statement

The select statement is similar go switch statement but this is specifically for channels. The select statement makes go-routine wait on multiple communication operations. In this, each of the case statement waits for a send or receive operation from a channel. The select statement blocks until any of the case statements are ready and if there is a situation when there are multiple case statements ready, then it selects one randomly and proceeds.

```go
select {
    case channel_send_receive:
        // statements
    case channel_send_receive:
        // statements
    default:
        // optional block    
}
```

This select statement lets go routine wait on multiple communication operations. Select with channels and go-routine becomes a powerful tool for managing synchronization and concurrency. For example, in the [given example](src/36-select-statements.go), the output is non-deterministic. It can be "One" or "Two". The default case will be executed if none of the case statements have send or receive operation ready. This helps make select statement non-blocking. If all case statements are blocked, the default statement will be executed. Similar to `switch..case`, we can use `break` statement in the case blocks to terminate the select statement.

Switch is non-blocking whereas select statements can block as they are using channels. Switch is deterministic because it goes in sequence whereas select statement is not deterministic.

## Best Practices

Whenever we launch a go-routine function, we must make sure it will eventually exit. if a go-routine would never terminate, it will occupy the memory reserved by the routine forever. This kind of memory leak is called *go-routine leak*. They leak if they end up either blocked forever on IO like channel communication or fall into infinite loops.

```go
func main() {
    var wg = sync.WaitGroup
    wg.Add(2)
    go leakingRoutine(&wg)
    wg.Wait()
}

func leakingRoutine(wg *sync.WaitGroup) {
    ch := make(chan int)
    go func() {
        value := <- ch        
        fmt.Println("Received", value)
        wg.Done()
    }()
    fmt.Println("Finished go routine")
    wg.Done()
}
```

### Spinning up Go-routine inside Closure

A closure is [a function defined inside another function](src/37-incorrect-closure.go). When closure is called, it has access to outer function's local variables.
The start is always delayed by Go routine scheduler and we may not even have CPU available to run go-routine. By the time Go routine starts, the value of variable might have changed. It picks the value of the variable at the time go routine starts. So, never use closure to start go routine this way. If we want to pass outer variable inside go routine, we have to pass the variable explicitly as a method parameter. The output is still unordered because go routines do not execute in sequence.

```go
func main() {
    var wg sync.WaitGroup
    wg.Add(10)
    for i := 1; i <= 10; i++ {
        func () {
            fmt.Println(i)
            wg.Done()
        }
    }
    wg.Wait()
    fmt.Println("Finished")
}
```

When to use buffered or unbuffered channels. Unbuffered channels are easy to follow whereas Buffered channels require size decision upfront. For using buffered channel, we have to handle the situation where the channel is blocking. This could be due to waiting for sender or receiver. These are useful when we know following.
- how many go-routines we have launched 
- we want to limit the number of go-routine we can launch 
- we want to limit the amount of work that can be queued up.

Most programs must return a response within certain amount of time. Hence, we can timeout the running code if it doesn't respond in specific amount of time. In select statements, we can block the timeout period using `time.After()` function. The `After()` function waits for a specified duration `d` to finish and return current time on a channel. The function signature looks like this.

```go
func After(d Duration) <- chan Time
```

[Blocking timeout in Select](src/38-select-timeout.go) is an example code.

## Modules, Packages and Repos

1. Repository: place in VCS where code is stored.
2. Module: The root of the go module or library stored in repository. We can store more than one module in a single repo but not advisable. Every module is versioned so having same version for two modules is not actually correct. To make a module reusable, we have to declare that the project is a module. Every module has a globally unique identifier. Go modules can be downloaded from Github repository or private repository but it has to be globally unique. A module specifies the dependencies needed for it and go version in the `go.mod` file.
3. Packages: Golang code is put in packages and packages are grouped to form modules.

```shell
mkdir module01 && cd module01
go mod init <module_path>
go mod init example.com/module01
```

This `go.mod` tracks dependencies. The directory module01 becomes a module as it has `go.mod` file. This file starts with module and module's unique path. This is the path that will be used when we want to use this as dependency. This also specifies minimum go version required.

```shell
go mod tidy # add any missing modules
go run main.go
```

Once we run `go mod tidy` it will download required dependencies and also indirect dependencies. This will update the `go.mod` file accordingly. A collection of Go source code becomes module when we have a valid `go.mod` file in the root directory. This file lists module declaration, minimum compatible version of Go and dependencies for the imported third-party packages.

```shell
mkdir cryptit && cd cryptit
go mod init github.com/piyushpatel2005/cryptit
```

Go uses capitalization to identify if package level identifier is accessible outside of the package. An identifier whose name starts with uppercase is accessible outside the package. Idenfiers with lowercase or starting with underscore can only be accessible inside the package. Anything that is exported is part of the package API. So, be sure to verify that you intend to export something. Import statement allows us to use exported functions, variables and constants in another package. It cannot be used without importing it.

We can use the packages in the current package easily by specifying their full url in `import` statement. If we want to import this new module into another module such as `module01`, we have to specify similar path and run `go mod tidy` command. This produces error `cannot find module providing package`. Until we publish our modules to github, we cannot import it in another module. To import a local module, we can use `replace` directive.

```go
go mod edit -replace github.com/piyushpatel2005/cryptit=../cryptit
go mod tidy
```

You will see corresponding entry into `go.mod` file.

### Go commands

- `go mod init`: Initializes and writes a new `go.mod` file in the current directory. It also accepts optional argument for module path that we are creating.
- `go mod tidy`: ensures that `go.mod` file matches source code in the module. IT will add any missing module requirements necessary to build the current module's packages and dependencies. This also removes requirements on modules that dont provide relevant packages.
- `go run <go_file>`: Compiles and runs the program. Internally, it will compile and build executable in temporary location, launches that executable and finally cleans it once app exits. This is very useful for testing of code interactively. This acts verify much like an interpreted language.
- `go build`: compiles the packages named by import path along with their dependencies into executable. The executable gets created in the current source directory.
- `go install`: This will compile the code and move the executable to `$GOPATH/bin` directory so that we can run this command from anywhere in terminal. We can check GOPATH using `go env GOPATH`.
- `go get`: resolves the command line arguments to packages at specific module versions. This also updates `go.mod` file to require those versions and downloads source code into the module cache. This is used to add a dependency for a package or upgrade it to the latest version using `go get github.com/package` or to specific version using `go get github.com/package@v1.2.2`.

### Compiling and Installing application

```shell
go build # compile into executable
./cryptit # produces desired output
```

This requires us to be in the directory. If we run it from any other location in terminal, we have to specify full path of the application. In order for us to be able to run this application from anywhere, we have to add GOPATH to PATH variable. To find GOPATH, we can use `go env GOPATH` and then export this in the PATH variable using `export PATH=$PATH:<GOPATH>`. Again, we have to run `go build` of application, we can check in `GOPATH/bin` if there are any binaries in there. We can use `go install` and it will install the binary in `GOPATH/bin`. Now, we can run the application from any other directory in the terminal.

```shell
go env GOPATH
export PATH=$PATH:GOPATH
go build
ls $GOPATH/bin
go install
./cryptit # from any location in terminal.
```

## Publishing a module

While developing a module, we should design and code the packages that module will include. We commit the repo using Go conventions that ensure it's available to other Go tools. Publish the module to make it discoverable by developers. As we update code, make sure to revise the module with versions. The version number should specify backward compatability.

When designing a module's public API, keep functionality focused and ensure backward compatability.
We can publish module by tagging code in your repository to make it available for other developers. Developers can use Go tools like `go get` command to download module's code to use in their code. After we've published module and someone has fetched it with Go tools, it will become visible on the Go package discovery website `pkg.go.dev`. This is where developers can search for modules and read respective documentation.

For **Versioning**, we use semantic versioning system. When we update the code, we modify the version number which will specify its backward compatability and stability. We use module version number by tagging module's source in the repository with the number. In this model of versioning, the first number indicates major version which means if a major version is changed, it may not backward compatible. It does not guarantee backward compatability Second part reflects minor version which adds potentially new APIs, but still largely backward compatible. The last part is patch version which is like bug fix or security fix. This also guarantees backward compatability. The fourth part is optional and may include pre-release identifier like `alpha.1` or `beta.5`, etc. This signals that this is pre-release version and the API may likely change as it evolves.

For publishing our newly developed module, we can use below.

```shell
cd src/cryptit
git init
git remote add origin https://github.com/piyushpatel2005/cryptit.git
git remote -v
git add .
git commit -am "Cryptit module created"
git tag v0.1.0
git push origin v0.1.0
```

Now, we can remove the `replace` directve in `module01/go.mod` file along with `require` declaration. We can add the published module using `go get github.com/piyushpatel2005/cryptit@v0.1.0`. This will add entry into `go.mod` file and we can run the application using `go run main.go` in module01.

While naming packages, we should name them descriptively. The names should possibly with single word. Also, the functions defined inside the package should not include the package name. For example, it should have names like below. There is no point of including package name `string` in the method name like `formatString`.

```
string 
    format
    print
    read
```

As a general naming rules, you should always try to make the name of the package match the name of the directory that contains the package. This helps in easily locating respective source code. It's not absolutely necessary and we have been using `main` package without this name for directory. We cannot import `main` package so this does not cause confusion.

### Best Practices

In Go, documentation is important and should be always incorporated. Go makes writing documentation very easy and accessible. We use `go doc` to access documentation for any module. It parses go source code, including comments and produces documentation in HTML format. This allows us to write documentation as comments and is always part of the actual source code. To document anything (a type, variable, constant, function or package) in Go, we just write a comment before the declaration of that component in source code file. Go doc will parse that as a documentation for that type or component. We can also use a blank comment to create new paragraphs in documentation. For lengthy comments, we can put the comments in a file `doc.go`.

For viewing the documentation for a package, we can use `go doc package_name` and similarly to view documentation for an identifier defined in the package, we can use `go doc <package_name>.<identifier>`.

```shell
go doc decrypt
go doc decrypt.Nimbus # this displays doc for Numbus function
```

When importing packages, we may find two packages with same name. In this case, we may want to rename one of the package, to avoid name collision.

```go
import (
    crand "crypto/rand"
    "math/rand"
)
```

In above code, we cna use `crypto/rand` by using `crand`. This is like a package alias.

## Important Packages

Go is made of packages and they allow for modular code. Packages allow us to have reusable code with an abstraction. Abstraction allows us to reuse existing code without knowing the details of how that is developed. When the code is written by someone else, we can use them by adding a dependency in our project. These are also referred to as third party packages. When packages come as a part of programming language itself, you do not need to add any third party dependency in your project. 

In Golang, the built in packages come as a library known as standard library. These are group of packages bundled together. It contains various packages some of which we will see in this section.

### `strings` package

This package provides several useful methods to make it easier to [work with string data types](./src/).

- `Count(s, substring) int`: used to count number of occurrences of `substring` in `s`. If `substring` is empty string, it will return 1 + number of characters in `s`.
- `Contains(s, substring) bool`: used to find if a `substring` is present in string `s`.
- `ToUpper(s) string`: convert to uppercase
- `ToLower(s) string`: convert to lowercase
- `HasPrefix(s, prefix) bool`: checks whether the string `s` starts with `prefix`.
- `HasSuffix(s, suffix) bool`: checks whether the string `s` ends with `suffix`.
- `Index(s, substr) int`: returns the index of the first instance of `substr` in `s`. If it does not occur, return -1.
- `LastIndex(s, substr) int`: returns the index of the last instance of `substr` in `s` otherwise -1.
- `Compare(a, b) int`: compares two strings and returns integer. If `a == b` return 0, if `a < b` return -1. If `a > b` return +1.
- `Replace(str, old, new, n) string`: returns new string by replacing `n` occurrences of `old` with `new`. If `n < 0`, it will replace all occurrences. Strings are immutable so cannot be modified once created.
- `ReplaceAll(str, old, new) string`: returns new string with all `old` substring replaced by `new` in original string `str`.
- `Trim(s, cutset string) string`: trims leading and trailing characters from `cutset` removed.
- `Split(s, sep string) []string`: slices `s` into substrings separated by `sep` and returns a slice of substrings
- `Join(elems []string, sep string) string`: joins the elements of slice `elems` by using separator `sep` and returns as a string