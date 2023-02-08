package main

import "fmt"

type book struct {
	isbn      string
	author    string
	publisher string
}

func main() {
	fmt.Println("Maps:")
	books := [5]book{
		{"123-456", "A. Hendker", "Wort und Bild Verlag"},
		{"345-476", "T. Müller", "Random House"},
		{"133-898", "I. Wellmann", "Random House"},
		{"423-001", "I. Wellmann", "Random House"},
		{"193-753", "T. Müller", "Wort und Bild Verlag"},
	}
	fmt.Println(books)

	var booksByIsbn map[string]book = groupBooksByIsbn(books[:])
	fmt.Println(booksByIsbn)

	var booksByAuthor map[string][]book = groupBooksByAuthor(books[:])
	fmt.Println(booksByAuthor)

	// bonus exercise calls from here on:
	booksByAuthor = groupBooksUsingKeyExtractor(books[:], createKeyExtractor(KEY_AUTHOR))
	fmt.Println(booksByAuthor)
	booksByPublisher := groupBooksUsingKeyExtractor(books[:], createKeyExtractor(KEY_PUBLISHER))
	fmt.Println(booksByPublisher)
}

func groupBooksByIsbn(books []book) map[string]book {
	res := make(map[string]book)
	for _, v := range books {
		res[v.isbn] = v
	}
	return res
}

func groupBooksByAuthor(books []book) map[string][]book {
	res := make(map[string][]book)
	for _, v := range books {
		if res[v.author] == nil {
			res[v.author] = make([]book, 0, 5)
		}
		res[v.author] = append(res[v.author], v)
	}
	return res
}

//
// bonus:
//

type keyExtractorFunc func(book) string

const (
	KEY_AUTHOR = iota
	KEY_PUBLISHER
)

func createKeyExtractor(key int) keyExtractorFunc {
	switch key {
	case KEY_AUTHOR:
		return func(b book) string { return b.author }
	case KEY_PUBLISHER:
		return func(b book) string { return b.publisher }
	default:
		panic(fmt.Sprintf("Unhandled key: %v", key))
	}
}

func groupBooksUsingKeyExtractor(books []book, keyExtractor keyExtractorFunc) map[string][]book {
	res := make(map[string][]book)
	for _, b := range books {
		key := keyExtractor(b)
		if res[key] == nil {
			res[key] = make([]book, 0, 5)
		}
		res[key] = append(res[key], b)
	}
	return res
}
