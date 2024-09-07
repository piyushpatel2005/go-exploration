# Structures

Structures are one of the ways to organize data into structured data structure. This is also the data structure which helps avoid having objects and classes because we can define the blueprint of data using `struct`s.

## Overview

Structure is a user-defined data type which groups several data elements for related entity. This allows to store multiple values of different data types under single variable name. For example, a person may have name, address, citizenship, security number, etc. So, we can create a struct for Person and it can include all these information for a person under single variable.

## Define and Initialize Structure

You can define a `struct` using below syntax.

```go
struct {
	field1 data_type1
	field2 data_type2
	... ...
}
```

For example, below code defines a struct type variable called `example` and later assigns values to its fields.

```go
package main

import "fmt"

func main() {
	var example struct {
		name string
		age  int
	}

	example.name = "John"
	example.age = 30
	fmt.Println(example)
}
```

You can also use struct literal to define and initialize a struct in one line.

```go
package main

import "fmt"

func main() {
    example := struct {
        name string
        age  int
    }{
        name: "John",
        age:  30,
    }
    fmt.Println(example)
}
```

## Accessing Structure Fields

You can access structure fields using `.` operator. For example, to access `name` field of `example` struct, you can use `example.name`.

```go
package main

import "fmt"

func main() {
    var example struct {
        name string
        age  int
    }

    example.name = "John"
    example.age = 30
    fmt.Println(example.name) // John
	fmt.Println(example.age)  // 30
}
```

The structures are really the most important data type in Go. They allow us to define our own data types by combining one or more types. We can use structures to define complex data types that are made up of multiple fields. In Object-oriented programming, you would use classes to define custom data types. In Go, you can achieve similar functionality using structures as you will see in next lesson.