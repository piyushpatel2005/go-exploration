package main

func main() {
	go func() {
		fmt.Println("An anonymous function")
	}()
	time.Sleep(1 * time.Second)
}
