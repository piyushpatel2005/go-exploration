# Select Statement

The `select` statement in Go allows multiple goroutine to wait on multiple communication operations. It is used to execute a channel among many options.

## Overview

The `select` statement is similar go `switch` statement but this is specifically for channels. The `select` statement makes goroutine wait on multiple communication operations. In this, each case statement waits for a send or receive operation from a channel. The select statement blocks until any of the case statements are ready. If there is a situation when there are multiple case statements ready, then it selects one randomly and proceeds. 

## Syntax

The syntax for `select` statement looks like this.

```go
select {
    case channel_send_receive:
        // statements
    case channel_send_receive:
        // statements
    default:
        // optional block    
}
```

## Examples

The `select` statement with channels and goroutine become a powerful tool for managing synchronization and concurrency. In `select` statement, the `default` case will be executed if none of the `case` statements have send or receive operation ready. This helps make the `select` statement non-blocking. If all `case` statements are blocked, the `default` statement will be executed. 

```go
package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendOne(ch1)
	go sendTwo(ch2)

	select {
	case value1 := <-ch1:
		fmt.Println(value1)
	case value2 := <-ch2:
		fmt.Println(value2)
	default:
		fmt.Println("Default")
	}
}

func sendOne(ch chan string) {
	ch <- "One"
}

func sendTwo(ch chan string) {
	ch <- "Two"
}
```

In this case, because the goroutine takes some time to execute, we almost always get `Default` as the output. In order to allow the goroutine to execute, we have to provide some delay before the `select` statement.

```go
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendOne(ch1)
	go sendTwo(ch2)

	time.Sleep(1 * time.Second)
	select {
	case value1 := <-ch1:
		fmt.Println(value1)
	case value2 := <-ch2:
		fmt.Println(value2)
	default:
		fmt.Println("Default")
	}
}
```

In above code snippet, I have added 1 second sleep time before the select statement. This gives enough time for both goroutine `sendOne` and `sendTwo` to execute. Because these two goroutine execute independent of each other, the output of the program is non-deterministic. Sometimes it will output `One` and sometimes it will be `Two`.

Another important consideration is that we can use `break` statement in the `case` blocks to terminate the `select` statement.

The `switch` statement is non-blocking whereas the `select` statements can block as they are using channels. The `switch` statement is deterministic because it goes in sequence whereas the `select` statement is not deterministic.