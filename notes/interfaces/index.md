# Interfaces

In Golang, there are primarily two ways to handle encapsulation. These are structures and interfaces. Interfaces define method sets available for a struct. This helps with featues like Polymorphism in object oriented programming languages.

- Interfaces specify method sets and is used to create modularity in Go programming language. 
- It is like a blueprint for a method set. This is defined so that it can be implemented by other types like Structures. 
- Interfaces only provide method signatures for each method of a method set. They do not implement these methods. These methods are then implemented by other types. 
- It cannot have any variable declarations. They can only have method signatures.

## Defining Interfaces

Interface is defined in Golang using below syntax.

```go
type <interface_name> interface {
    // method signatures
}
```

**Example:

```go
type Account interface {
    withdrawMoney(int amount) bool
    depositMoney(int amount) bool
    getBalance() float64
}
```

## Interface Example using Shape

Interfaces are defined so that other types can implement them. In order to implement an interface, a type has to implement all of interface's methods. There is no `implements` keyword like Java or C#. These are implemented implicitly and does not use specific keyword to implement an interface. This is how we can implement features like polymorphism from other object oriented programming languages.

**Example:**

```go
package main

import "fmt"

const PI float64 = 3.14

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

func (c Circle) area() float64 {
	return PI * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * PI * c.radius
}

func main() {
	square := Square{side: 2}
	circle := Circle{radius: 2}
	printDetails(circle)
	printDetails(square)
}

func printDetails(s Shape) {
	fmt.Printf("Area = %.2f, Perimeter = %.2f\n", s.area(), s.perimeter())
}
```

In this case, we define `interface` of type `Shape`. Further, we define two structs of type `Circle` and `Square`. Both of them have methods defined `area` and `perimeter`. These methods calculate correct values based on type of the shape. Here `Square` and `Circle` implicitly implements `Shape` interface. We do not use any explicit keyword for this implementing this.

**Output:**

```output{ lineNos=false }
Area = 12.56, Perimeter = 12.56
Area = 4.00, Perimeter = 8.00
```