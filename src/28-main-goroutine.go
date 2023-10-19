package main

import "fmt"
import "time"

func process1() {
	time.Sleep(1 * time.Second)
	fmt.Println("Running Process 1")
}

func process2() {
	go process1()
	time.Sleep(1 * time.Second)
	fmt.Println("Running Process 2")
}

func main() {
	go process2()
	time.Sleep(2 * time.Second)
}
