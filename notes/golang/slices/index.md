# Slices

In Go, you don't have a datastructure like list or dynamic array which will grow as we need to add more elements into it. However, you have **slices** which are made from arrays and can grow like `ArrayList` in other languages.

## Overview

Slice is a continuous segment of an array. It provides numbered sequence of elements from original array. Unlike arrays, these can be of variable typed and are lot more flexible. Elements can be added or removed easily from a slice. 
Slices contain three major components. 

1. *Pointer* to the first element of the slice. 
2. The *length* of the slice. Length of the slice are the number of elements the slice contains. 
3. The *capacity* of slice is calculated from the first element of the slice to the last element of the underlying array. So, capacity will always be greater than or equal to length of the slice. 


## Creating Slice

Declaration of a slice is similar to array except that it does not need the size of the slice. 

```go
<slice_variable_name> := []<data_type>{.... value .....}
```

**Example:**

```go
package main

import "fmt"

func main() {
    numbers := []int {1, 2, 3, 4}
	fmt.Println(numbers)
}
```

**Output:**
```output{ lineNos=false }
1 2 3 4]
```

With such initialization, first an array is created and then the slice reference is returned.
Another method to create a slice is to use `make` function. The syntax for this looks like below.

```go
slice := make ([] <data_type>, length, capacity)
```

The first argument to the [make function](https://pkg.go.dev/builtin#make) is the data type of the slice you want to create then you provide the length of the slice and the capacity of the slice.

**Example:**
```go
package main

import "fmt"

func main() {
    slice := make([] int, 3, 5)
    fmt.Println(slice)
}
```

In this case, uninitialized slice is created using `make` function and all the elements of the slice will get zero value.

**Output:**
```output{ lineNos=false }
[0 0 0]
```

You can check the length and capacity of the slice using `len` and `cap` function respectively.

```go
package main

import "fmt"

func main() {
    slice := make([] int, 3, 5)
    fmt.Printf("Length %d, Capacity: %d\n", len(slice), cap(slice))
}
```

**Output:**
```output{ lineNos=false }
Length 3, Capacity: 5
```

### How to create a Slice from an Array

you can also create a slice from an array. To do that, you can use `array[start_index: end_index]` syntax. In this operation, the `start_index` element is included in the output slice but `end_index` element is excluded.

**Example:**
```go
package main

import "fmt"

func main() {
    numbers := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	slice1 := numbers[0:5]
	fmt.Println(slice1)
	fmt.Printf("Length of array: %d, Capacity of Array: %d\n", len(numbers), cap(numbers))
	fmt.Printf("Length of slice1: %d, Capacity of slice1: %d\n", len(slice1), cap(slice1))
}
```

**Output:**

```output{ lineNos=false }
[0 1 2 3 4]
Length of array: 8, Capacity of Array: 8
Length of slice1: 5, Capacity of slice1: 8
```

As mentioned earlier, the slice capacity is calculated from the start of the slice to end of the underlying array. So, if we modify slice to start from index position `1`. then slice capacity changes to `7` as shown in below code snippet.

```go
func main() {
    numbers := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	slice1 := numbers[1:5]
	fmt.Println(slice1)
	fmt.Printf("Length of array: %d, Capacity of Array: %d\n", len(numbers), cap(numbers)) // len = capacity for array
	fmt.Printf("Length of slice1: %d, Capacity of slice1: %d\n", len(slice1), cap(slice1))
}
```

**Output:**
```output{ lineNos=false }
[1 2 3 4]
Length of array: 8, Capacity of Array: 8
Length of slice1: 4, Capacity of slice1: 7
```

If you do not specify the `start_index`, slice operation default to position starting at index `0`. Similarly, if you do not specify `end_index` position, it defaults to end of underlying array.

**Example:**
```go
package main

import "fmt"

func main() {
    numbers := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	slice2 := numbers[:4] // Without start index, it's 0 by default
	fmt.Println(slice2)
	slice3 := numbers[6:] // end index is length of the array by default, even though we slice from index 6, slice index starts from 0
	fmt.Println(slice3)
}
```

**Output:**

```output{ lineNos=false }
[0 1 2 3]
[6 7]
```

## Slice Indexing

Slice indexing is a way to access individual elements at a given index position. We can access elements at any index position using slice indexing which is very similar to array indexing. We can also modify an element at a given index position for slice.

**Example:**
```go
package main

import "fmt"

func main() {
    numbers := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	slice := numbers[6:]
	fmt.Println(slice)
	fmt.Println(slice[0])
	slice[0] = 100
	fmt.Println(slice)
}
```

As you can see on line `10`, we modified element at index position `0`. The output slice has `100` at index position `0` after this operation

**Output:**

```output{ lineNos=false }
[6 7]
6
[100 7]
```

Slice is basically using the same data as the original array but provides a flexible view on the top of it. It has a different pointer at the beginning of the slice. If we modify the data of a slice, the data in the original array also gets updated.

**Example:**

```go
package main

import "fmt"

func main() {
    numbers := [5]int{1, 2, 3, 4, 5}
	slice := numbers[1:]
	fmt.Println(slice)
	slice[2] = 2
	fmt.Println(slice)
	fmt.Println(numbers)
}
```

In this case, we started with elements 1 to 5 for `numbers` array. When we modify the slice on line `9`, this operation modifies the slice value at index position `2`. However, this also modifies original array value at index position `1`.

**Output:**

```output{ lineNos=false }
[2 3 4 5]
[2 3 2 5]
[1 2 3 2 5]
```

We can append new elements to a slice using built-in function [append](https://pkg.go.dev/builtin#append). From the official documentation, the first argument is the slice to which we want to append new data. The following arguments after these are the list of values we want to append into a slice. The datatype of these values should match the original slice data type.

**Example:**

```go
package main

import "fmt"

func main() {
    array := [8]int{1, 2, 3, 4, 5}
	slice := array[:3]
	fmt.Printf("%v has length %d and capcity %d\n", slice, len(slice), cap(slice))
	slice = append(slice, 40, 50, 60)
	fmt.Printf("%v has length %d and capcity %d\n", slice, len(slice), cap(slice))
}
```

**Output:**
```output{ lineNos=false }
[1 2 3] has length 3 and capcity 8
[1 2 3 40 50 60] has length 6 and capcity 8
```

In this case, because the original array on line `6` was created with length `8`, it has capacity to store three additional elements on line `9`.
If the underlying array has enough capacity to hold new elements, these elements get added and slice pointer is returned. If it does not have enough capacity to append these new elements, a new array will be allocated in the memory and all values are copied into this array. Finally this new array pointer is returned. This new slice will have twice the capacity of the original slice.

```go
package main

import "fmt"

func main() {
    array := [5]int{1, 2, 3, 4, 5}
	slice := array[:3]
	fmt.Printf("%v has length %d and capcity %d\n", slice, len(slice), cap(slice))
	slice = append(slice, 40, 50, 60)
	fmt.Printf("%v has length %d and capcity %d\n", slice, len(slice), cap(slice))
}
```

In this case, we had original array of length `5` on line `6`. So, there is no room to add additional three elements. So, Go compiler will copy original array elements to a new array and capacity will be double after that copy operation.

**Output:**

```output{ lineNos=false }
[1 2 3] has length 3 and capcity 5
[1 2 3 40 50 60] has length 6 and capcity 10
```

## Merge Two Slices

We can append a new slice into existing slice using `append` function. In this case, the second slice is expanded using `...` operator.

```go
append(slice1, slice2...)
```

**Example:**

```go
package main

import "fmt"

func main() {
    numbers := [4]int{1, 2, 3, 4}
	numbers2 := [4]int{5, 6, 7, 8}
	slice1 := numbers[:]
	slice2 := numbers2[:2]
	fmt.Printf("Slice1: %v, Slice2: %v\n", slice1, slice2)
	output_slice := append(slice1, slice2...)
	fmt.Println(output_slice)
}
```

Here, we are appending `slice2` using `...` operator on line 11. This `append` function is variadic which means it takes variable number of arguments.

**Output:**

```output{ lineNos=false }
Slice1: [1 2 3 4], Slice2: [5 6]
[1 2 3 4 5 6]
```

## Deleting Elements from a Slice

We can delete an element from a slice by splitting the slice at that index position and then appending two slices.

**Example:**

```go
package main

import "fmt"

func main() {
    numbers := [6]int{1, 2, 3, 4, 5, 6}
	slice := numbers[:]
	fmt.Println(slice)
	output_slice1 := slice[:2]
	output_slice2 := slice[3:]
	deleted_element_slice := append(output_slice1, output_slice2...)
	fmt.Println(deleted_element_slice)
}
```

In above code, we start with a slice same as an array that is with elements 1 to 6. Now, if we want to delete an element at index position 2, we have to create two slices around that index and append those two together to create a new slice. This way we will end up deleting the element at index position 2.

**Output:**

```output{ linenNos=false }
[1 2 3 4 5 6]
[1 2 4 5 6]
```

- In Go 1.21 and above, you can use `clear` function to delete all elements from a slice.

```go
numbers := []int{1, 2, 3, 4, 5, 6}
clear(numbers)
```

## Copying Slice

There is also [copy function](https://pkg.go.dev/builtin#copy) which copies all elements of the source slice into destination slice. This function returns the number of elements copied. Here, both slices have to be of the same type otherwise Go compiler throws an error.

**Example:**

```go
package main

import "fmt"

func main() {
    var src = []int{1, 2, 3, 4, 5}
	var dest = make([]int, 4)
	count := copy(dest, src)
	fmt.Printf("Destination slice: %v, number of elements copied: %d\n", dest, count)
}
```

**Output:**

```output{ lineNos=false }
Destination slice: [1 2 3 4], number of elements copied: 4
```

## Looping through Slices

To iterate through slice, we can use the same syntax as [arrays](../golang-arrays/).

```go
package main

import "fmt"

func main() {
	var src = []int{1, 2, 3, 4, 5}
	var slice = src[:4]
    for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
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

Again, we can use the `range` to get index position of each element of a slice and if we don't care about the index position, we can ignore them using `_` syntax as shown in below code.

```go
func main() {
	var src = []int{1, 2, 3, 4, 5}
	var slice = src[:4]
	for _, elem := range slice {
		fmt.Println(elem)
	}
}
```

## Conclusion

In this lesson, we learned about slices, one of the most powerful data structure in Go. We learned how to create a slice from an array and how we can index elements or even modify slice elements using indexing. Later, we learned how to add new elements into a slice and copy a slice from another slice. Slice is basically built on top of an array so modifying elements of a slice will have impact on underlying array. Finally, we saw how to iterate through each element of a slice.