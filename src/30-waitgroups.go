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
