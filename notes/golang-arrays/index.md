# Arrays

Often times, it's necessary to store not just single value for any data type, but we want to store a collection of values for a data type. This is where collection data types from a language come in handy. In this lesson, we look at first collection data type called **arrays**.

## Overview

Arrays are a collection of similar data type elements located at contiguous memory locations. Each of these elements are placed next to each other so accessing them sequentially or even randomly is very efficient. Arrays can be of any data type, but they have to have same data type for all elements stored in a single array. Most of the times, we will be using collection of data rather than individual single data. For example, to store information about employee salary, we would need a collection or an array. 

Declaration of an array looks like below.

```go
var <name> <size> <datatype>
```

## Array Declaration and Initialization

To assign an array with list of values, we can use curly braces (`{}`). We can also use shorthand array declaration and initialization.

**Example**

```go
var salaries [5] int
var names [5] string
var names [5]string = [5]string{"John", "Jenny", "Zoe"} // the number of elements on the right should be less than or equal to capacity on the left side.
names2 := [5]string{"John", "Jenny", "Zoe"} 
```

An array has a pointer at the beginning of the array and also has `length` property which represents number of elements in an array and a `capacity` which is the size of the array or the number of elements it can hold. If we know the starting memory address and the data type of an array, we can find memory address of each of the elements in an array. 

By default, uninitialized array contains zero value of the data type.

```go
package main

import "fmt"

func main() {
	// Basic syntax

	var numbers [4]int
	// numbers = [4]int{1, 2, 3, 4}
	fmt.Println(numbers)
}
```

**Output:**
```output{ lineNos=false }
[0 0 0 0]
```

We can initialize the array as shown below.

```go
var numbers [4]int
numbers = [4]int{1, 2, 3, 4}
fmt.Println(numbers)
var numbers2 [4]int = [4]int{1, 2, 3, 4}
fmt.Println(numbers2)
```

**Output:**
```output{ lineNos=false }
[1 2 3 4]
[1 2 3 4]
```

We can also use shorthand initialization as shown below. There is another way to initialize Go arrays using `...` syntax in which compiler automatically picks the length of the array.

```go
package main

import "fmt"

func main() {
	names := [3]string{"John", "Jenny", "Zoe"}
	fmt.Println(names)

	colors := [...]string{"orange", "red", "blue", "black"} // compiler picks length automatically
	fmt.Println(colors)
}
```

**Output:**

```outpu{ lineNos=false }
[John Jenny Zoe]
[orange red blue black]
```

## Finding Array Length

If we want to check the length property of an array, we can use the `len` function. For Golang arrays, the length and the capacity are the same.

```go
package main

import "fmt"

func main() {
	numbers := [4]int{1, 2, 3, 4}
	names := [3]string{"John", "Jenny", "Zoe"}
	fmt.Printf("Length of %v array: %d\n", numbers, len(numbers))
	fmt.Printf("Length of %s array: %d\n", names, len(names))
}
```

**Output:**
```output{ lineNos=false }
Length of [1 2 3 4] array: 4
Length of [John Jenny Zoe] array: 3
```

## Array Indexing

 To get individual elements of an array, we can use array indexing. Indices start at 0. So, to access element at index position `1`, we can use `names[1]`.

 **Example:**

 ```go
package main

import "fmt"

func main() {
	numbers := [4]int{1, 2, 3, 4}
	names := [3]string{"John", "Jenny", "Zoe"}
	// Accessing elements at an index
	fmt.Println(names[1])
	fmt.Println(numbers[3])
}
 ```

**Output:**
```output{ lineNos=false }
Jenny
4
```

 We can also change the value at any index position using assignment operators. Index must be between `0` and `len(arr) - 1`. If we go above or below this index, we can compilation error (out of bounds).

 **Example:**

 ```go
 package main

import "fmt"

func main() {
	names := [3]string{"John", "Jenny", "Zoe"}
	fmt.Println(names)
	names[1] = "Jane"
	fmt.Println(names)
}
 ```

 **Output:**

```output{ lineNos=false }
[John Jenny Zoe]
[John Jane Zoe] 
```

## How to Iterate through Golang Array

For iterating through an array, we can use for loop. The iterating variable is usually called `i` for index. 

```go
package main

import "fmt"

func main() {
	numbers := [4]int{1, 2, 3, 4}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
}
```

**Output:**
```output{ lineNos=false }
1
2
3
4
```

There is another method to loop through an array without using index variable in which we don't have to worry about off by 1 error in normal for loop. This is using `range` keyword. With this syntax, we can also access the `index` position of each element.

```go
package main

import "fmt"

func main() {
	numbers := [4]int{1, 2, 3, 4}s
	for index, elem := range numbers {
		fmt.Printf("Element at %d: %d\n", index, elem)
	}
}
```

**Output:**

```output{ lineNos=false }
Element at 0: 1
Element at 1: 2
Element at 2: 3
Element at 3: 4
```

If we don't need index of elements, we can replace that with _ and it will be ignored.

```go
for _, elem := range numbers {
    fmt.Println(elem)
}
```

## Multi-Dimensional Arrays

Multi-dimensional array is like an array of array with multiple dimensions like in Mathematics. To access individual elements, we have to use multiple indices.

Multi-dimensional arrays are declared using below syntax.

```go
arr := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
```

We can think of this like a mathematical matrix where we have rows and columns. To access individual elements, we can refer to those elements using double index.

```go
package main

import "fmt"

func main() {
    arr := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(arr[2]) // This returns one-dimensional array located at index 2
	fmt.Println(arr[2][0])
	fmt.Println(arr[2][1])
}
```

**Output:**

```output{ lineNos=false }
[5 6]
5
6
```

### How to iterate through Multi-Dimensional Array

To iterate through multi-dimensional array, we will need as many loop variables as the depth of the multi-dimensional array. In below example, we have depth of 2, so we will need to looping variables `i` and `j`.

```go
package main

import "fmt"

func main() {
    arr := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[i][j])
		}
	}
}
```

**Output:**
```output{ lineNos=false }
1
2
3
4
5
6
```

## Conclusion

In this lesson, we looked at how to initialize a single dimensional and multi-dimensional array. We saw how we can access individual elements at a given index position using *Array Indexing*. Further in this tutorial, we looked at how to iterate through arrays.