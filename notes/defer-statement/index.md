# Defer Statement

Go has normal control flow statements like conditionals, loops, but it also has an interesting idea of `defer`. A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns. Let's see a simple example.

**Example:**

```go
package main

import "fmt"

func main() {
	fmt.Println("hello")
	defer fmt.Println("how are you?") // defer statement runs after the next function call
	fmt.Println("another hello")
}
```

In this case, the function `Println` on line `7` is deferred until the function call on line `8` finishes. So, this gets printed at the end unlike usual control flow.

**Output:**
```output{ lineNos=false }
hello
another hello
how are you?
```

Let's check another example.

```go
package main

import "fmt"

func main() {
	a2 := 1
	fmt.Println(a2)
	fmt.Println("another")
	defer fmt.Println("Hey there!")
	fmt.Println("Another defer")
	fmt.Println("Final")
}
```

**Output:**

```output{ lineNos=false }
1
another
Another defer
Final
Hey there!
```

As you can see, `defer` statement makes the function execution at the end. In general, deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out (LIFO) order. Let's see an example with multiple `defer` statements.

```go
package main

import "fmt"

func main() {
	fmt.Println("staring count")
	defer fmt.Println("Go!")
	for i := 1; i < 4; i++ {
		defer fmt.Println(i)
	}
}
```

In this case, the stack will have `Go!` printing statement first and then numbers from 1 to 3. So, when executing, it executes from the end.

**Output:**

```output{ lineNos=false }
staring count
3
2
1
Go!
```

The `defer` statement can be very useful in file handling or database connection handling where we want to ensure we close the connection or file handle if there is an error somewhere in code.