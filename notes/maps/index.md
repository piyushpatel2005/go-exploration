# Maps

Sometimes you do not need data to be independent. Usually data is associated with a specific object, like a person's salary or user's profile. In such situation, you can still use array indices to track each object, but that becomes very difficult. This is where Maps become useful. Maps are just a key-value association.

## Overview

Map is unordered collection of key-value pairs. They are usually associated with key. So, they are called associative arrays or dictionary in other languages. They are implemented using hash tables and due to that they provide constant time lookup. So, add, get and delete operations are very fast with maps.


## Creating Maps in Golang
The general syntax for creation of a map looks like this.

```go
var <name> map[<key_data_type>]<value_data_type>
```

**Example of map Declaration:**
```go
var employees map[string]int
```

If we try to assign key-value pairs to uninitialized map, we get an error.

**Example:**
```go
package main

import "fmt"

func main() {
	var employees map[string]int
	employees["Akshay"] = 90000 // throws error
	fmt.Println(employees)
}
```

To create a map with key-value pairs we also have to initialize it with key-value pairs.

**Example:**

```go
package main

import "fmt"

func main() {
	numbers := map[string]int{"one": 1, "two": 2}
	fmt.Println(numbers)
}
```

**Output:**
```output{ lineNos=false }
map[one:1 two:2]
```

You can find how many key value pairs exist in a map using `len` function.

**Example:**
```go
package main

import "fmt"

func main() {
	numbers := map[string]int{"one": 1, "two": 2}
	fmt.Println(len(numbers)) // 2
}
```

You can also declare and initialize a map with make function. You can also assign new key-value pairs using assignment operator as shown in line `9` below.

```go
package main

import "fmt"

func main() {
	numbers := make(map[string]int)
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	numbers["one"] = 1 // now we can assign key-value
	fmt.Println(numbers)
}
```

**Output:**
```output{ lineNos=false }
map[]
0
map[one:1]
```

**Output:**

```output{ lineNos=false }
map[1:hi 2:hello 3:there 4: 5:]
```

## Accessing elements of a Map in Go

You can get values for a specific key using its keys. Just like arrays and slices we can use `[]` to access value associated with a key.

**Example:**
```go
package main

import "fmt"

func main() {
	numbers := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(numbers["one"])
}
```

**Output:**

```output{ lineNos=false }
1
```

When retrieving values, we get value and a boolean flag if the key exists or not. If key doesn't exist, the first value will be zero value. This we can use with conditionals to verify if a key exists and if exists, get the value else do something else.

```go
package main

import "fmt"

func main() {
	numbers := map[string]int{"one": 1, "two": 2, "three": 3}
	value, found := numbers["three"]
	fmt.Println(found, value)

	value, found = numbers["four"]
	fmt.Println(found, value)
}
```

**Output:**
```output{ lineNos=false }
true 3
false 0
```

You can also assign new key-value pairs using `map_name[key] = value` syntax as you've already seen. However, if we try to assign the existing key with a new value, it will simply overwrite existing values.

```go
package main

import "fmt"

func main() {
	numbers := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(numbers)
	numbers["one"] = 11
	fmt.Println(numbers)
}
```

In above code, line `8` reassigns the key `one` with new value `11` which can be seen from the below output.

**Output:

```output{ lineNos=false }
map[one:1 three:3 two:2]
map[one:11 three:3 two:2]
```

### Create a Map from an Array

Sometimes, it may be useful to create a map from an existing array, below code shows an example of this. In this case, we are using array indices as map keys. Also, note that additional keys will get zero value for specific data type as their value, for example, keys `4` and `5`.

**Example:**

```go
package main

import "fmt"

func main() {
	arr := [5]string{"hi", "hello", "there"}
	greetings := make(map[int]string, 5)
	for index, elem := range arr {
		greetings[index+1] = elem
	}
	fmt.Println(greetings)
}
```

## Iterating through Maps in Go

For looping through each key-value pair in a map, we can use `range` to get both `key` and `value` of each pairs.

**Example:**

```go
package main

import "fmt"

func main() {
	numbers := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, value := range numbers {
		fmt.Println(key, "has value", value)
	}
}
```

Again, the `map` data-structure is unordered. So, output below sometimes may not match exactly.

**Output:**

```output{ lineNos=false }
one has value 1
two has value 2
three has value 3
```

Alternatively, if you only need key, you can ignore the value by using `_` as shown below.

```go  
package main

import "fmt"

func main() {
	numbers := map[string]int{"one": 1, "two": 2, "three": 3}
	for key := range numbers {
		fmt.Println(key)
	}
}
```

## Deleting a Key from the map

To delete a key-value pair from the map, you could use `delete` function. This function takes two arguments, first is the map and second is the key to be deleted.

**Example:**

```go
package main

import "fmt"

func main() {
    numbers := map[string]int{"one": 1, "two": 2, "three": 3}
    fmt.Printf("%v\n", numbers)
    delete(numbers, "one")
    fmt.Printf("%v\n", numbers)
}
```

**Output:**

```output{ lineNos=false }
map[one:1 three:3 two:2]
map[three:3 two:2]
```

## Clean up Map

For clearing any key-value pair from a map, we can use [built-in function delete](https://pkg.go.dev/builtin#delete).

```go
package main

import "fmt"

func main() {
	numbers := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Printf("%v\n", numbers)
	for key, _ := range numbers {
		delete(numbers, key)
	}
	fmt.Printf("%v\n", numbers)
}
```

**Output:**

```output{ lineNos=false }
map[one:1 three:3 two:2]
map[]
```

## Conclusion

In this lesson, we looked at various operations we can perform on Map data structure in Go. These are basically associative arrays with key-value pairs. We learned how to create a map, get element for a specific key, add new key-value pairs. We also learned how we can iterate over a map or delete map from memory.