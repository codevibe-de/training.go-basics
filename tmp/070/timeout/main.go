package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	response := performRequest("https://www.example.com", 1)
	fmt.Println(response)
}

func performRequest(url string, timeout int) *http.Response {
	result := make(chan *http.Response, 1) // buffer for late request
	go func() {
		result <- networkRequest(url)
	}()
	select {
	case data := <-result:
		return data
	case <-time.After(time.Duration(timeout) * time.Second):
		return nil
	}
}

func networkRequest(url string) *http.Response {
	response, _ := http.Get(url)
	return response
}
