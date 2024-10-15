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

type People []Person

func (a People) Len() int { return len(a) }

func (a People) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (p Person) String() string {
	return fmt.Sprintf("Person(%s: %d: %s)", p.Name, p.Age, p.City)
}

type ByAge struct{ People }

//func (a ByAge) Len() int { return len(a) }
//
//func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByAge) Less(i, j int) bool { return a.People[i].Age < a.People[j].Age }

type ByName struct{ People }

//func (a ByName) Len() int { return len(a) }
//
//func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByName) Less(i, j int) bool { return a.People[i].Name < a.People[j].Name }

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
	sort.Sort(ByAge{people})
	printPeople(people)

	fmt.Println("=====================================")
	// Sort by name
	fmt.Println("Sort By Name")
	sort.Sort(ByName{people})
	printPeople(people)

	fmt.Println("=====================================")

	// Sort by name and age
	fmt.Println("Sort By Name in reverse order")
	sort.Sort(sort.Reverse(ByName{people}))
	printPeople(people)

	fmt.Println("=====================================")
	fmt.Println("Sort By Name in descending order and Age in ascending order")
	sort.Sort(ByName{people})
	sort.Sort(sort.Reverse(ByAge{people}))
	printPeople(people)

}
