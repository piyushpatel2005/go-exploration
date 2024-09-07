# Channels

Channels are a way to send and receive data between Go routines. It acts as pipes that connects different goroutines with each other for data exchange.

<!--more-->

## What are Go channels?

Channels are a way through which different go-routines communicate. It's like a data exchange mechanism in Go which allows us to exchange data between different parts of go-routine. Channels make concurrent programming easier. Traditional threading models were communicating by sharing memory, however, channels try to avoid that because that can result in nasty bugs. In traditional concurrent programming, these shared memory is locked by threads. There is possiblity of deadlock and thread contention over that data. 

Go routines and channels allow access to data using channels. Channels share memory by communicating which ensures that at a given time, only one routine has access to the memory. The communication in the channels is bi-directional, that is you can send and receive values using the same channel. When you send data to a channel, the program execution control is blocked until another goroutine reads from this channel. This way go-routines synchronize without the use of locks or conditional variables. The same thing happens when reading from a channel. The reading is blocked until another goroutine writes to the channel. 

Each channels holds data of particular data type. Syntax for this looks like below. We use `chan` keyword to declare channel of specific type. For example, below is example of string data type channel.

```go
var c chan string
c := make(chan string)
```

Let's look at some common operations with channels.

1. You can send a value to a channel using `<-` operator. For this to work, the data type of input must match the channel type.
2. You can receive a value from a channel using `<-ch`
3. Close a channel using `close(ch)` function.

```go
package main

import (
	"fmt"
)

func main() {
	var ch chan string
	ch = make(chan string)
	go hello(ch)

	// receive message from the channel
	msg := <-ch
	fmt.Println(msg)
    fmt.Println("End of Program")
}

func hello(ch chan string) {
	ch <- "Hello"
}
```

This code declares a channel `ch` of type `chan string`. The next line creates this variable using `make()` function. The goroutine `hello` is passed reference to the channel and inside this routine, it sends a message `"Hello"` to the channel. In the `main()` function, we receive this data from the channel in `msg` variable which is printed when we run this code.

The data sent in channel can be received only once in any of the goroutines. By default, when we created channel, it creates unbuffered channel. Earlier, we need `time.Sleep()` or `WaitGroup` in order to make sure that we allow other goroutines to finish before `main` goroutine. However, in the output, notice that the `main` goroutine waits until you've read from the channel. This means it's going to block the execution of goroutine sending the data until the other go routine has received previous value. If you add a delay in `hello` goroutine, this can be easily seen that `main` goroutine is waiting for `hello` routine to finish.

```go
func hello(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Hello"
}
```

This feature allows us to write concurrent programs. For example, below code crunches numbers in parallel where it can where it's finding sum and max value from a slice.

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Main routine")
	numbers := []int{1, 2, 12, 4, 5, 6, 7, 8}
	start := time.Now()


	ch1 := make(chan int)
	go sum(numbers, ch1)
	sum := <-ch1
	fmt.Println("Sum:", sum)

	ch2 := make(chan int)
	go max(numbers, ch2)
	max := <-ch2
	fmt.Println("Max:", max)

	fmt.Println("Time taken:", time.Since(start))
}

func sum(numbers []int, ch chan int) {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	time.Sleep(2 * time.Second)
	ch <- sum
}

func max(numbers []int, ch chan int) {
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	time.Sleep(2 * time.Second)
	ch <- max
}
```

Here, channel `ch1` is created of type `chan int` and passed to `sum` function where it takes 2 seconds. Notice that the execution is blocking, so it takes 2 seconds. The program waits for message to be written to channel. Similarly new channel `ch2` is created and used for calculating `max` from an array which also takes 2 seconds. In total, this code execution takes about 4 seconds.

```{.show-prompt lineNos=false }
go run main.go
Main routine
Sum: 45
Max: 12
Time taken: 4.002665167s
```

The code actually waits where we are reading from the channel, so if I move my read statements together, I might be able to get some parallelism like below snippet for `main` function.


```go
func main() {
	fmt.Println("Main routine")
	numbers := []int{1, 2, 12, 4, 5, 6, 7, 8}
	start := time.Now()


	ch1 := make(chan int)
	go sum(numbers, ch1)

	ch2 := make(chan int)
	go max(numbers, ch2)
	sum := <-ch1
	fmt.Println("Sum:", sum)
	max := <-ch2
	fmt.Println("Max:", max)

	fmt.Println("Time taken:", time.Since(start))
}
```

This code execution takes only 2 seconds.

## Unbuffered Channel

When we created channels above, they were by default **unbuffered channels**. This means it will block the execution of go routine sending the data until the other go routine has received the previously sent data. The receiver on a channel is blocked until there is some data coming in from another go routine into that channel. This type of channels are called Unbuffered channels. Tehy need a receiver as soon as a message is emitted to the channel. They do not have buffer which means they do not have capacity to hold on to data so we don't have to specify `capacity` of the channel when creating a channel using `make` function. They do not store data, so they do not have any length.

For unbuffered channels, the length and the capacity is always zero.

## Buffered Channels

On the other hand, if we want channel to hold some data, we can define a **buffered channel**. These types of channel do not block the execution until the buffer is full and at that point sending data into this channel is blocked until data is received by another goroutine. Similarly, receiver is blocked if the channel is totally empty. To create a buffered channel, we can define the `capacity` in the `make` function using the syntax below.

```go
ch := make(chan <data_type>, capacity)
ch := make(chan int, 10)
```

Again, to find the capacity of the channel, we can use `cap()` function and to find currently available data in the channel, we can use `len()` function. The `len()` function will represent number of elements queued in this channel. Based on this, the length of the channel can never be higher than the capacity of the channel.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int, 3)
	var wg sync.WaitGroup
	wg.Add(2)
	go send(ch, &wg)
	wg.Wait()
}

func send(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Inside send go-routine")
	ch <- 1
	ch <- 2
	ch <- 3
	go receive(ch, wg)
	fmt.Println("Sent 3 elements to the channel")
	wg.Done()
}

func receive(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for incoming data")
	fmt.Println("Received:", <- ch)
	wg.Done()
}
```

In this code, we are creating a buffered channel of capacity three. We are adding 2 to wait group and sending 3 elements to the channel. With this, we can receive the incoming data from the channel using `receive` function.

```output{ .show-prompt lineNos=false }
go run main.go
Inside send go-routine
Sent 3 elements to the channel
Waiting for incoming data
Received: 1
```

Again, if we call `receive` before sending any data into the channel, we run into deadlock error because channel is empty and we are trying to read from the channel. Below snippet shows demonstration where it throws deadlock error.

```go
func send(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Inside send go-routine")
	go receive(ch, wg)
	fmt.Println("Sent 3 elements to the channel")
	wg.Done()
}
```

Also, because the channel has size of three, we cannot send forth element until we have read previous elements. So, below code also throws deadlock error.

```go
func send(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Inside send go-routine")
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	go receive(ch, wg)
	fmt.Println("Sent 3 elements to the channel")
	wg.Done()
}
```

To avoid this error, we have to make sure that we do not write more than the capacity of the channel. If we want to write more data, we first have to read from the channel like below snippet.

```go
func send(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Inside send go-routine")
	ch <- 1
	ch <- 2
	ch <- 3
	go receive(ch, wg)
	ch <- 4
	go receive(ch, wg)
	fmt.Println("Sent 3 elements to the channel")
	wg.Done()
}
```

Below code shows produce consume kind of pattern with the help of goroutines and channels.

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    ch := make(chan int, 10)
    go produce(ch, &wg)
    wg.Wait()
}

func produce(ch chan int, wg *sync.WaitGroup) {
    for i := 10; i <= 100; i += 10 {
        ch <- i
    }
    fmt.Println("Exiting Produce")
    close(ch)
    go consume(ch, wg)
    wg.Done()
}

func consume(ch chan int, wg *sync.WaitGroup) {
    for val := range ch {
        fmt.Println("Received: ", val)
    }
    fmt.Println("Exiting Consume")
    wg.Done()
}
```

While receiving data from the channel, we can check if the channel is closed by assigning second value when receiving from the channel. This will have boolean values. If `ok` has `true`, it means channel is open and if `ok` is `false`, channel is closed and we will not receive any more values.

```go
v, ok := <- ch
```

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2

	data, ok := <-ch
	fmt.Println(data, ok)
	close(ch)
	data, ok = <-ch
	fmt.Println(data, ok)
	data, ok = <-ch
	fmt.Println(data, ok)
}
```

```output{ .show-prompt lineNos=false }
go run main.go
1 true
2 true
0 false
```

## Closing a Channel

To close a channel, we can use `close()` function. We need to close the channel to make sure no more data can be sent to the channel. This is done when we do not want to send any more data to the channel. 

```go
package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	close(ch)

	data, ok := <-ch
	fmt.Println(data, ok)
	data, ok = <-ch
	fmt.Println(data, ok)
	data, ok = <-ch
	fmt.Println(data, ok)
}
```

In this example, we write data into the channel. After that, we close the channel using `close` function. We can still read the data using channel read operation.

## Iterating Channel Data

To iterate through channel data we can use `for ... range` structure. In this case, we have to ensure that we close the channel.

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 3)
	wg.Add(1)
	go assign(ch, &wg)
	wg.Add(1)
	go work(ch, &wg)
	wg.Wait()
}

func assign(ch chan int, wg *sync.WaitGroup) {
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println("Assigned work to worker")
	close(ch)
	wg.Done()
}

func work(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for some work")
	for value := range ch {
		fmt.Println("Received:", value)
	}
	wg.Done()
}
```