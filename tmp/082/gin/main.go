package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusAccepted, "We have received your greeting")
	})
	r.POST("/items", func(ctx *gin.Context) {
		var reqData struct {
			ID   string `json:"id"`
			Name string `json:"name" binding:"required"`
		}
		if ctx.Bind(&reqData) == nil {
			fmt.Printf("Got create-item request: %+v\n", reqData)
			ctx.Status(http.StatusCreated)
		}
	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
