package main

import "fmt"

type Day int

//const (
//	Sunday    Day = 1
//	Monday    Day = 2
//	Tuesday   Day = 3
//	Wednesday Day = 4
//	Thursday  Day = 5
//	Friday    Day = 6
//	Saturday  Day = 7
//)

//func findDayNumber(d Day) int {
//	return int(d) - 1
//}

const (
	Sunday Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func findDayNumber(d Day) int {
	return int(d)
}

func (d Day) String() string {
	switch d {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	default:
		return "Unknown"
	}
}

func (d Day) isWeekend() bool {
	switch d {
	case Saturday, Sunday:
		return true
	default:
		return false
	}
}

func (d Day) dayNumber() int {
	return int(d)
}

func main() {
	//fmt.Println("sunday is ", findDayNumber(Sunday))

	fmt.Println("Sunday is ", Sunday)
	fmt.Println("Monday is day number: ", Monday.dayNumber())
	fmt.Printf("Is %s a weekday? %v\n", Sunday, Sunday.isWeekend())
}
