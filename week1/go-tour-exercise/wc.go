package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	for _, word := range strings.Fields(s) {
		m[word]++
	}
	return m
}

func main() {
	fmt.Print(WordCount("I am learning Go!"))
}
