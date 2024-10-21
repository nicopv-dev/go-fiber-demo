package main

import (
	"go-fiber-demo/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/hello", handlers.Hello)

	log.Fatal(app.Listen(":3000"))
}