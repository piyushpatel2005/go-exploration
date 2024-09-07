# Goroutines

Go routines is the mechanism used in Go to create concurrent programs. A goroutine is a simple function or method which runs concurrently with other functions or methods. This is considered as a lightweight thread that has its own independent execution. 

## Overview

Go routines can be executed concurrently with other go-routines and it is the fundamental way to execute concurrent programs in Go. Go routines are managed by Go runtime. The syntax looks like below which applies to functions or methods. `go` keyword makes it run in a separate go routine.

```go
go func_name()
```

Other programming languages use the concept of threads to achieve concurrency. In the Go programming language, instead of using threads, you utilize lightweight and more compact functions or procedures to execute concurrently. When compared to operating numerous threads concurrently, the resource cost of constructing goroutines is extremely low. Go routines only consume a small amount of memory. Threads, on the other hand, often require a fixed amount of memory which often around 100 times more than goroutines.

### Comparison of Threads and Go routine

- go routines are cheaper. Just few KBs in stack and stack can grow and shrink in size in Go. With threads, stack size has to be specified upfrtont and is fixed in size. OS threads generally start with 1MB size. Go routines are cheap so we can start hundreds of thousands of go-routines.
- Go routines are multiplexed to a fewer number of OS threads. Even a single OS thread can handle 1000s of go routines.
- Scheduling of go routines is managed by Go runtime. So, it's faster. For threads, scheduling is done by OS runtime. So, context switching time for go-routines is much faster.
- Go routines communicate using channels. Channels do not cause race condition when using shared memory. This is a powerful construct built into the language. It can be thought of bytes using with Go routines communicate with each other.

## Simple Goroutine

A simple goroutine looks like below code.

```go
package main

import "fmt"

func test_routine() {
	fmt.Println("Hello from Goroutine")
}

func main() {
	go test_routine()
	fmt.Println("Hello from main")
}
```

In this code, we have a simple function `test_routine()` which we call from `main()` using `test_routine()`. Notice that we have added `go test_routine()` which makes it a goroutine.

```output{ lineNos=false .show-prompt }
go run main.go 
Hello from main
```

When you run this code, you may or may not see output from the `test_routine()` function. This is because goroutine runs in separate parallel process. The `main()` function does not wait until the routine has finished. It executes the next statement of printing from main and then terminates the `main()` function ending the program. There is not enough time for `test_routine` to execute before the `main()` function ends. 

To see output from `test_routine()` we can provide manual pause to `main()` function control flow. We can use `time.Sleep()` to sleep for a second or more if we want.

```go
package main

import (
	"fmt"
	"time"
)

func test_routine() {
	fmt.Println("Hello from Goroutine")
}

func main() {
	go test_routine()
	fmt.Println("Hello from main")
	time.Sleep(1 * time.Second)
}
```

With this, we always see the output from `Println` statement in `test_routine()`.

```output{ .show-prompt lineNos=false }
go run main.go 
Hello from main
Hello from Goroutine
```

This is just a hack that I've applied for now. Later on, I will show better options to control flow using WaitGroups and channels.

To actually see parallel execution of Goroutines, we can use the previous sequential program write that using goroutines.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i <= 100; i++ {
		go printSequence(i) // run as go routine
	}
	elapsed := time.Since(start)
	fmt.Println("Elapsed time", elapsed)
	time.Sleep(2 * time.Second)
}

func printSequence(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Sequence", i)
}
```

In this code, we are using `printSequence` function as a goroutine. Previously, this code was taking around 100 seconds. With this little change, if we execute it, this finishes in 2 seconds. Also those two seconds are due to the `time.Sleep` call we have in the `main()` function.

## Goroutine Scheduler

When we launch go program, it will launch OS threads that is proportional to logical CPUs available. They are completely managed by kernel at OS level. From creating, managing and terminating threads is managed by OS. We can find available number of logical processors using runtime package with `Numcpus()` methods. The logical threads are the number of physical cores available in your system multiplied by the number of threads that can run on each core (hardware threads). Go routines are lightweight application threads that run independent of one another. The go routine has scheduler that will multiplex the go-routines on OS level threads in the go runtime. It schedules arbitrary number of go-routines onto an arbitrary number of OS threads(m x n multiplexing).

**OS Scheduler** manages OS threads for each logical core. Within Go runtime, each of these threads will have one queue associated with it called LRQ (Local Run quque). It consists of all go-routines to be executed in that thread. Go runtime scheduler will perform scheduling and context switching belonging to LRQ. We also have GRQ (Global run queue). It consists of all go-routines which are not assigned any LRQ. Go scheduler will assign these go routines to any of the Local run queue of OS thread. This is how Go scheduler works and multiplexes go routines on OS threads. Properties of Go scheduler include:

In Go, we have **Cooperative scheduler** which means it's not pre-emptive. There is no time based pre-emption from the OS. In this, OS never interrupts a running process to initiate a context switch from one process to another. Processes must voluntarily yield control periodically or when logically blocked on a resource. Of course, there are some specific checkpoints which will ensure go-routine can yield execution to other go-routines. These are called context switches. For example, situations where runtime calls the scheduler on function calls to check if new go routine needs to be scheduled. In this case, a context switch might happen. It is also possible that current go routine may continue. Examples when these kinds of context switching can happen include: function calls, network calls, garbage collection, channel operations, go keyword usage. The scheduler gets an option to context switch, but it does not mean it will always do so.


### Goroutines run in Parallel

The goroutines are multiplexed in a small number of OS threads. That means single thread can handle thousands of goroutines. When goroutine is blocking the thread, Go will automatically spin up another thread and move the other goroutines to the new one. By default, `main` function in the main package is a goroutine itself called the **main go-routine**. All go-routines are started from main goroutine. it represents main program. When main go-routine exits, it assumes all go-routines have finished executing. 

Go routines do not have parents or children. They all execute in parallel. All go routines exit when main go routine exits. To prove that there is no parent-child relationship, we can use below code. We would expect, the output from `process1()` to be printed first before output from `process2()`. However, when we execute this code, the output of this code will not be deterministic because both methods `process1` and `process2` execute in separate go routines.

```go
package main

import "fmt"
import "time"

func process1() {
	time.Sleep(1 * time.Second)
	fmt.Println("Running Process 1")
}

func process2() {
	go process1()
	time.Sleep(1 * time.Second)
	fmt.Println("Running Process 2")
}

func main() {
	go process2()
	time.Sleep(2 * time.Second)
}
```

## Anonymous goroutine

Anonymous goroutines are anonymous functions. They don't have identifier associated with name as a function name. They are used only once. They can be called by prefixing their definition with `go` to make them goroutines. 

```go
go func() { }(args...)
```

These  routines behave just like any other goroutine. Below code demonstrates very simple goroutine.

```go
package main

func main() {
    go func() {
        fmt.Println("An anonymous function")
    }()
    time.Sleep(1 * time.Second)
}
```