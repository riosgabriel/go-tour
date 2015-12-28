package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := make(map[string]int)

	for _, value := range strings.Fields(s) {

		if e, ok := words[value]; ok {
			words[value] = e + 1
		} else {
			words[value] = 1
		}
	}

	return words
}

func main() {
	wc.Test(WordCount)
}