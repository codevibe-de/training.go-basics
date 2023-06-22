package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var r io.Reader
	r = strings.NewReader("Read me!")

	io.Copy(os.Stdout, r)
	io.Copy(os.Stdout, os.Stdin) // hallo echo! :)

	var writer io.Writer
	writer, _ = os.Open("blah.txt")

	var closer io.Closer
	closer, _ = os.Open("blah.txt")

	fmt.Println(closer, writer)
}
