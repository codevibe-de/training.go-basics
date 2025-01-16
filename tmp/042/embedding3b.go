package main

import (
	"errors"
	"fmt"
)

type RequestError struct {
	error
	StatusCode int
}

// overwrites promoted code
func (r RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.error.Error())
}

func main() {
	// basic usage
	var err error = RequestError{errors.New("unavailable"), 503}
	fmt.Println(err.Error())

	// since we upcasted our error to type `error`,
	// we need to downcast for details
	reqErr, ok := err.(RequestError)
	if ok {
		fmt.Printf("request-error, status %d\n", reqErr.StatusCode)
	}
}
