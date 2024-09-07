# Waitgroups

Golang has another primitive structure to help with synchronization of go routines. That is using waitgroups.

## What is WaitGroup?

During the usage of go-routines, we saw that main go-routine was terminating before the other go-routines completed or even began their execution. I used `time.Sleep()` in order to make sure all our go routines finish before the `main()` function. However, in practice, that's not ideal because of two reasons. First, we do not know exactly how long all go routines may take. They may take varying amount of time each time they are called. Second, even if we know the time it may take, it's not a good idea to pause the execution. The code might finish in 500ms and we may be falsely pausing execution for a second. This is unncessary delay that may not give a good user experience.

To wait for all go-routines to finish, we can use **waitgroup**. A wait group is a construct for synchronization that allows multiple go-routines to wait for each other. In Go, there is package `sync` which includes `WaitGroup` and it acts like a counter that keeps information about spawn go routines. It blocks execution in a structured manner until its internal counter becomes zero. That is all go-routines has finished. The syntax for creating a waitgroup looks like below.

## Syntax of WaitGroups

```go
var wg sync.Waitgroup
```

We can call various methods on the `WaitGroup`. `Add(int)` method is used to track the number of go-routines to wait for. This acts like a counter. Another method `Wait()` blocks the execution of the code until the counter reduces to zero. The `Done()` method decreases the internal counter by one. Each go routine at the end calls `Done()` method which reduces the counter by 1. 

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(100)
	for i := 1; i <= 100; i++ {
		go printSequence(i, &wg) // run as go routine
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Elapsed time", elapsed)
	time.Sleep(2 * time.Second) // add this timeout so that it doesn't exit without printing results
}

func printSequence(i int, wg *sync.WaitGroup) {
	time.Sleep(1 * time.Second)
	fmt.Println("Sequence", i)
	wg.Done()
}
```

In this code, we are printing 1 to 100 in go routines. You will notice that the execution finishes in around a second. When using `WaitGroup`, we first need to import `sync` package. Next everytime we create a go routine, we add 1 to the wait group. Inside the go routine, we can call `Done()` to indicate that the go routine has finished. In the `main()` code, we indicate to wait for all go routines to finish using `Wait()` method. Also something to note is that we pass `WaitGroup` as a reference to go routines.

The output of above code is not in specific order because go routines are not deterministic. So, it may print different sequence every time you execute this code. I have kept `time.Sleep()` but not that it finishes before 2 seconds. You can also remove those `time.Sleep()` calls to ensure it works like below.

```go
func main() {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go printSequence(i, &wg) // run as go routine
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Elapsed time", elapsed)
}

func printSequence(i int, wg *sync.WaitGroup) {
	fmt.Println("Sequence", i)
	wg.Done()
}
```