---
title: "Structures in Golang"
date: 2023-08-12T20:38:44Z
draft: false
---

Structures are one of the ways to organize data into structured data structure. This is also the data structure which helps avoid having objects and classes because we can define the blueprint of data using `struct`s.
<!--more-->

## Overview

Structure is a user-defined data type which groups several data elements for related entity. This allows to store multiple values of different data types under single variable name. For example, a person may have name, address, citizenship, security number, etc. So, we can create a struct for Person and it can include all these information for a person under single variable. 

## Define and Initialize Structure

To initialize a structure we can use below syntax. 

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

**Output:**

```output{ lineNos=false }
{0 0}
{x:0 y:0}
{0 0}
```

Struct assigns contiguous blocks of memory for each of its members. So, if we have a struct with 2 fields of `int16`, it will assign 4 bytes plus a padding byte.

### `new` Keyword

We can also use `new` keyword to create a variable of struct type. This allocates memory for each of the fields of the struct, assigns them zero values and returns the pointer to the struct. 

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

The use of `new` keyword is less common because usually we want to provide initial values to fields. We can initialize structure while defining a structure.

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

## Accessing Struct Fields

We can access individual fields of the struct using dot operator (`.`). We can also modify fields of a structure.

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

## Structures as Function Argument

Passing structures as parameters is similar to using any other variable type. In Go, there is `math` package which contains math related functions. In below example, we use `Pow` function to calculate power and `Sqrt` function to calculate square-root of a number.

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

We can also pass by reference.

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

In this case, we are passing a `Person` struct by reference as we can see in the `populateFullName` function definition. It has a parameter which is of type `*Person`.

**Output:**

```output{ lineNos=false }
{firstName:John lastName:Doe fullName:}
{firstName:John lastName:Doe fullName:John Doe}
```

## Comparing Structures

We can compare `struct`s using equality operator. This checks the fields for comparison. If all fields of a structure match, they are considered equal.

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

## Define Structure Methods

In Go, we have methods which are defined differently. A method is something that augments a function with an extra parameter that specifies the receiver. This argument is added right after `func` keyword and it accepts only one argument. This is called a receiver. A method is a function that has a defined receiver. This ensures that any instance of `Person` struct will have `name` method available with them. This method is called with a dot operator after the struct variable.

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