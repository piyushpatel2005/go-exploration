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
	// go receive(ch, wg) // If we call receive() before sending 4th element, it will avoid deadlock
	// ch <- 4            // if we call one more time, it creates deadlock situation because we exceed the capacity before calling receive
	go receive(ch, wg) // If we call receive before sending any data into channel, again this causes deadlock when channel is empty
	fmt.Println("Sent 3 elements to the channel")
	wg.Done()
}

func receive(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for incoming data")
	fmt.Println("Received:", <-ch)
	wg.Done()
}


// package main

// import (
//     "fmt"
//     "sync"
// )

// func main() {
//     var wg sync.WaitGroup
//     wg.Add(2)
//     ch := make(chan int, 10)
//     go produce(ch, &wg)
//     wg.Wait()
// }

// func produce(ch chan int, wg *sync.WaitGroup) {
//     for i := 10; i <= 100; i += 10 {
//         ch <- i
//     }
//     fmt.Println("Exiting Produce")
//     close(ch)
//     go consume(ch, wg)
//     wg.Done()
// }

// func consume(ch chan int, wg *sync.WaitGroup) {
//     for val := range ch {
//         fmt.Println("Received: ", val)
//     }
//     fmt.Println("Exiting Consume")
//     wg.Done()
// }