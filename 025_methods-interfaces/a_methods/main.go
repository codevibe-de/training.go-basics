package main

import (
	"fmt"
	"local/025_methods-interfaces/a_methods/stack"
)

func main() {
	fmt.Println("----------------- STACK ----------------------------")
	stackDemo()
	fmt.Println("----------------- BOOKSHELF ------------------------")
	bookshelfDemo()
}

func stackDemo() {
	s := stack.NewStack()
	s.Push("first")
	s.Push("second")
	fmt.Println(s.Peek()) // "second"
	s.Push("third")
	fmt.Println(s.Pop())  // "third"
	fmt.Println(s.Pop())  // "second"
	fmt.Println(s.Peek()) // "first"
}

func bookshelfDemo() {
	fmt.Println("Methods:")
	books := [5]book{
		{"123-456", "A. Hendker", "Wort und Bild Verlag"},
		{"345-476", "T. Müller", "Random House"},
		{"133-898", "I. Wellmann", "Random House"},
		{"423-001", "I. Wellmann", "Random House"},
		{"193-753", "T. Müller", "Wort und Bild Verlag"},
	}

	shelf := bookshelf{}
	for _, b := range books {
		shelf.add(b)
	}

	fmt.Println("Get one book:")
	fmt.Println(shelf.forIsbn("123-456"))

	fmt.Println("\nGet all books:")
	fmt.Println(shelf.all())
}

type book struct {
	isbn      string
	author    string
	publisher string
}

type bookshelf struct {
	booksByIsbn map[string]book
}

func (bs *bookshelf) add(b book) error {
	bs.init()
	_, exists := bs.booksByIsbn[b.isbn]
	if exists {
		return fmt.Errorf("This bookshelf contains already a book with ISBN `%s`", b.isbn)
	} else {
		bs.booksByIsbn[b.isbn] = b
		return nil
	}
}

func (bs *bookshelf) forIsbn(isbn string) book {
	bs.init()
	return bs.booksByIsbn[isbn]
}

func (bs *bookshelf) all() []book {
	bs.init()
	booksSlice := make([]book, 0, len(bs.booksByIsbn))
	for _, b := range bs.booksByIsbn {
		booksSlice = append(booksSlice, b)
	}
	return booksSlice
}

func (bs *bookshelf) init() {
	if bs.booksByIsbn == nil {
		bs.booksByIsbn = make(map[string]book)
	}
}
