package main

import (
	"fmt"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	svc "github.com/ryanyktan/opening-hand-sim/api/service"
)

const apiKey string = "f028221e-b952-4b18-9f83-22d507a4ed7b"

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {
	c := tcg.NewClient(apiKey)

	svc := svc.New(c)

	err := svc.InitSetMap()
	if err != nil {
		return err
	}

	return nil
}
