package main

import "fmt"

type book struct {
	isbn      string
	author    string
	publisher string
}

func main() {
	fmt.Println("OO-Design:")
	books := [5]book{
		{"123-456", "A. Hendker", "Wort und Bild Verlag"},
		{"345-476", "T. Müller", "Random House"},
		{"133-898", "I. Wellmann", "Random House"},
		{"423-001", "I. Wellmann", "Random House"},
		{"193-753", "T. Müller", "Wort und Bild Verlag"},
	}
	fmt.Println(books)
}
