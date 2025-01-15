package main

import (
	"fmt"
	"io"
	"strings"
)

type CountingReader struct {
	io.Reader
	BytesRead uint64
}

func (cr *CountingReader) Read(p []byte) (int, error) {
	n, err := cr.Reader.Read(p)
	cr.BytesRead += uint64(n)
	return n, err
}

func main() {
	var reader io.Reader = &CountingReader{Reader: strings.NewReader("ähm")}
	bytes := make([]byte, 10)
	reader.Read(bytes)

	countingReader, ok := reader.(*CountingReader)
	if ok {
		fmt.Println(countingReader.BytesRead) // Output: 4
	}
}
