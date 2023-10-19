package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		// func (i) {
		func() {
			fmt.Println(i)
			wg.Done()
		}()
		// }(i)
	}
	wg.Wait()
	fmt.Println("Finished")
}
