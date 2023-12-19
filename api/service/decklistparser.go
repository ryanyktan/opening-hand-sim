package service

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
			id := basicEnergyMap[splitLine[n-1]]
			for i := 0; i < cardCount; i++ {
				deck[currCount] = tcg.PokemonCard{
					ID:   id,
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

var basicEnergyMap = map[string]string{
	"1": "sve-1", "9": "sve-1", "18": "sve-1", "27": "sve-1", "36": "sve-1", "45": "sve-1",
	"2": "sve-2", "10": "sve-2", "19": "sve-2", "28": "sve-2", "37": "sve-2", "46": "sve-2",
	"3": "sve-3", "11": "sve-3", "20": "sve-3", "29": "sve-3", "38": "sve-3", "47": "sve-3",
	"4": "sve-4", "12": "sve-4", "21": "sve-4", "30": "sve-4", "39": "sve-4", "48": "sve-4",
	"5": "sve-5", "13": "sve-5", "22": "sve-5", "31": "sve-5", "40": "sve-5", "49": "sve-5",
	"6": "sve-6", "14": "sve-6", "23": "sve-6", "32": "sve-6", "41": "sve-6", "50": "sve-6",
	"7": "sve-7", "15": "sve-7", "24": "sve-7", "33": "sve-7", "42": "sve-7", "51": "sve-7",
	"8": "sve-8", "16": "sve-8", "25": "sve-8", "34": "sve-8", "43": "sve-8", "52": "sve-8",
	"17": "sm1-172", "26": "sm1-172", "35": "sm1-172", "44": "sm1-172",
}
