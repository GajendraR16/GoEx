package main

import (
	"fmt"
	"io"
	"os"
)

type writerCounter struct {
	originalwriter io.Writer
	count          int64
}

func (wc *writerCounter) Write(p []byte) (int, error) {
	n, err := wc.originalwriter.Write(p)
	if err != nil {
		return 0, err
	}
	wc.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newWriter := &writerCounter{
		originalwriter: w,
		count:          0,
	}
	return newWriter, &newWriter.count
}

func main() {
	// 1. Define the underlying writer (os.Stdout prints to the console)
	underlyingWriter := os.Stdout

	// 2. Call CountingWriter to get the wrapper and the counter pointer
	countingWrapper, countPtr := CountingWriter(underlyingWriter)

	// 3. Write data using the new wrapper
	data1 := "Hello, Go programmers! 🚀\n"
	n1, _ := countingWrapper.Write([]byte(data1))

	// 4. Write more data using fmt.Fprintf (which uses the Write method internally)
	data2 := "The counter is working."
	n2, _ := fmt.Fprintf(countingWrapper, "%s\n", data2)

	// 5. Print the individual write counts and the total count
	fmt.Printf("\n--- Status ---\n")
	fmt.Printf("Wrote %d bytes for the first string.\n", n1)
	fmt.Printf("Wrote %d bytes for the second string.\n", n2)

	// The value pointed to by countPtr is the cumulative total
	fmt.Printf("Total bytes written to the wrapper: %d\n", *countPtr)
}
