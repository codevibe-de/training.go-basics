package main

import (
	"fmt"
	"io"
)

type Foo struct {
	io.Reader
}

func main() {
	var reader io.Reader = Foo{}
	fmt.Println(reader)
	buffer := make([]byte, 100)
	reader.Read(buffer) // panic
}
