package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var count = make(map[string]int)

func main() {
	dup3()
}

func dup1() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		count[input.Text()]++
	}

	for line, n := range count {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup2() {
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, count)
	} else {
		for _, args := range files {
			f, err := os.Open(args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, count)
			f.Close()
		}
	}

	for line, n := range count {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}

	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func dup3() {
	file_count := make(map[string]map[string]int)
	files := os.Args[1:]
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			if file_count[file] == nil {
				file_count[file] = make(map[string]int)
			}
			file_count[file][line]++
		}
	}

	for file, count := range file_count {
		hasdup := false
		for line, n := range count {
			if n > 1 {
				hasdup = true
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
		if hasdup {
			fmt.Println(file)
		}
	}

}
