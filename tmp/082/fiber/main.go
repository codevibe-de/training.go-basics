package main

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func main() {
	api := fiber.New()

	api.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusAccepted).SendString("We have received your greeting")
	})
	api.Post("/items", func(ctx *fiber.Ctx) error {
		var reqData struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}
		err := json.Unmarshal(ctx.Body(), &reqData)
		if err == nil {
			fmt.Printf("Got create-item request: %+v\n", reqData)
			ctx.SendStatus(http.StatusCreated)
		}
		return err
	})
	api.Listen(":8080")
}
