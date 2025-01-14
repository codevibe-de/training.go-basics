package main

import (
	"fmt"
	"local/042_methods-interfaces/a_methods/stack"
)

func main() {
	fmt.Println("----------------- BOOKSHELF ------------------------")
	bookshelfDemo()
	fmt.Println("----------------- STACK ----------------------------")
	stackDemo()
}

func bookshelfDemo() {
	books := [5]book{
		{"123-456", "A. Hendker", "Wort und Bild Verlag"},
		{"345-476", "T. Müller", "Random House"},
		{"133-898", "I. Wellmann", "Random House"},
		{"423-001", "I. Wellmann", "Random House"},
		{"193-753", "T. Müller", "Wort und Bild Verlag"},
	}

	shelf := newBookshelf()
	for _, b := range books {
		err := shelf.add(b)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("Get one book:")
	fmt.Println(shelf.forIsbn("123-456"))
	fmt.Println(shelf.forIsbn("does-not-exist"))

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

func newBookshelf() bookshelf {
	bs := bookshelf{}
	bs.init()
	return bs
}

func (bs *bookshelf) add(b book) error {
	_, exists := bs.booksByIsbn[b.isbn]
	if exists {
		return fmt.Errorf("bookshelf already contains a book with ISBN `%s`", b.isbn)
	} else {
		bs.booksByIsbn[b.isbn] = b
		return nil
	}
}

func (bs *bookshelf) forIsbn(isbn string) (book, bool) {
	book, exists := bs.booksByIsbn[isbn]
	return book, exists
}

func (bs *bookshelf) all() []book {
	booksSlice := make([]book, 0, len(bs.booksByIsbn))
	for _, b := range bs.booksByIsbn {
		booksSlice = append(booksSlice, b)
	}
	return booksSlice
}

func (bs *bookshelf) init() {
	bs.booksByIsbn = make(map[string]book)
}

// ----------------- stack -------------------

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
