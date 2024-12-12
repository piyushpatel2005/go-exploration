# Pointers

Pointers are one of the most powerful idea in programming, but this is also one of the most complicated structure as well. This is very useful for efficiently manipulating memory values.

## Overview

In this lesson, we will discuss pointers which are used to reference the memory address of a variable or data structure. When a variable is declared, a certain amount of memory is allocated based on the variable data type during the program execution. So, every time we execute the code, we may get different memory location. Pointer is something that holds the memory address of a variable. They do not only hold memory addresses but also point to where the memory is allocated and allows us to modify the value stored in that memory location. The same thing happens for functions as well, they are stored in memory. So, we can referene those functions and use them using pointers.

There are two main operators when working with pointers.
1. **address of operator** (`&`) gives the memory address of a variable. To use this, we precede `&` before the name of the variable.
2. **derefernce operator** (`*`) is used to get the value at a given memory address. This operator is preceded before the memory address.

## Declaring and Initializing a Pointer

The syntax for declaring a pointer variable looks like this.

```go
var <pointer_name> *<data_type>
```

Below we declare pointer variable for various data types.

```go
var ptr_id *int
var ptr_name *string
```

For uninitialized pointers, we get `nil`.

**Example:**

```go
package main

import "fmt"

func main() {
	var a *int
	var str *string
	fmt.Println(a)   // <nil>
	fmt.Println(str) // <nil>
}
```

**Output:**

```output{ lineNos=false }
<nil>
<nil>
```

For assigning values to pointer variable, we use the address of operator.

```go
var ptr_id *int
var ptr_name *string
var id int
var address string
ptr_id = &id
ptr_name = &address
```

Once we have assigned pointer variable, we can use dereference operator to extract value from that pointer.

**Example:**

```go
package main

import "fmt"

func main() {
	var a int = 50
    var a_ptr = &a
	fmt.Printf("%T %v \n", a_ptr, &a)
	fmt.Printf("%T %v\n", *(&a), *(a_ptr)) 
}
```

This code gives a random memory address in the first print statement. We can see the type of pointer as `*int`. Here, we used `a_ptr` as a variable, but we can directly use `&a` to get the same value. I have used both in print statements just for reference.

**Output:**

```output{ lineNos=false }
*int 0xc000016058 
int 50
```

We can also create a pointer variable using built-in [function new](https://pkg.go.dev/builtin#new). This sets the pointer to zero value. So, dereferencing this pointer will give zero value for respective data type.

```go
package main

import "fmt"

func main() {
	new_ptr := new(int) // sets to zero value
	fmt.Println("new_ptr", *new_ptr)
}
```

**Output:**

```output{ lineNos=false }
new_ptr 0
```

## Modify Variable value using Pointer

Pointers are actually reference to memory and if we can create another pointer by dereferencing the same variable, we get the same memory address as you can see in the below example. Also note that if we modify the value at a given address, all pointers referencing that memory address will have new value after the change. This modification using dereference operator is a powerful concept.

**Example:**

```go
package main

import "fmt"

func main() {
	greeting := "Hello"
	var ptr *string
	ptr = &greeting
	fmt.Println(ptr, *ptr)

	var ptr2 *string = &greeting
	fmt.Println(ptr2, *ptr2)

	ptr3 := &greeting
	fmt.Println(ptr3, *ptr3)
	*ptr3 = "Hi"
	fmt.Println(greeting, *ptr, *ptr2, *ptr3)
}
```

In your case, the output below might refer to different memory address.

**Output:**

```output{ lineNos=false }
0xc000014280 Hello
0xc000014280 Hello
0xc000014280 Hello
Hi Hi Hi Hi
```

As you can see in the output above, all pointers refer to same memory address and when we modified the value in that address, it is reflected in all pointers as well as original variable `greeting`.

## Argument - Pass by Value

Now that we are aware of pointers, we can understand argument passing by value. In this case, the parameters are copied into another local variable and different location of the memory. That copy is accessible only within the scope of the function and when we modify that variable, it will modify only the copy and will not change the original variable we passed when calling the function. Go by default uses pass by value for both primitive types as well as structs. You still have the option of passing by reference by using pointers.

**Example:**

```go
package main

import "fmt"

func main() {
	number := 10
	fmt.Println("number =", number)
	increment(number)
	fmt.Println("number =", number) // still remains the same
	input := "hello"
	fmt.Println("input =", input)
	concatenate(input)
	fmt.Println("input =", input)
}

func increment(number int) {
	number += 1 // this modifies the copy of parameter
	fmt.Println("Inside increment function, input:", number)
}

func concatenate(input string) {
	input += " there"
	fmt.Println("Inside concatenate function, input:", input)
}
```

As we can see from the output below, inside the function `increment` the value of `number` variable was changed to `3`, but outside that function, it still has the same value of `2`. This is pass by value where a copy of `number` variable is created inside the `increment` function.

**Output:**

```output{ lineNos=false }
number = 10
Inside increment function, input: 11
number = 10
input = hello
Inside concatenate function, input: hello there
input = hello
```

## Argument - Pass by Reference

On the other hand, when a function parameter is passed by reference, the actual memory address is passed into the function call as a pointer. So, if the function modifies the value of this variable, the value in the calling function also changes.

**Example:**

```go
package main

import "fmt"

func main() {
	input := "hello"
	fmt.Println("input =", input)
	concatenate_ptr(&input)
	fmt.Println("input =", input)
}

func concatenate_ptr(input *string) {
	*input += " there"
}
```

In this case, we are passing `input` variable using their memory address (pointer). This makes it pass by reference. So, when we modify the variable inside the function, it modifies the original value.

**Output:**

```output{ lineNos=false }
input = hello
input = hello there
```

Slices, maps and other complex data types are passed by reference. So, if we modify the slice in the function, it will modify the original slice as well as the underlying original array.

**Example:**

```go
package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4}
	fmt.Println("arr =", arr)
	slice := arr[:3]
	fmt.Println("slice =", slice)
	update_slice(slice)
	fmt.Println("slice =", slice)
	fmt.Println("arr =", arr)
}

func update_slice(slice []int) {
	slice[1] = 200
}
```

**Output:**

```output{ lineNos=false }
arr = [1 2 3 4]
slice = [1 2 3]
slice = [1 200 3]
arr = [1 200 3 4]
```

Maps also, if we modify inside the function, they get updated outside the function scope.

**Example:**

```go
package main

import "fmt"

func main() {
	language := make(map[string]string)
	language["USA"] = "English"
	language["Spain"] = "Spanish"
	language["France"] = "French"
	language["Canada"] = "English"
	fmt.Println(language)
	update_map(language) // update Canada key
	fmt.Println(language)
}

func update_map(m map[string]string) {
	m["Canada"] = "French"
}
```

**Output:**

```output{ lineNos=false }
map[Canada:English France:French Spain:Spanish USA:English]
map[Canada:French France:French Spain:Spanish USA:English]
```

> **Note:** In Go, maps are implemented as a pointer to a struct. This means passing map to a function is passed by reference. 
## Conclusion

In this lesson, we saw how we can use pointers to modify the values of a variable or any memory address. This also explains how we can pass variables by value or by reference. Slice, Maps and other complex data types are, by default, pass by reference.


