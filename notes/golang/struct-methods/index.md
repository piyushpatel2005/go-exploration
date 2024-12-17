# Struct Methods

In this tutorial, I explain more about methods in Go. In Go, we don't have concept of classes and objects. Structs are primarily used to create complex data structures. So, we may need to add additional functionalities in the form of methods.

## Function vs Method:

Firstly, let's first understand the difference between functions and methods in Go. Below code shows syntax difference.

```go
type MyStruct struct {
   field int
}

func MyFunction(a myStruct) { }  // This is a function

func (a myStruct) MyMethod() { } // This is a method
```

Functions are pieces of code that can be called from anywhere else in your code while methods are attached to *a particular type* and can be invoked specifically on an instance of that type. This is very important difference. So, function call and method call would look very different as shown in below snippet.

```go
MyFunction(m) // function call
m.MyMethod() // method call
```

Let's take an example of a bank.

```go
package main

import "fmt"

type Account struct {
	number  int
	name    string
	balance float64
}

func CanWithdraw(a Account, amount float64) bool {
	return a.balance >= amount
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.balance >= amount
}

func main() {
	a := Account{number: 1, name: "Bob", balance: 300.0}
	fmt.Println(CanWithdraw(a, 100.0))
	fmt.Println(a.CanWithdraw(100))
}
```

Function needs `Account` type to be passed as an argument whereas methods work on `Account` type. The declaration and invocation syntax differs between these two.

```output{ lineNos=false }
true
true
```

## Define new methods on existing type

You can define a new type, that is basically same as existing type. For example, below code creates a new type `Number` which looks like an alias to `int` type.

```go
package main

import "fmt"

type Number struct {
   digit int
}

func (n Number) SetValue(val int) {
   n.digit = val
}

func main() {
   num := Number{digit: 5}
   num.SetValue(10) // This modifies local value
   fmt.Println(num.digit) // 5 
}
```

This might raise a question, why would you do this? In Go, we cannot define a new method on `int` type. As a matter of fact, in Go, the definition of the receiver type must be in the same package as the new method. Receiver type is the type to which the new method is attached.So, we cannot define any methods on existing types because they are not in the package, we define. If we try to create a method on `int` type, we get compile error.

```go
package main

import "fmt"

// cannot define new methods on non-local type int
func (n int) SetValue(val int) {
	n.digit = val
}

func main() {
	var a int = 4
	fmt.Println(a.SetValue(100))
}
```

## Value Receiver
A method with value receiver gets the copy of the original value that you passed into it, meaning any changes won't affect the original struct instance.

```go
package main

import "fmt"

type Number struct {
   digit int
}

func (n Number) SetValue(val int) {
   n.digit = val
}

func main() {
   num := Number{digit: 5}
   num.SetValue(10) // This modifies local value
   fmt.Println(num.digit) // Outputs: 5 
}
```

In the example above, setting a new value inside method did not change outside variable because it was passed by value.

To fix this, we can definitely return new `Number` type after setting a value and re-assign the value to existing `num` variable like below.

```go
package main

import "fmt"

type Number struct {
	digit int
}

func (n Number) SetValue(val int) Number {
	n.digit = val
	return n
}

func main() {
	num := Number{digit: 5}
	num = num.SetValue(10)
	fmt.Println(num.digit) // 10
}
```

However, this potentially creates unnecessary copy of the variable when we really wanted original `num` to be modified. We can fix this using pointers.

## Pointer Receiver
What if we want modifications inside the function to affect our original struct? Well, in that case, you need to pass pointer to struct instead of value. 

```go
package main

import "fmt"

type Number struct {
    digit int
}

func (n *Number) SetPointerValue(val int){
    n.digit = val
}

func main(){
    num := &Number{digit: 5} // declare it as pointer using &
    num.SetPointerValue(10)
    fmt.Println(num.digit)// 10
}
```

Attaching pointer receiver allows us to modify the value stored in the same memory location.

This works but what if we pass the variable but receiver is of pointer type? 

```go
type Number struct {
    digit int
}

func (n *Number) SetPointerValue(val int){
    n.digit = val
}

func main(){
    num := Number{digit: 5}
    num.SetPointerValue(10)
    fmt.Println(num.digit)// 10
}
```

Above code also produces same results even though we pass `Number` and not `&Number`. When we use pointer receiver, the method accesses the memory location of the stored value rather than value itself. When we update, it will actually update the value stored in the same memory location.

If we modify the code to accept value receiver, it will stop working again even though we can pass pointer type when calling.

```go
type Number struct {
    digit int
}

func (n Number) SetPointerValue(val int){
    n.digit = val
}

func main(){
    num := &Number{digit: 5} // declare it as pointer using &
    num.SetPointerValue(10)
    fmt.Println(num.digit)// 5
}
```

- If function accepts value type, it will only accept value arguments and similarly, if it accepts pointer type, it can only accept reference arguments.
- On the other hand, if a method accepts value receiver, it will accept either pointer and value arguments and if it accepts pointer receiver, it will accept either type of argument.

### Points to Consider when choosing recevier type

1. If your method modifies the receiver, it should be of pointer type.
2. If your method accesses the receiver, it should be of value type.
3. If your method needs to handle `nil` value, it should be of pointer type.