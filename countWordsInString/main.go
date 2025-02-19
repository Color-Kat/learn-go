package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	wordCounts := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		wordCounts[word] = wordCounts[word] + 1
	}

	return wordCounts
}

func main() {
	fmt.Println(WordCount("Hello, Hello, how are you how do you do"))
}
