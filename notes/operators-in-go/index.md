---
title: "Operators in Go"
date: 2023-07-29T19:54:28-04:00
draft: false
---

Golang provides different operators to perform different kinds of operations. The main operators are arithmetic, comparison,logical, assignment and bitwise operators.

<!--more-->

## Overview

**Operators** specify the operation we want to perform on elements. These elements are also called **operand**s on which the operation is being performed. There are five main types of operators.

1. Arithmetic Operators
2. Comparison Operators
3. Logical Operators
4. Assignment Operators
5. Bitwise Operators

## Arithmetic Operators
Arithmetic operators are used to perform arithmetic operations on operands.

| Operator | Description |
|:---------|:------------|
| `+` | Addition |
| `-` | Subtraction |
| `*` | Multiplication of two numbers |
| `/` | Division. This returns float result for the division of two numbers. |
| `%` | This returns the remainder for the division of two numbers. This is also called modulo operator |
| `++` | Unary increment by 1 |
| `--` | Unary decrement by 1 |

```go
package main

import "fmt"

func main() {

	a, b := 2, 5
	fmt.Printf("%d + %d: %d\n", a, b, a+b)

	fmt.Printf("%d - %d: %d\n", b, a, b-a)
	fmt.Printf("%d * %d: %d\n", a, b, a*b)
	fmt.Printf("%d / %d: %d\n", b, a, b/a) 
    // '/' returns quotient of division operation
	fmt.Printf("%d mod %d: %d\n", b, a, b%a) 
    // '%' returns remainder of division operation

	// unary operators
	i := 10
	i++
	fmt.Printf("After i++ : %d\n", i)
	i--
	fmt.Printf("After i-- : %d\n", i)
}
```
**Output**
```output{ lineNos=false }
2 + 5: 7
5 - 2: 3
2 * 5: 10
5 / 2: 2
5 mod 2: 1
After i++ : 11
After i-- : 10
```

## Comparison Operators
These are used to compare two values or variables.

| Operator | Description |
|:---------|:------------|
| `>` | Greater than |
| `<` | Less than |
| `<=` | Less than or equal to |
| `>=` | Greater than or equal to |
| `!=` | Is not Equal to |
| `==` | Equal to |

```go
package main

import "fmt"

func main() {
	a, b := 10, 10
	fmt.Printf("%d == %d: %t\n", a, b, a == b)
	fmt.Printf("%d != %d: %t\n", a, b, a != b)

	a = 20
	fmt.Printf("%d > %d: %t\n", a, b, a > b)
	fmt.Printf("%d < %d: %t\n", a, b, a < b)

	fmt.Printf("%d >= %d: %t\n", a, b, a >= b)
	fmt.Printf("%d <= %d: %t\n", a, b, a <= b)

	a = 10
	fmt.Printf("%d > %d: %t\n", a, b, a > b)
	fmt.Printf("%d < %d: %t\n", a, b, a < b)

	fmt.Printf("%d >= %d: %t\n", a, b, a >= b)
	fmt.Printf("%d <= %d: %t\n", a, b, a <= b)

}
```

**Output**

```output{ lineNos=false }
10 == 10: true
10 != 10: false
20 > 10: true
20 < 10: false
20 >= 10: true
20 <= 10: false
10 > 10: false
10 < 10: false
10 >= 10: true
10 <= 10: true
```

## Logical Operators
Logical operators are usually used with conditions to verify more than one condition or to negate a condition. These has three main operators.

| Operator | Description |
|:---------|:------------|
| `and` | AND operation |
| `or` | OR operation |
| `not` | Complement operation |

```go
package main

import "fmt"

func main() {
	i := 10

	fmt.Printf("(i < 100) && (i > 5): %t\n", (i < 100) && (i > 5))     // true
	fmt.Printf("(i < 100) && (i > 500): %t\n", (i < 100) && (i > 500)) // false

	fmt.Printf("(i > 100) && (i > 500): %t\n", (i < 100) && (i > 5))   // false
	fmt.Printf("(i < 100) && (i > 500): %t\n", (i < 100) && (i > 500)) // true

	// unary operator - reverses logical output
	fmt.Printf("!true: %t\n", !true)
	fmt.Printf("!false: %t\n", !false)
	fmt.Printf("!(i > 100): %t\n", !(i > 100))
}
```

**Output**
```output{lineNos=false}
(i < 100) && (i > 5): true
(i < 100) && (i > 500): false
(i > 100) && (i > 500): true
(i < 100) && (i > 500): false
!true: false
!false: true
!(i > 100): true
```

## Assignment Operators

Assignment operators are used to assign values to a variable. These include very basic `=` assignment but include some more.

| Operator | Description |
|:---------|:------------|
| `=` | Assign a value to a variable |
| `+=` | Add and assigne a value |
| `-=` | Subtract and assgn a value |
| `*=` | Multiply and assign |
| `/=` | Divide and assign a value |
| `//=` | Divide and assign quotient |
| `%=` | Divide and assign remainder |

```go
package main

import "fmt"

func main() {
	var a int
	a = 10

	// short hand operator
	b := 10

	a += 5
	fmt.Println("a += 5:", a)

	b -= 2
	fmt.Println("b -= 2:", b)

	b *= 2
	fmt.Println("b *= 2:", b)

	b /= 2
	fmt.Println("b /= 2:", b)

	a %= 2
	fmt.Println("a %= 2:", a)
}
```

**Output**
```output{ lineNos=false }
a += 5: 15
b -= 2: 8
b *= 2: 16
b /= 2: 8
a %= 2: 1
```

## Bitwise Operators
Bitwise operators perform operations at bit-level. All values in computer system are stored at bit level using 0 and 1.

| Operator | Description |
|:---------|:------------|
| `&` | Bitwise AND operation |
| `|` | Bitwise OR operation |
| `^` | Bitwise XOR operation |
| `~` | Bitwise NOT operation |
| `<<` | Shift Left operation |
| `>>` | Shift Right operation |

```go
package main

import "fmt"

func main() {
	// Perform bit by bit operations
	a := 4                                  // 0100
	b := 2                                  // 0010
	c := 5                                  // 0101
	fmt.Printf("%d & %d = %d\n", a, b, a&b) // 0000 = 0

	fmt.Printf("%d | %d = %d\n", a, b, a|b) // 1010 = 6

	fmt.Printf("%d ^ %d = %d\n", a, c, a^c) // 0001 = 1

	fmt.Printf("%d << 1 = %d\n", b, b<<1) // 0100 = 4
	fmt.Printf("%d << 2 = %d\n", b, b<<2) // 1000 = 8
	fmt.Printf("%d << 1 = %d\n", c, c<<1) // 1010 = 10

	fmt.Printf("%d >> 1 = %d\n", c, c>>1) // 0010 = 2
}
```

**Output**
```output{ lineNos=false }
4 & 2 = 0
4 | 2 = 6
4 ^ 5 = 1
2 << 1 = 4
2 << 2 = 8
5 << 1 = 10
5 >> 1 = 2
```

## Conclusion

In this tutorial, we looked at various kinds of operations in Golang. We have five main operator types in Go. These are arithmetic operators, bitwise operators, comparison operators, logical operators and assignment operators.