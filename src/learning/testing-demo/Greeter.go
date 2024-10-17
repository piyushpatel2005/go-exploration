package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name + "!"
}

func MorningGreet(name string) string {
	return "Good morning, " + name + "! How are you doing today?"
}

func main() {
	fmt.Println(Hello("Chris"))
}
