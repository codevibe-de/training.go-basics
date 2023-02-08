package main

import "fmt"

type book struct {
	isbn   string
	author string
	year   uint16
}

func main() {
	fmt.Println("Maps:")
	books := [5]book{
		{"123-456", "A. Hendker", 2002},
		{"345-476", "T. Müller", 2009},
		{"133-898", "I. Wellmann", 2017},
		{"423-001", "I. Wellmann", 2018},
		{"193-753", "T, Müller", 2017},
	}
	fmt.Println(books)
	//groupBooksByIsbn()
	//groupBooksByAuthor()
}
