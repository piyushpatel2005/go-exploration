package main

import (
	"fmt"
	"time"
)



func main() {
	resultch := make(chan string) // unbuffered channel
	// A channel in Golang will always block if it's full

	result := <-resultch
	resultch <- fetchResource(1)
	// result := <-resultch
	fmt.Println(result)
}

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("the result %d", n)
}	