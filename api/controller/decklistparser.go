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
	var noCode bool

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
		noCode = false
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
