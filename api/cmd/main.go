package main

import (
	"fmt"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

const apiKey string = "f028221e-b952-4b18-9f83-22d507a4ed7b"

func main() {
	// do something
	c := tcg.NewClient(apiKey)

	cards, err := c.GetCards(
		request.Query("name:Lightning Energy"),
	)
	if err != nil {
		panic(err)
	}

	for _, card := range cards {
		fmt.Println(card.ID)
	}
}
