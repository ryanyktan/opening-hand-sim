package controller

import (
	"fmt"
	"strconv"
	"strings"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
)

// ParseDecklist takes in a decklist and returns a deck
func (i impl) ParseDecklist(decklist []string) ([]tcg.PokemonCard, error) {

	deck := make([]tcg.PokemonCard, 60)
	var currCount int
	var noCode bool

	for _, line := range decklist {
		if line == "" {
			continue
		}

		// line should come in the form: count(1-4), name, ptcgocode, set number
		// parse line: ptcgocode and set number are optional
		splitLine := strings.Split(line, " ")

		n := len(splitLine)
		cardCount, err := strconv.Atoi(splitLine[0])
		if err != nil {
			continue
		}

		var name string
		_, err = strconv.Atoi(splitLine[n-1])
		if err != nil {
			noCode = true
			name = strings.Join(splitLine[1:], " ")
		}

		if noCode {
			if name == "" {
				return nil, fmt.Errorf("line only has a number")
			}
			for i := 0; i < cardCount; i++ {
				deck[currCount] = tcg.PokemonCard{
					Name: name,
				}
			}
		} else {
			for i := 0; i < cardCount; i++ {
				deck[currCount] = tcg.PokemonCard{
					ID:     ptcgoCodeToSetID[splitLine[n-2]] + "-" + splitLine[n-1],
					Name:   name,
					Number: splitLine[n-1],
					Set: tcg.Set{
						PtcgoCode: splitLine[n-2],
					},
				}
				currCount++
			}
		}
	}

	if currCount != 60 {
		return nil, fmt.Errorf("a valid deck has 60 cards")
	}

	return deck, nil
}

// This map contains conversion information from 3 letter set codes to set ids as used in PokemonTCG API
var ptcgoCodeToSetID = map[string]string{
	// SV block
	"PAR":   "sv4",
	"MEW":   "sv3pt5",
	"OBF":   "sv3",
	"PAL":   "sv2",
	"SVI":   "sv1",
	"PR-SV": "svp",

	// SSH block
	"CRZ":    "swsh12pt5",
	"CRZ-GG": "swsh12pt5gg",
	"SIT":    "swsh12",
	"SIT-TG": "swsh12tg",
	"LOR":    "swsh11",
	"LOR-TG": "swsh11tg",
	"PGO":    "pgo",
	"ASR-TG": "swsh10tg",
	"ASR":    "swsh10",
	"BRS":    "swsh9",
	"FST":    "swsh8",
	"CEL":    "cel25",
	"EVS":    "swsh7",
	"CRE":    "swsh6",
	"BST":    "swsh5",
}
