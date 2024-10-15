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

## Embedding Interfaces

Interfaces can also be embedded in other interfaces. This is similar to embedding structures in Golang as you will see in upcoming tutorials. This is useful when we want to extend an interface with additional methods. In Go, the interfaces are meant to be relatively small and simple. This is why embedding interfaces is not used as much as embedding structures.

```go
package main

import  "fmt"

type Walker interface {
	Walk()
}

type Talker interface {
	Talk()
}

type Flyer interface {
	Fly()
}

type Human struct {
	Name string
	Age  int
}

func (h Human) Walk() {
	fmt.Println(h.Name, "is walking")
}

func (h Human) Talk() {
	fmt.Println(h.Name, "is talking")
}

type Bird struct {
	Name string
}

func (b Bird) Fly() {
	fmt.Println(b.Name, "is flying")
}

func (b Bird) Talk() {
	fmt.Println(b.Name, "is talking")
}

func (b Bird) Walk() {
	fmt.Println(b.Name, "is walking")
}

func main() {
	h := Human{"John", 25}
	h.Walk()
	h.Talk()

	b := Bird{"Pigeon"}
	b.Fly()
	b.Talk()
	b.Walk()
}
```

In this case, I've defined three interfaces `Walker`, `Talker` and `Flyer`. Next, I've defined two structs `Human` and `Bird`. Both of them implement these interfaces. `Human` implements `Walker` and `Talker` interfaces. `Bird` implements `Flyer`, `Talker` and `Walker` interfaces. 

If I wanted I could define a new interafce `Humanoid` which embeds `Walker` and `Talker` interfaces. It can also have additional methods if needed.

```go
type Humanoid interface {
	Walker
	Talker
	Study()
}
```

With this, `Human` struct would need to implement `Study` method. Here, `Humanoid` has `Walker`, `Talker` interface embedded in it. So, any `Humanoid` will have to implement their methods `Walk`, and `Talk`. Additionally, `Humanoid` also defines new method `Study` which any type implementing `Humanoid` will have to implement.

```go
package main

import "fmt"

type Walker interface {
	Walk()
}

type Talker interface {
	Talk()
}

type Flyer interface {
	Fly()
}

type Humanoid interface {
	Walker
	Talker
	Study()
}

type Human struct {
	Name string
	Age  int
}

func (h Human) Walk() {
	fmt.Println(h.Name, "is walking")
}

func (h Human) Talk() {
	fmt.Println(h.Name, "is talking")
}

func (h Human) Study() {
	fmt.Println("Teacher is teaching", h.Name)
	fmt.Println(h.Name, "is studying")
}

type Bird struct {
	Name string
}

func (b Bird) Fly() {
	fmt.Println(b.Name, "is flying")
}

func (b Bird) Talk() {
	fmt.Println(b.Name, "is talking")
}

func (b Bird) Walk() {
	fmt.Println(b.Name, "is walking")
}

func teach(h Humanoid) {
	h.Study()
}

func main() {
	h := Human{"John", 25}
	h.Walk()
	h.Talk()
	teach(h)

	fmt.Println("=====================================")

	b := Bird{"Pigeon"}
	b.Fly()
	b.Talk()
	b.Walk()
	//teach(b) // This will throw an error
}
```

The cool thing with interfaces is that if you now try to pass `Bird` into `teach` function, it will throw an error. This is because `Bird` does not implement `Study` method which is required by `Humanoid` interface.

```output{ lineNos=false }
cannot use b (variable of type Bird) as Humanoid value in argument to teach: Bird does not implement Humanoid (missing method Study)
```
