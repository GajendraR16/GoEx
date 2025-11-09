package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		m[word]++
	}
	fmt.Println(m)
}
