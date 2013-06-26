package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(s)
	for _, v := range words {
		// wordMap[v] will return 0 if the word is not yet in the map
		wordMap[v] = wordMap[v] + 1
	}
	return wordMap
}

func main() {
	wc.Test(WordCount)
}
