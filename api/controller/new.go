package controller

import (
	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/ryanyktan/opening-hand-sim/api/service"
)

type controller interface {
	ProcessOpeningHandSimulator(c *fiber.Ctx) error
}

func New(client tcg.Client) controller {
	return impl{
		svc: service.New(client),
	}
}

type impl struct {
	svc service.Service
}
