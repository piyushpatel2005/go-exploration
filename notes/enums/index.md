# Enums in Go

Go doesn't have some of the common constructs available in other programming langauge. This makes the language relatively easier to learn and understand. One of the constructs that Go doesn't have is enums. Enums are a way to define a set of named integral constants that may be assigned to a variable. This is a way to define a set of named constants of a specific type.

Although Go doesn't have it out of box, we can still create similar structure using `const`. Let's see how we can do that in multiple different ways. In this case, I want to define the days of the week as an enum.

## Using `const` and `int` type

In order to define a custom data type, you can use `type` keyword. I'm going to represent the day as an integer.

```go
type Day int
```

Now, to define different days, we can use `const` keyword.

```go
package main

import "fmt"

const (
    Sunday int = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func main() {
	fmt.Println("sunday is ", Sunday)
}
```

You can also define a method which will return the weekday number. That is Sunday is 0, Monday is 1 and so on.

```go
func findDayNumber(d Day) int {
	return int(d) - 1
}
```
In this case, I have to convert the `Day` type to `int` type otherwise I cannot perform arithmetic operations on it. So, Go doesn't do implicit type conversion. Using this, you can get the day number as below.

```go
func main() {
	fmt.Println("sunday is ", findDayNumber(Sunday))
}
```

## Using `const` and `iota` type

You can simplify this declaration of enum by using `iota` keyword. `iota` is a predeclared identifier representing the untyped integer ordinal number of the current const specification in a `const` declaration. It is reset to 0 whenever the reserved word `const` appears in the source code. So, with this you can define the days of the week as below.

```go
package main

import "fmt"

type Day int

const (
    Sunday Day = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func findDayNumber(d Day) int {
	return int(d)
}

func main() {
    fmt.Println("sunday is ", Sunday)
}
```

Notice that in this case, you just have to `iota` for the first value and rest of the values will be automatically assigned the next value. This is a very simple and easy way to define enums in Go. Also because it defines first value as 0, you don't have to do any arithmetic operations to get the day number in `findDayNumber()` function.

## Assigning Methods to Custom types

You can also assign methods to custom types. This is a very powerful feature of Go and helps you implement object-oriented code structures. You can define a method on any type that is defined in the same package. This is a way to define methods on custom types. Let's see how we can define a method on `Day` type. In this case, I'm implementing `fmt.Stringer` interface and providing string representation of the day. So, whenever you try to print the `Day` type, you will get the string representation of the day which will be lot more readable.

```go
package main

import "fmt"

type Day int

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d Day) String() string {
	switch d {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	default:
		return "Unknown"
	}
}

func (d Day) isWeekend() bool {
	switch d {
	case Saturday, Sunday:
		return true
	default:
		return false
	}
}

func (d Day) dayNumber() int {
	return int(d)
}

func main() {
	fmt.Println("Sunday is ", Sunday)
	fmt.Println("Monday is day number: ", Monday.dayNumber())
	fmt.Printf("Is %s a weekday? %v\n", Sunday, Sunday.isWeekend())
}
```

Notice that in this case, `dayNumber()` and `isWeekend()` is called on the `Day` type. This is a very powerful feature of Go and helps you define methods on custom types. This is a way to define methods on custom types. Similarly, when you print  the `Day` type, you will get the string representation of the day. This is because I have defined a `String()` method on the `Day` type.