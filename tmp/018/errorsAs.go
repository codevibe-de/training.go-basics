package main

import (
	"errors"
	"fmt"
)

var badInput = "asd"

type BadInputError struct {
	input string
}

func (e *BadInputError) Error() string {
	return fmt.Sprintf("bad input: %s", e.input)
}

func validateInput(input string) error {
	if input == badInput {
		return fmt.Errorf("validateInput: %w", &BadInputError{input: input})
	}
	return nil
}

func main() {
	input := badInput

	err := validateInput(input)
	var badInputErr *BadInputError // nil pointer
	if errors.As(err, &badInputErr) {
		// pointer has been changed to point to actual data
		fmt.Printf("bad input error occured: %s\n", badInputErr)
	}
	G()
}

func G() {
	defer fmt.Println("g")
	defer func() {}()
	F()
}

func F() {
	defer fmt.Println("f")
	panic("coffee is empty")
}

// f
// g
// panic: coffee is empty
