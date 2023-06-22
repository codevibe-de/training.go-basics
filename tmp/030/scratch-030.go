package main

import (
	"errors"
	"fmt"
)

func main() {
	simpleErrors()
	createError()
	customError()

	//file, err := os.Open("*crazy*filename*")
	//if err != nil {
	//	fmt.Println("Cannot open file: ", err.Error())
	//	os.Exit(1)
	//}
	//fmt.Println(file)
	//
	//if file2, err2 := os.Open("*crazy*filename*"); err2 != nil {
	//	fmt.Println("Cannot open file: ", err2.Error())
	//	os.Exit(1)
	//} else {
	//	fmt.Println(file2)
	//}
}

func simpleErrors() {
	someError := errors.New("uh oh")
	anotherError := fmt.Errorf("cannot read %d bytes from %s", 123, "someFile.txt")
	fmt.Println(someError, anotherError)
}

var (
	errOutOfFuel        = errors.New("empty fuel tank")
	errNoPilot          = errors.New("missing pilot")
	errDisconnectedWire = errors.New("disconnected wire")
)

func createError() {
	fluxError := fmt.Errorf(
		"the Flux Capacitor is broken due to %w",
		errDisconnectedWire,
	)
	fmt.Println(fluxError)

	fmt.Println(errors.Is(fluxError, errOutOfFuel))
	fmt.Println(errors.Is(fluxError, errNoPilot))
	fmt.Println(errors.Is(fluxError, errDisconnectedWire))
}

type RequestError struct {
	error
	StatusCode int
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.error)
}

func customError() {
	var err error = &RequestError{errors.New("unavailable"), 503}
	fmt.Println(err)

	reqErr, ok := err.(*RequestError)
	if ok {
		fmt.Printf("request-error, status %d\n", reqErr.StatusCode)
	}
}
