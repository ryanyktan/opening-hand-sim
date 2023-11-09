package controller

import tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"

type Controller interface {
	// ParseDecklist parses a TCG Live decklist into an array of Pokemon Card IDs
	ParseDecklist(decklist []string) ([]tcg.PokemonCard, error)
}

func New() Controller {
	return impl{}
}

type impl struct{}
