# Switch Case Statement

Even though we have `if...else` construct in Go, there may be situation when we have to check so many conditions and even we may want to execute two `if` blocks than just one. `switch-case` statements provide this kind of flexibility and are more readable when we have to write multiple `if` blocks.

### Common Switch Case Statement

Switch..case statements provide multi-way control flow. This also has a `default` block which is executed if none of the `case` conditions match. The syntax for this looks like below.

```go
switch expression {
    case value1:
        // execute block
    case value2:
        // block 2
    default:
        // default block if no match found
}
```

This is a typical `switch...case` block syntax. Let's see an example of this.

```go
package main

import "fmt"

func main() {
	// normal switch case
	var month string = "March"
	switch month {
	case "January":
		fmt.Println("This is January")
	case "February", "March": // check for multiple values
		fmt.Println("This is feb or March")
	default:
		fmt.Println("This is neither first or second month of the year.")
	}
}
```

In this code, we are checking for the `month` variable in `switch` statement. As you can see on line `11`, we can also compare multiple values in a single `case` statement.

**Output:**
```output{ lineNos=false }
This is feb or March
```

### Switch Case with Fallthrough

Unlike other programming languages, in Go, the `switch` statement does not fallthrough by default. That is we do not have to write `break` statement after each `case` block. If we want this behavior to change, then we have to explicitly write `fallthrough` keyword at the end of the `case` block.

Interestingly, we can specify this `fallthrough` keyword only in some of the `case` block and have pretty complex control flow with `switch..case`.

```go
package main

import "fmt"

func main() {
	// use of fallthrough will move control to next case block
	var month = "February"
	switch month {
	case "January":
		fmt.Println("This is January")
	case "February", "March": // check for multiple values
		fmt.Println("This is feb or March")
		fallthrough
	default:
		fmt.Println("This is default block.")
	}
}
```

In above code block, we have used `fallthrough` keyword in the second `case` block. This means when this condition matches (i.e. when `month` variable is either `February` or `March`), it will fall through and check for other conditions which matches. In this case, no other conditions after second `case` block matches, so it will execute the `default` block statements.

**Output:**
```output{ lineNos=false }
This is feb or March
This is default block.
```

### Switch Case as Conditional

We can also omit the expression in the `switch` block and actually use the variable for verifying a condition in the `case` block. This will behave like an `if..else` structure.

```go
switch {
    case condition1:
        // statements
    case condition2:
        // statements 2
    default:
        // default statements
}
```

In below example, we are checking multiple conditions with a variable `num`. If the number is greater than `100` or `50` or `10`, it will print corresponding message from the `case` block. One thing to note here is that this variable is also accessible inside the `case` or `default` statement, as we can see on line `16`.

```go
package main

import "fmt"

func main() {
	// using switch case as if..else
	var num int = 9
	switch {
	case num > 100:
		fmt.Println("The number is greater than 100")
	case num > 50:
		fmt.Println("The number is greater than 50")
	case num > 10:
		fmt.Println("The number is greater than 10")
	default:
		fmt.Println("The number is", num)
	}
}
```

**Output:**
```output{ lineNos=false }
The number is 9
```

## Conclusion

In this lesson, we looked at various ways in which `switch...case` statement may be used. This can be used as a conditional statement or it can also `fallthrough` and execute multiple matching statements.