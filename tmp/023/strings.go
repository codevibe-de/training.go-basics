package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hey ho let's GO!"

	var punctPresent = strings.ContainsAny(s, "?!.")
	var indexOf = strings.Index(s, "ho")
	var trimmed = strings.Trim(s, " ?!.")
	var noWhitespace = strings.ReplaceAll(s, " ", "")

	var bldr *strings.Builder = &strings.Builder{}
	fmt.Fprintln(bldr, "Add this string")

	var rdr *strings.Reader = strings.NewReader("Read me")

	fmt.Println(punctPresent, indexOf, trimmed, noWhitespace, rdr)
}
