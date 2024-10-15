package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) String() string {
	return fmt.Sprintf("Person(%s: %d: %s)", p.Name, p.Age, p.City)
}

func printPeople(p []Person) {
	for _, person := range p {
		fmt.Println(person)
	}
}

func main() {
	people := []Person{
		{Name: "John", Age: 30, City: "New York"},
		{Name: "Danny", Age: 25, City: "Los Angeles"},
		{Name: "Smith", Age: 35, City: "Chicago"},
		{Name: "Danny", Age: 26, City: "New Jersey"},
		{Name: "Alex", Age: 40, City: "San Francisco"},
	}

	// Sort by age

	fmt.Println("Sort By Age")
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	printPeople(people)

}
