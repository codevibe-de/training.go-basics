package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	client.SetHeader("Content-Type", "application/json")
	client.SetHeader("Accept", "application/json")

	rslt := new(struct {
		Id     int
		UserId int
		Title  string
	})
	rsp, _ := client.R().
		SetQueryString("details=true&page=1").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		SetResult(rslt).
		Get("https://jsonplaceholder.typicode.com/posts/3")
	fmt.Println(rsp.Status())
	fmt.Println(prettyPrint(*rslt))

	rsp, _ = client.R().
		SetBody(`{"productId":"S-123", "name":"Pizza Funghi"}`).
		Post("https://myapp.com/login")
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}
