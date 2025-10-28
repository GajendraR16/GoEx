package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args
	start := time.Now()
	ineff(args)
	end := time.Since(start)
	fmt.Println("Inefficient time = ", end)

	start = time.Now()
	eff(args)
	end = time.Since(start)
	fmt.Println("efficient time = ", end)
}

func ineff(s []string) {
	for i, v := range s[1:] {
		fmt.Println(i, v)
	}
}

func eff(s []string) {
	fmt.Println(strings.Join(s[1:], " "))
}
