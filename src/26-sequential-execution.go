package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i <= 100; i++ {
		printSequence(i)
	}
	elapsed := time.Since(start)
	fmt.Println("Elapsed time", elapsed)
}

func printSequence(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Sequence", i)
}
