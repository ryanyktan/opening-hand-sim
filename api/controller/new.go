package controller

import (
	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/ryanyktan/opening-hand-sim/api/service"
)

type handler interface {
	ProcessOpeningHandSimulator(c *fiber.Ctx) error
}

func New(client tcg.Client) handler {
	return impl{
		svc: service.New(client),
	}
}

type impl struct {
	svc service.Service
}
