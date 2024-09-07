package main

import (
	"fmt"
	"sync"
)

func main() {
	// var wg sync.WaitGroup
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	// close(ch)

	data, ok := <-ch
	fmt.Println(data, ok)
	close(ch)
	data, ok = <-ch
	fmt.Println(data, ok)
	data, ok = <-ch
	fmt.Println(data, ok)

	// wg.Add(1)
	// go receive(ch, &wg)
	// wg.Wait()
}

func receive(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Receiving data")
	data, ok := <-ch
	for ok {
		fmt.Println("Received:", data)
		data, ok = <-ch
	}
	fmt.Println("Done!")
	wg.Done()
}
