package main

import "fmt"

func main() {
	msgch := make(chan string, 128)
	msgch <- "A"
	msgch <- "B"
	msgch <- "C"
	close(msgch) // to avoid deadlock
	
	// for msg := range msgch {
	// 	fmt.Println(msg)
	// }

	// infinite loop
	for {
		msg, ok := <-msgch
		if !ok {
			break
		}
		fmt.Println(msg)
	}

	// msgch <- "D" // send on closed channel

	// for msg := range msgch {
	// 	fmt.Println(msg)
	// }
}