package main

import (
	"fmt"
	"sync"
)

// import "time"

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go assign(ch, &wg)
	wg.Add(1)
	go work(ch, &wg)
	wg.Wait()
}

func assign(ch chan int, wg *sync.WaitGroup) {
	ch <- 1
	ch <- 2
	fmt.Println("Assigned work to worker")
	close(ch)
	wg.Done()
}

func work(ch chan int, wg *sync.WaitGroup) {
	fmt.Println("Waiting for some work")
	for value := range ch {
		fmt.Println(value)
	}
	wg.Done()
}
