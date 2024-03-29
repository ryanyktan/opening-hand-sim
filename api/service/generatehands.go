package service

import (
	"math/rand"
	"time"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
)

func (i impl) GenerateOpeningHands(deck []tcg.PokemonCard, n int) []OpeningHand {
	res := make([]OpeningHand, n)

	for i := 0; i < n; i++ {
		var hand, prizes []tcg.PokemonCard
		mulligans := 0

		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		// some logic
		for {
			for j := 0; j < 7; j++ {
				randomInt := r.Intn(60 - j)
				if randomInt == 59-j {
					continue
				}
				deck[randomInt], deck[59-j] = deck[59-j], deck[randomInt]
			}

			// check for valid opening hand
			hand = deck[53:]
			// if hand is valid, generate prizes and break
			if isValidHand(hand) {
				for j := 7; j < 13; j++ {
					randomInt := r.Intn(60 - j)
					if randomInt == 59-j {
						continue
					}
					deck[randomInt], deck[59-j] = deck[59-j], deck[randomInt]
				}
				prizes = deck[47:53]
				break
			}

			// otherwise, increase mulligan
			mulligans++
		}

		res = append(res, OpeningHand{
			Hand:      hand,
			Prizes:    prizes,
			Mulligans: mulligans,
		})
	}
	return res
}

type OpeningHand struct {
	Hand      []tcg.PokemonCard
	Prizes    []tcg.PokemonCard
	Mulligans int
}

// A valid hand contains 7 cards and at least 1 basic pokemon.
func isValidHand(hand []tcg.PokemonCard) bool {
	if len(hand) != 7 {
		return false
	}

	for _, card := range hand {
		for _, subtype := range card.Subtypes {
			if subtype == "Basic" {
				return true
			}
		}
	}
	return false
}
