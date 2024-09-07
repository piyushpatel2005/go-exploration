package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// Finding data type: 1
	var str string = "hello there"
	var isTrue bool = true
	var age int = 20
	var amount float32 = 23.23

	fmt.Printf("The type of str is %T \n", str)
	fmt.Printf("The type of isTrue is %T\n", isTrue)
	fmt.Printf("The type of age is %T\n", age)
	fmt.Printf("The type of amount is %T\n", amount)

	// Finding data type: 2
	fmt.Printf("Type: %v \n", reflect.TypeOf(str))
	fmt.Printf("Type: %v\n", reflect.TypeOf(20.99))

	var i int = 25
	var convertedInt float64 = float64(i)
	fmt.Printf("%.2f\n", convertedInt)

	var f float64 = 39.99
	var convertedFloat int = int(f) // truncates the number, removes the decimal point values
	fmt.Printf("%d\n", convertedFloat)

	// below results in compilation error
	// var s str = "hello"
	// var strToInt int = int(str)
	// fmt.Printf("%d\n", strToInt)

	var s string = strconv.Itoa(i)
	fmt.Printf("%q\n", s)

	var num int = 10
	str = strconv.Itoa(num)
	fmt.Printf("%q\n", str)

	str = "200"
	// str = "hello"
	count, err := strconv.Atoi(str)
	fmt.Printf("Count: %d, err: %v\n", count, err)

	cnt, err := strconv.ParseInt(str, 10, 32) // (string, base, bitsize)
	fmt.Printf("Count: %d, err: %v\n", cnt, err)
}
