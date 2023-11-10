package controller

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	pkgerrors "github.com/pkg/errors"
)

// ParseDecklist takes in a decklist and returns a deck
func (i impl) ParseDecklist(decklist []string) ([]tcg.PokemonCard, error) {

	deck := make([]tcg.PokemonCard, 60)
	var currCount int

	ptcgoCodeToSetID, err := retrieveSetMapping(setMapPath)
	if err != nil {
		return nil, err
	}

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

		switch "energy" {
		case strings.ToLower(splitLine[n-1]):
			for i := 0; i < cardCount; i++ {
				deck[currCount] = tcg.PokemonCard{
					Name: strings.Join(splitLine[1:], " "),
				}
				currCount++
			}
			continue
		case strings.ToLower(splitLine[n-2]):
			// TODO: Figure out energy mapping from live to api
			for i := 0; i < cardCount; i++ {
				deck[currCount] = tcg.PokemonCard{
					Name: strings.Join(splitLine[1:n-2], " "),
				}
				currCount++
			}
			continue
		}

		setID, containsCode := ptcgoCodeToSetID[splitLine[n-2]]

		if containsCode {
			number := splitLine[n-1]

			if isTG[setID] {
				p := &number
				switch len(number) {
				case 1:
					*p = "TG0" + number
				case 2:
					*p = "TG" + number
				}
			}

			if setID == "swsh12pt5gg" {
				p := &number
				switch len(number) {
				case 1:
					*p = "GG0" + number
				case 2:
					*p = "GG" + number
				}
			}

			for i := 0; i < cardCount; i++ {
				deck[currCount] = tcg.PokemonCard{
					ID:     setID + "-" + splitLine[n-1],
					Name:   strings.Join(splitLine[1:n-2], " "),
					Number: number,
					Set: tcg.Set{
						PtcgoCode: splitLine[n-2],
					},
				}
				currCount++
			}
		} else {
			for i := 0; i < cardCount; i++ {
				deck[currCount] = tcg.PokemonCard{
					Name: strings.Join(splitLine[1:], " "),
				}
			}
		}
	}

	if currCount != 60 {
		return nil, fmt.Errorf("a valid deck has 60 cards")
	}

	return deck, nil
}

func retrieveSetMapping(file string) (map[string]string, error) {
	ptcgoCodeToSetID := make(map[string]string)

	csvFile, err := os.ReadFile(file)
	if err != nil {
		return nil, pkgerrors.WithStack(err)
	}

	csvReader := csv.NewReader(bytes.NewReader(csvFile))
	parsedCsv, err := csvReader.ReadAll()
	if err != nil {
		var parseError *csv.ParseError
		if errors.As(err, &parseError) {
			return nil, pkgerrors.WithStack(parseError)
		}
		return nil, pkgerrors.WithStack(err)
	}

	for _, row := range parsedCsv {
		// ptcgo code not in csv, skip mapping
		if row[0] == "" {
			continue
		}
		ptcgoCodeToSetID[row[0]] = row[1]
	}

	return ptcgoCodeToSetID, nil
}

var isTG = map[string]bool{
	"swsh9tg":  true,
	"swsh10tg": true,
	"swsh11tg": true,
	"swsh12tg": true,
}
