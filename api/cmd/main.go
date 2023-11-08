package main

import (
	"log"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

func main() {
	// If an empty string is used here, you can stil use the API with stricter limits.
	// See: https://docs.pokemontcg.io/#documentationrate_limits
	c := tcg.NewClient("your_api_key")

	cards, err := c.GetCards(
		request.Query("name:jirachi", "types:psychic"),
		request.PageSize(5),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, card := range cards {
		log.Println(card.Name)
	}
}
