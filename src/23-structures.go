package main

import (
	"fmt"
	"math"
)

// define struct for Point
type Point struct {
	x int
	y int
}

type Citizen struct {
	firstName string
	lastName  string
	age       string
	address   string
	phones    []string
}

func main() {
	var pt Point // allocates memory and zero value to fields
	fmt.Println(pt)
	fmt.Printf("%+v\n", pt)
	fmt.Printf("%v\n", pt)

	// another way is to use new keyword
	pt2 := new(Point)
	fmt.Printf("%+v\n", pt2)

	// usually the values will be assigned during creation of the variable
	pt3 := Point{
		x: 1,
		y: 2,
	}
	fmt.Printf("%+v\n", pt3)

	// To access fields of the struct, we can use dot operator.
	fmt.Printf("%d, %d\n", pt3.x, pt3.y)

	// Modify fields using dot operator
	pt3.y = 1
	fmt.Printf("%d, %d\n", pt3.x, pt3.y)

	// Passing structures to function
	fmt.Printf("%+v\n", pt)
	fmt.Printf("%+v\n", pt3)
	fmt.Printf("%f\n", calcDistance(pt, pt3))

	// Passing structure by reference
	p := Person{
		firstName: "John",
		lastName:  "Doe",
		fullName:  "",
	}
	fmt.Printf("%+v\n", p)
	populateFullName(&p)
	fmt.Printf("%+v\n", p)

	// Methods
	p = Person{
		firstName: "John",
		lastName:  "Doe",
		fullName:  "",
	}
	fmt.Printf("%+v\n", p)
	// p.name()
	// fmt.Printf("%+v\n", p)
	p.name2()
	fmt.Printf("%+v\n", p)
	fmt.Printf("%v\n", p.name2())

	// Method sets

	e1 := Employee{
		name: "John Doe", role: "Engineer", department: "Marketing", salary: 10000.0,
	}
	e2 := Employee{
		name: "Jane Doe", role: "Manager", department: "Technology", salary: 100000.0,
	}
	fmt.Printf("Salary = %.2f, Bonus = %.2f, Gross salary = %.2f\n", e1.salary, e1.calculateBonus(), e1.calculateGrossSalary())
	fmt.Printf("Salary = %.2f, Bonus = %.2f, Gross salary = %.2f\n", e2.salary, e2.calculateBonus(), e2.calculateGrossSalary())

	// Interfaces

	square := Square{side: 2}
    circle := Circle{radius: 2}
    printDetails(circle)
    printDetails(square)

}

func calcDistance(p1 Point, p2 Point) float64 {
	distance := math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2))
	return distance
}

type Person struct {
	firstName string
	lastName  string
	fullName  string
}

func populateFullName(p *Person) {
	(*p).fullName = p.firstName + " " + p.lastName
}

func (p *Person) name() { // If this is not pointer, it will not modify the value of the fullName field.
	p.fullName = p.firstName + " " + p.lastName
}

func (p Person) name2() string {
	p.fullName = p.firstName + " " + p.lastName
	return p.fullName
}

// Method sets

type Employee struct {
	name       string
	role       string
	department string
	salary     float64
}

func (e *Employee) calculateBonus() float64 {
	bonus := 0.1 * e.salary
	return bonus
}

func (e *Employee) calculateGrossSalary() float64 {
	if e.department == "Sales" || e.department == "Marketing" {
		return e.salary + e.calculateBonus() + 0.05*e.salary
	} else {
		return e.salary + e.calculateBonus()
	}
}

// Interfaces

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

const PI float64 = 3.14

func (c Circle) area() float64 {
	return PI * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * PI * c.radius
}

func printDetails(s Shape) {
	fmt.Printf("Area = %.2f, Perimeter = %.2f\n", s.area(), s.perimeter())
}
