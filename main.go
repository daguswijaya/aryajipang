package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	app.Get("/:param?", func(c *fiber.Ctx) error {
		if c.Params("param") != "" {
			return c.SendString("Your param is: " + c.Params("param"))
			// => Hello john
		}
		return c.SendString("it's empty")
	})

	log.Fatal(app.Listen("localhost:3000"))
}
