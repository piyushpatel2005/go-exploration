package main

import (
	"fmt"
	"strings"
)

func wordFrequency(text string) map[string]int {
	// TODO: implement this function
	words := strings.Split(text, " ")
	results := make(map[string]int)
	for _, word := range words {
		results[word] = strings.Count(text, word)
	}
	return results
}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	fmt.Println(wordFrequency(text))
}
