package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
	}
}

func run() error {
	app := fiber.New()

	app.Post("/sim", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(":3000"); err != nil {
		return err
	}

	return nil
}
