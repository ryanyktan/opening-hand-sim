package main

import (
	"log"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/ryanyktan/opening-hand-sim/api/handler"
)

const apiKey string = "f028221e-b952-4b18-9f83-22d507a4ed7b"

func main() {
	if err := run(); err != nil {
		log.Println(err)
	}
}

func run() error {
	app := fiber.New()

	client := tcg.NewClient(apiKey)
	handler := handler.New(client)

	app.Post("/gethands", handler.ProcessOpeningHandSimulator)

	if err := app.Listen(":3000"); err != nil {
		return err
	}

	return nil
}
