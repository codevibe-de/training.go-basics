package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	client.SetHeader("Content-Type", "application/json")
	client.SetHeader("Accept", "application/json")

	type GetPostsResult struct {
		// ...
	}
	rsp, _ := client.R().
		SetQueryString("details=true&page=1").
		SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
		SetResult(GetPostsResult{}).
		Get("https://jsonplaceholder.typicode.com/posts")
	fmt.Println(rsp.Status())
	fmt.Println(string(rsp.Body()[0:100]))

	rsp, _ = client.R().
		SetBody(`{"productId":"S-123", "name":"Pizza Funghi"}`).
		Post("https://myapp.com/login")
}
