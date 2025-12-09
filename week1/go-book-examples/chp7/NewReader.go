package main

import (
	"fmt"
	"io"
)

type StringReader struct {
	s      string
	offset int
}

func NewReader(s string) *StringReader {
	return &StringReader{
		s:      s,
		offset: 0,
	}
}

func (s *StringReader) Read(p []byte) (int, error) {
	if s.offset >= len(s.s) {
		return 0, io.EOF
	}

	remaining := len(s.s) - s.offset
	want := len(p)
	n := min(remaining, want)

	copy(p, s.s[s.offset:s.offset+n])
	s.offset += n

	return n, nil
}

func main() {
	r := NewReader("hello")
	buf := make([]byte, 2)

	for {
		n, err := r.Read(buf)
		fmt.Printf("%q %v\n", buf[:n], err)
		if err == io.EOF {
			break
		}
	}

}
