package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// you can skip initialization and post iteration steps in for loop

	// i := 0
	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// If we want to print only even numbers from 1 to 10, we can use continue

	// for i := 1; i <= 10; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// If we want to stop control flow when we come across first number divisible by 5 in input numbers, we can use break keyword

	// for i := 1; i <= 10; i++ {
	// 	if i%5 == 0 {
	// 		fmt.Println("Breaking out of loop when i is", i)
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// Labelled for loop

	// outer:
	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if i * j >= 25 {
	// 			break outer
	// 		}
	// 		fmt.Printf("%d * %d = %d\n", i, j, i*j)
	// 	}
	// }

}
