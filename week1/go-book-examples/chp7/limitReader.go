package main

import (
	"fmt"
	"io"
	"strings"
)

type limitReader struct {
	r io.Reader
	n int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{
		r: r,
		n: n,
	}
}

func (l *limitReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}

	allowed := l.n
	want := len(p)
	actual := min(allowed, int64(want))

	k, err := l.r.Read(p[:actual])
	l.n -= int64(k)

	return k, err
}

func main() {
	r := strings.NewReader("Hello, world")
	lr := LimitReader(r, 5)

	buf := make([]byte, 10)

	n, err := lr.Read(buf)
	fmt.Printf("%q, n=%d, err=%v\n", buf[:n], n, err)

	n, err = lr.Read(buf)
	fmt.Printf("second read: n=%d, err=%v\n", n, err)

}
