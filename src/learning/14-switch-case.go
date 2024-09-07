package main

import "fmt"

func main() {
	// normal switch case
	var month string = "March"
	switch month {
	case "January":
		fmt.Println("This is January")
	case "February", "March": // check for multiple values
		fmt.Println("This is feb or March")
	default:
		fmt.Println("This is neither first or second month of the year.")
	}

	// use of fallthrough will move control to next case block
	switch month {
	case "January":
		fmt.Println("This is January")
	case "February", "March": // check for multiple values
		fmt.Println("This is feb or March")
		fallthrough
	default:
		fmt.Println("This is neither first or second month of the year. It is", month)
	}

	// using switch case as if..else
	var num int = 9
	switch {
	case num > 100:
		fmt.Println("The number is greater than 100")
	case num > 50:
		fmt.Println("The number is greater than 50")
	case num > 10:
		fmt.Println("The number is greater than 10")
	default:
		fmt.Println("The number is ", num)
	}

}
