package main

import (
	"fmt"
	"maps"
)

func sliceAndmaps() {
	s := make([]int, 3)
	for i := range len(s) {
		s[i] = i
	}

	t := make([]int, 0, 3)
	fmt.Println(t)

	t1 := []int{1, 3, 4, 5, 6}
	PrintSlice(t1)

	twoD := make([][]int, 3)
	for i := range len(twoD) {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}

	n1 := map[string]int{"foo": 1, "bar": 2}
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n1, n2) {
		fmt.Println("Equal")
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fact(7))
	fmt.Println(fib(7))
}

func PrintSlice[T any](s []T) {
	fmt.Printf("slice: %v, len: %d, capacity: %d\n", s, len(s), cap(s))
}
