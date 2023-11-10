package controller

import tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"

const setMapPath string = "./api/controller/setmap/setmap.csv"

type Controller interface {
	// InitSetMap makes an api call to initialise the setID to ptcgoCode
	InitSetMap() error
	// ParseDecklist parses a TCG Live decklist into an array of Pokemon Card IDs
	ParseDecklist(decklist []string) ([]tcg.PokemonCard, error)
}

func New(client tcg.Client) Controller {
	return impl{
		dbApi: client,
	}
}

type impl struct {
	dbApi tcg.Client
}
