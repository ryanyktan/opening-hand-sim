package main

import (
	"log"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

const apiKey string = "f028221e-b952-4b18-9f83-22d507a4ed7b"

func main() {
	// If an empty string is used here, you can stil use the API with stricter limits.
	// See: https://docs.pokemontcg.io/#documentationrate_limits
	c := tcg.NewClient(apiKey)

	cards, err := c.GetCards(
		request.Query("name:jirachi", "types:psychic"),
		request.PageSize(5),
	)
	if err != nil {
		log.Fatal(err)
	}

	if len(cards) == 0 {
		log.Println("no cards found")
	}

	for _, card := range cards {
		log.Println(card.Name)
		log.Println(card.Number)
		log.Println(card.Set.PtcgoCode)
	}
}
