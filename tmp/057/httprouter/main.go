package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func HelloHandler(wrt http.ResponseWriter, req *http.Request, prms httprouter.Params) {
	name := prms.ByName("name")
	if name == "" {
		fmt.Fprintf(wrt, "Hello!\n")
	} else {
		fmt.Fprintf(wrt, "Hello, %s!\n", name)
	}
}

func CreateItemHandler(wrt http.ResponseWriter, req *http.Request, prms httprouter.Params) {
	var reqData struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	if json.NewDecoder(req.Body).Decode(&reqData) == nil {
		fmt.Printf("Got create-item request: %+v\n", reqData)
		wrt.WriteHeader(http.StatusCreated)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/hello", HelloHandler)
	router.GET("/hello/:name", HelloHandler)
	router.POST("/items", CreateItemHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
