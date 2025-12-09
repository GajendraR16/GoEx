package main

import (
	"bytes"
	"fmt"
)

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	count := bytes.Count(p, []byte{'\n'})
	*l += LineCounter(count)
	return len(p), nil
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	words := bytes.Fields(p)
	*w += WordCounter(len(words))
	return len(p), nil
}

func main() {
	// Test WordCounter
	var wc WordCounter
	fmt.Fprintf(&wc, "hello world")
	fmt.Printf("Words: %d\n", wc)

	fmt.Fprintf(&wc, " more words here")
	fmt.Printf("Words: %d\n", wc)

	// Test LineCounter
	var lc LineCounter
	fmt.Fprintf(&lc, "line 1\nline 2\nline 3")
	fmt.Printf("Lines: %d\n", lc)

	fmt.Fprintf(&lc, "\nline 4\nline 5")
	fmt.Printf("Lines: %d\n", lc)
}
