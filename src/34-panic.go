package main

import "fmt"

func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	value, ok := <-ch
	if ok {
		fmt.Println(value)
	}
	close(ch)
	// ch <- 100 // panic: send on closed channel
	close(ch) // panic: close of closed channel
}
