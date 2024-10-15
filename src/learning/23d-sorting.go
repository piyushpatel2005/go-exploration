package main

import (
	"fmt"
	"sort"
	"strings"
)

type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) String() string {
	return fmt.Sprintf("Person(%s: %d: %s)", p.Name, p.Age, p.City)
}

type People []Person

func (p People) String() string {
	strs := make([]string, len(p))
	for i, person := range p {
		strs[i] = fmt.Sprintf("Person(%s: %d: %s)", person.Name, person.Age, person.City)
	}
	return fmt.Sprintf("[%s]", strings.Join(strs, ", "))
}

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	if p[i].Name != p[j].Name {
		return p[i].Name < p[j].Name
	} else {
		return p[i].Age > p[j].Age
	}
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	numbers := []int{4, 2, 3, 1}
	fmt.Printf("Before sorting: %v\n", numbers)
	sort.Ints(numbers)
	fmt.Printf("After sorting: %v\n", numbers)
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Printf("After reverse sorting: %v\n", numbers)

	fmt.Println("=====================")
	names := []string{"apple", "banana", "orange", "grapes"}
	fmt.Printf("Before sorting: %v\n", names) // Before sorting: [apple banana orange grapes]
	sort.Strings(names)
	fmt.Printf("After sorting: %v\n", names) // After sorting: [apple banana grapes orange]
	sort.Sort(sort.Reverse(sort.StringSlice(names)))
	fmt.Printf("After reverse sorting: %v\n", names) // After reverse sorting: [orange grapes banana apple]

	fmt.Println("========== Custom Sorting ============")
	people := People{
		{Name: "John", Age: 30, City: "New York"},
		{Name: "Danny", Age: 25, City: "Los Angeles"},
		{Name: "Smith", Age: 35, City: "Chicago"},
		{Name: "Danny", Age: 26, City: "New Jersey"},
		{Name: "Alex", Age: 40, City: "San Francisco"},
	}

	fmt.Printf("Before sorting: %v\n", people)
	sort.Sort(people)
	fmt.Printf("After sorting: %v\n", people)
	sort.Sort(sort.Reverse(people))
	fmt.Printf("After reverse sorting: %v\n", people)
}
