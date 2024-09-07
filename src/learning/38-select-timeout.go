package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go send(ch1)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("select timeout")
	}
}

func send(ch1 chan int) {
	// time.Sleep(3 * time.Second) // execute timeout code
	ch1 <- 10
}
