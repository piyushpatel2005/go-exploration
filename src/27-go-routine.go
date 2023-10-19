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
	time.Sleep(2 * time.Second) // add this timeout so that it doesn't exit without printing results
}

func printSequence(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Sequence", i)
}
