package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)

	splittedString := strings.Fields(s)

	for _, word := range splittedString {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	wc.Test(WordCount)
}
