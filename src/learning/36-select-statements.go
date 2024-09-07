package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendOne(ch1)
	go sendTwo(ch2)

	select {
	case value1 := <-ch1:
		fmt.Println(value1)
		break
		fmt.Println("exit")
	case value2 := <-ch2:
		fmt.Println(value2)
		// default:
		// 	fmt.Println("Default")
	}
	time.Sleep(1 * time.Second)
}

func sendOne(ch chan string) {
	ch <- "One"
}

func sendTwo(ch chan string) {
	// ch <- "Two" // comment out to get output "One"
}
