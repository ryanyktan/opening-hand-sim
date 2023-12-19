package service

import tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"

const setMapPath string = "./api/controller/setmap/setmap.csv"

type Service interface {
	// InitSetMap makes an api call to initialise the setID to ptcgoCode
	InitSetMap() error
	// ParseDecklist parses a TCG Live decklist into an array of Pokemon Card IDs
	ParseDecklist(decklist []string) ([]tcg.PokemonCard, error)
	// GenerateOpeningHands generates an array of opening hands and prizes for a given deck
	GenerateOpeningHands(deck []tcg.PokemonCard, n int) []OpeningHand
}

func New(client tcg.Client) Service {
	return impl{
		dbApi: client,
	}
}

type impl struct {
	dbApi tcg.Client
}
