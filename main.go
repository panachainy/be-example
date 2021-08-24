package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Woooooooooooooooooooooooooooooooooooooooooo")
	})

	app.Get("/example", func(c *fiber.Ctx) error {
		return c.SendString("this is example")
	})

	log.Fatal(app.Listen(":3000"))
}
