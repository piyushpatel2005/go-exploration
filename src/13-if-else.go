package main

import "fmt"

func main() {

	// Check whether a person is eligible to drive or not

	var age int
	fmt.Println("Enter you age: ")
	fmt.Scanf("%d", &age)
	if age > 18 {
		fmt.Println("Congratulations, you can drive")
	} else {
		fmt.Printf("You will have to wait %d years before you can drive.\n", 18-age)
	}

	// Check if a person is eligible and whether he has the valid driving license.

	// var age int
	// var isDrivingLicenceValid string
	// fmt.Println("Enter you age: ")
	// fmt.Scanf("%d", &age)
	// if age > 18 {
	// 	fmt.Println("Do you have a valid driving licence? y/n")
	// 	fmt.Scanf("%s", &isDrivingLicenceValid)
	// 	if isDrivingLicenceValid == "y" || isDrivingLicenceValid == "yes" {
	// 		fmt.Println("Congratulations, you can drive")
	// 	} else {
	// 		fmt.Println("Oh no, you need valid driving licence to drive.")
	// 	}
	// } else {
	// 	fmt.Printf("You will have to wait %d years before you can drive.\n", 18-age)
	// }
}
