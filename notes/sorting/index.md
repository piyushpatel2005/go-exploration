# Sorting in Go

Sorting is a common operation in programming. For example, when presenting your products sorted in the order of their price or when you want to sort based on age of the item, etc. In Go, you have a `sort` package which allows you to sort slices and other `comparable` data types. 

## Sorting Slices

For sorting integer slices, you could use `Ints` function from `sort` package. 

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{4, 2, 3, 1}
	fmt.Printf("Before sorting: %v\n", numbers) // Before sorting: [4 2 3 1]
	sort.Ints(numbers) 
	fmt.Printf("After sorting: %v\n", numbers) // After sorting: [1 2 3 4]
}
```

The `Ints` function sorts the slice in ascending order. Similarly, you could sort a slice of strings using `Strings` function which again by default sorts in ascending order. 

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
	names := []string{"apple", "banana", "orange", "grapes"}
	fmt.Printf("Before sorting: %v\n", names) // Before sorting: [apple banana orange grapes]
	sort.Strings(names)
	fmt.Printf("After sorting: %v\n", names) // After sorting: [apple banana grapes orange]
}
```

## Custom Sorting

If you wan tto implement custom sorting, you have to implement the `sort.Sort` interface. This interface has three methods `Len`, `Less` and `Swap` which you may need to implement. 

- `Len` returns the length of the slice. For example, if you have a slice of `Person` struct, you could return the length of the name.
- `Less` returns a boolean value based on the comparison of two elements. For example, if you want to sort based on the age of the person, you could compare the age of two persons and return the result.
- `Swap` swaps the elements at two indices. 

In this hypothetical example, I want to sort a slice of `Person` struct based on their name and if their names are same, then sort based on their age in descending order.

I can first define the `Person` type along with the `People` type which is a slice of `Person`. For each of those types I have implemented `String` method which returns a string representation of the type. For `People` type, I need to import `strings` package to join the strings.

```go
type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d: %s", p.Name, p.Age, p.City)
}

type People []Person

func (p People) String() string {
    strs := make([]string, len(p))
    for i, person := range p {
        strs[i] = fmt.Sprintf("Person(%s: %d: %s)", person.Name, person.Age, person.City)
    }
    return fmt.Sprintf("[%s]", strings.Join(strs, ", "))
}
```

Next, I can implement the `sort.Sort` interface for `People` type. 

```go
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
```

Now, if I try to sort a slice of `People`, it will sort based on the custom logic I have implemented. 

```go
func main() {
    people := People {
        {Name: "John", Age: 30, City: "New York"},
        {Name: "Danny", Age: 25, City: "Los Angeles"},
        {Name: "Smith", Age: 35, City: "Chicago"},
        {Name: "Danny", Age: 26, City: "New Jersey"},
        {Name: "Alex", Age: 40, City: "San Francisco"},
    }

    fmt.Printf("Before sorting: %v\n", people)
    sort.Sort(people)
    fmt.Printf("After sorting: %v\n", people)
}
```

### Reverse Sorting
If you want to change the sorting order to descending, you could use the `sort.Interface` function along with `sort.Reverse` function.

```go
sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
sort.Sort(sort.Reverse(sort.StringSlice(names)))
sort.Sort(sort.Reverse(people))
```

## `Slice` function

If you want to sort a slice based on a custom function, you could use the `Slice` function from `sort` package. This function takes a slice and a function which returns a boolean value. The function should return `true` if the element at index `i` should come before the element at index `j`. 

```go
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
```

## Sort Wrapper

You can define your own wrapper type for sorting. This is useful when you want to sort a slice of a type which is not directly comparable. For example, if you want to sort a slice of `Person` based on their age, you could define a `ByAge` type which is a slice of `Person` and implement the `sort.Interface` for it. 

```go
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

type ByAge []Person

func (a ByAge) Len() int { return len(a) }

func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person

func (a ByName) Len() int { return len(a) }

func (a ByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

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
	sort.Sort(ByAge(people))
	printPeople(people)

	fmt.Println("=====================================")
	// Sort by name
	fmt.Println("Sort By Name")
	sort.Sort(ByName(people))
	printPeople(people)

	fmt.Println("=====================================")

	// Sort by name and age
	fmt.Println("Sort By Name in reverse order")
	sort.Sort(sort.Reverse(ByName(people)))
	printPeople(people)

}
```

You could also reuse the methods `Len` and `Swap` as they essentially have the same implementation above. For this, you would need to embed `Person` slice in `ByAge` and `ByName` types. With these changes, you simply have to implement the `Less` method for each type. 

```go
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

func (a ByAge) Less(i, j int) bool { return a.People[i].Age < a.People[j].Age }

type ByName struct{ People }

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
	sort.Sort(ByAge{people}) // pass people as struct
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
	// Sort by Name ascending and Age descending
	fmt.Println("Sort By Name in descending order and Age in ascending order")
	sort.Sort(ByName{people})
	sort.Sort(sort.Reverse(ByAge{people}))
	printPeople(people)
}
```