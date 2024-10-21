# Defining Custom Types

You can define your own custom data types using the `type` keyword. This is usually used with another data type called `struct`s.

```go
type <Name> struct {
    field1 <data_type>
    field2 <data_type>
    ... ...
}
```

**Example:**

```go
package main

import "fmt"

// define struct for Point
type Point struct {
	x int
	y int
}

func main() {
	var pt Point // allocates memory and zero value to fields
	fmt.Println(pt)
	fmt.Printf("%+v\n", pt)
	fmt.Printf("%v\n", pt)
}
```

Here, I've defined custom type `Point` which consists of its x and y coordinates. You can create a variable of type `Point` and access its fields. Notice that I have defined the new type with pascal case notation. This way it is exported and can be used outside the package.

**Output:**

```output{ lineNos=false }
{0 0}
{x:0 y:0}
{0 0}
```

Struct assign contiguous blocks of memory for each of its members. So, if we have a struct with 2 fields of `int16`, it will assign 4 bytes plus a padding bytes.

### `new` Keyword

You can also use `new` keyword to create a variable of your custom types. This allocates memory for each of the fields of the struct, assigns them zero values and returns the pointer to the struct.

```go
package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	pt2 := new(Point)
	fmt.Printf("%+v\n", pt2)
	fmt.Printf("%+v\n", *pt2)
}
```

**Output:**

```output{ lineNos=false }
&{x:0 y:0}
{x:0 y:0}
```

The use of `new` keyword is less common because usually you would want to provide initial values to fields. You can initialize structure while defining a new variable of this type.

**Example:**

```go
package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	pt := Point{
		x: 1,
		y: 2,
	}
	fmt.Printf("%+v\n", pt)
}
```

**Output:**

```output{ lineNos=false }
{x:1 y:2}
```

It is also possible to omit the field name and the values will be assigned in the sequence of field names. This is less common due to code readability. So, we can also create `Point` struct using below syntax.

```go
pt := Point{1, 2}
fmt.Printf("%+v", pt) // {x:1 y:2}
```

## Accessing Fields

You can access individual fields of our custom types just like you would do for struct types. After all, it's simply a Go `struct`. You can also modify fields of a structure.

**Example:**

```go
package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	pt := Point{
		x: 1,
		y: 2,
	}
	fmt.Printf("%d, %d\n", pt.x, pt.y)
	pt.y = 1 // modify field of a struct
	fmt.Printf("%d, %d\n", pt.x, pt.y)
}
```

**Output:**

```output{ lineNos=false }
1, 2
1, 1
```

## Passing Custom types to Functions

Passing custom types as parameters is similar to using any other variable type. In Go, there is `math` package which contains math related functions. In below example, we use `Pow` function to calculate power and `Sqrt` function to calculate square-root of a number.

**Example:**

```go
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

func main() {
	var pt Point
	pt3 := Point{
		x: 1,
		y: 2,
	}
	fmt.Printf("%+v\n", pt)
	fmt.Printf("%+v\n", pt3)
	fmt.Printf("%f\n", calcDistance(pt, pt3))
}

func calcDistance(p1 Point, p2 Point) float64 {
	distance := math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2))
	return distance
}
```

**Output:**

```output{ lineNos=false }
{x:0 y:0}
{x:1 y:2}
2.236068
```

You can also pass a structure by reference.

**Example:**

```go
package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	fullName  string
}

func main() {
	p := Person{
		firstName: "John",
		lastName:  "Doe",
		fullName:  "",
	}
	fmt.Printf("%+v\n", p)
	populateFullName(&p)
	fmt.Printf("%+v\n", p)
}

func populateFullName(p *Person) {
	(*p).fullName = p.firstName + " " + p.lastName
}
```

In this case, I am passing a `Person` struct by reference as we can see in the `populateFullName` function definition. It has a parameter which is of type `*Person`.

**Output:**

```output{ lineNos=false }
{firstName:John lastName:Doe fullName:}
{firstName:John lastName:Doe fullName:John Doe}
```

## Comparing Custom Types

You can compare `struct`s using equality operator. This checks the fields for comparison. If all fields of a structure match, they are considered equal.

**Example:**

```go
package main

import "fmt"

type Point struct {
	x int
	y int
}

func main() {
	p1 := Point{1, 2}
	p2 := Point{1, 1}
	p3 := Point{1, 2}
	fmt.Println(p1 == p2) // type and fields should be same
	fmt.Println(p1 == p3)
}
```

**Output:**

```output{ lineNos=false }
false
true
```

## Strong Type Safety in Go

In Go, even if two structures have same fields and types, they are considered different types. This is because Go is strongly typed language. Below code will not compile because `Point` and `Point2` are considered different types.

```go
package main

import "fmt"

type Point struct {
	x int
	y int
}

type Point2 struct {
	x int
	y int
}

func main() {
	p1 := Point{1, 2}
	p2 := Point2{1, 2}
	p1 = p2 // error: cannot use p2 (variable of type Point2) as Point value in assignment
	fmt.Println(p1, p2)
}
```

Even though `Point` and `Point2` have exact same fields and types, they are considered different types. This is because Go is strongly typed language. It clearly separates two types unlike other languages where implicit type conversion is possible.

If you want to really assign `Point2` type to `Point`, you will have to do explicit type conversion.

```go
p1 = Point(p2)
```

## Composition of Custom Types

Structures can also be nested inside other structures. This is useful when you want to group related data together. For readability, you can define structures as custom types and then use them to compose other custom types.

```go
package main

import "fmt"

type Address struct {
    city  string
    state string
}

type Person struct {
    name    string
    age     int
    address Address
}

func main() {
    p := Person{
        name: "John",
        age:  30,
        address: Address{
            city:  "New York",
            state: "NY",
        },
    }
    fmt.Println(p)
}
```

## Define Structure Methods

In Go, you can have methods which are defined differently. A method is something that augments a function with an extra parameter that specifies the receiver. This argument is added right after `func` keyword and it accepts only one argument. This is called a receiver. A method is a function that has a defined receiver. The syntax for defining a method with receiver is as below.

```go
func (r ReceiverType) MethodName(parameters) {
    // method body
}
```

This can help ensure that any instance of `Person` struct will have `name` method available with them. This method is called with a dot operator after the struct variable.

**Example:**

```go
package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	fullName  string
}

func main() {
	p := Person{
		firstName: "John",
		lastName:  "Doe",
		fullName:  "",
	}
	fmt.Printf("%+v\n", p)
	p.name()
	fmt.Printf("%+v\n", p)
	fmt.Printf("%v\n", p.name())
}

func (p Person) name() string {
	p.fullName = p.firstName + " " + p.lastName
	return p.fullName
}
```

Here, the struct variable is passed by value. So, it does not modify the original fields of the structure.

**Output:**

```output{ lineNos=false }
{firstName:John lastName:Doe fullName:}
{firstName:John lastName:Doe fullName:}
John Doe
```

## Method Sets

*Method set*s are a set of methods available to a data type and are used to encapsulate functionality to specific data type. Here, we defined two methods `calculateBonus` and `calculateGrossSalary` on structure of type `Employee`.

**Example:**

```go
package main

import "fmt"

type Employee struct {
	name       string
	role       string
	department string
	salary     float64
}

func main() {
	e1 := Employee{
		name: "John Doe", role: "Engineer", department: "Marketing", salary: 10000.0,
	}
	e2 := Employee{
		name: "Jane Doe", role: "Manager", department: "Technology", salary: 100000.0,
	}
	fmt.Printf("Salary = %.2f, Bonus = %.2f, Gross salary = %.2f\n", e1.salary, e1.calculateBonus(), e1.calculateGrossSalary())
	fmt.Printf("Salary = %.2f, Bonus = %.2f, Gross salary = %.2f\n", e2.salary, e2.calculateBonus(), e2.calculateGrossSalary())
}

func (e *Employee) calculateBonus() float64 {
	return 0.10 * e.salary
}

func (e *Employee) calculateGrossSalary() float64 {
	if e.department == "Sales" || e.department == "Marketing" {
		return e.salary + e.calculateBonus() + 0.10*e.salary
	} else {
		return e.salary + e.calculateBonus()
	}
}
```

**Output:**

```output{ lineNos=false }
Salary = 10000.00, Bonus = 1000.00, Gross salary = 12000.00
Salary = 100000.00, Bonus = 10000.00, Gross salary = 110000.00
```