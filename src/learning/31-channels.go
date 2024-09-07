package main

import (
	"fmt"
	"sync"
)

// import "time"

func main() {
	var wg sync.WaitGroup
	ch := make(chan string) // can transfer only string data type
	wg.Add(1)
	// go assign(ch)
	go assign(ch, &wg)
	wg.Add(1)
	// go work(ch)
	go work(ch, &wg)
	wg.Wait()
	// time.Sleep(2 * time.Second)
}

// send data to channel
// func assign(ch chan string) { // channel is sent as a reference implicitly, so no need of & operator
// 	ch <- "Make tutorial"
// 	fmt.Println("Assigned work to worker")
// }

func assign(ch chan string, wg *sync.WaitGroup) {
	ch <- "Make tutorial"
	fmt.Println("Assigned work to worker")
	wg.Done()
}

// receive data from channel
// func work(ch chan string) {
// 	fmt.Println("Waiting for some work")
// 	value := <-ch
// 	fmt.Println("Work assigned for:", value)
// }

func work(ch chan string, wg *sync.WaitGroup) {
	fmt.Println("Waiting for some work")
	value := <-ch
	fmt.Println("Work assigned for:", value)
	wg.Done()
}

// Simple Example:
// func main() {
// 	messages := make(chan string)
// 	go hello(messages)
// 	msg := <-messages
// 	fmt.Println("Received message:", msg)
// }

// func hello(messages chan string) {
// 	fmt.Println("Inside hello function")
// 	messages <- "Hello, World!"
// 	fmt.Println("Sent message to channel")
// }
