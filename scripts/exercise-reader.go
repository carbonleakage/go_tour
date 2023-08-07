package main

import (
	"fmt"
	"io"
	// "golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (m MyReader) Read(b []byte) (int, error) {
	if cap(b) < 1 {
		return 0, io.EOF
	}
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	// reader.Validate(MyReader{})
	reader := MyReader{}
	b := make([]byte, 100)

	n, err := reader.Read(b)
	fmt.Printf("%d characters read, error raised %q. The actual array is %v\n", n, err, b)
}
