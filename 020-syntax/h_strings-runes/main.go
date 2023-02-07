package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Strings and Runes:")

	s := "Heute lerne ich Go"
	words := strings.Split(s, " ")
	for i, w := range words {
		runes := []rune(w)
		runes[0] = unicode.ToUpper(runes[0])
		words[i] = string(runes)
	}
	camelCase := strings.Join(words, "")
	fmt.Println(camelCase)
}
