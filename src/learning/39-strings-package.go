package main

import (
	"fmt"
	"strings"
)

func main() {
	test := "Hello"
	// test[1] = "h" // compile error
	// fmt.Println(test)

	fmt.Println(strings.Count(test, "l"))

	hello := "Hello World!"
	fmt.Println(strings.Contains(hello, "World"))
	fmt.Println(strings.Contains(hello, "Hi"))

	fmt.Println(strings.ToLower(hello))
	fmt.Println(strings.ToUpper(hello))

	fmt.Println(strings.HasSuffix(hello, "llo"))
	fmt.Println(strings.HasSuffix(hello, "!"))
	fmt.Println(strings.HasPrefix(hello, "Hello"))

	fmt.Println(strings.Index(hello, "l"))
	fmt.Println(strings.Index(hello, "z"))
	fmt.Println(strings.Index(hello, "L")) // case-sensitive
	fmt.Println(strings.LastIndex(hello, "l"))

	fmt.Println(strings.Compare("abc", "bcd"))
	fmt.Println(strings.Compare("abcd", "abc"))
	fmt.Println(strings.Compare("abc", "abc"))

	// none of these modify original strings, immutable
	fmt.Println(strings.Replace(hello, "l", "L", 2))
	fmt.Println(strings.Replace(hello, "l", "L", -1))
	fmt.Println(strings.ReplaceAll(hello, "l", "L"))

	fmt.Println(strings.Trim(hello, "!"))
	fmt.Println("'" + strings.Trim("    Hello!!!", " !") + "'")

	fmt.Println(strings.Split(hello, " ")) // returns slice
	helloSlice := strings.Split(hello, " ")
	for _, elem := range helloSlice {
		fmt.Println(elem)
	}

	fmt.Println(strings.Join(helloSlice, ":"))
}
