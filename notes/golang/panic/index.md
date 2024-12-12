# Panic

In Go, panic is like an exception arising at runtime. Due to this execution of program is terminated. We can use `panic` to signal abnormal behavior in the program. 

While working with channels, we may notice panic, for example, when sending to a channel once it has been closed or closing an already closed channel. Panic is used to signal the program to fail when there are errors that we cannot handle.


```go
package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Enter two numbers to divide: ")
	fmt.Scanf("%d %d", &a, &b)
	if (b == 0) {
		panic("Cannot divide by zero")
	} else {
		fmt.Println(a / b)
	}
}
```

In above code, we are taking two integer values from the user using `Scanf` function and assigned them to `a` and `b` variables. Next, we want to divide `a` by `b`. Now, if the denominator is zero, we cannot divide `a`. In this situation, we have to exit the program, but we want to provide proper information. For this, we can use `panic` function to throw an error.

The program by default does not necessarily stop running when it encounters an error. We can use `panic` function to inform the program to stop allowing program to exit gracefully.

```output { .show-prompt lineNos=false }
go run main.go
Enter two numbers to divide: 
5 0
panic: Cannot divide by zero
```