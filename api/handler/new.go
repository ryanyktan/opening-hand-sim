package handler

import (
	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/ryanyktan/opening-hand-sim/api/controller"
)

type handler interface {
	ProcessOpeningHandSimulator(c *fiber.Ctx) error
}

func New(client tcg.Client) handler {
	return impl{
		ctrl: controller.New(client),
	}
}

type impl struct {
	ctrl controller.Controller
}
