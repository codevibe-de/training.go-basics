package main

import "fmt"

//type error interface {
//	Error() string
//}

type RequestError struct {
	error
	StatusCode int
}

func main() {
	var err error = &RequestError{}
	fmt.Println(err.Error()) // was passiert jetzt wohl?
}
