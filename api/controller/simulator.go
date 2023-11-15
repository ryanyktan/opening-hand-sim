package controller

import (
	"bytes"
	"strings"
)

func (i impl) ProcessOpeningHandSimulator(decklistData []byte, cardCache map[string]string, numberOfHands int) error {
	// prepare decklist
	deckString := bytes.NewBuffer(decklistData).String()
	preparedDecklist := strings.Split(deckString, "\n")

	// parse decklist
	deck, err := i.ParseDecklist(preparedDecklist)
	if err != nil {
		return err
	}

	// make api call
	for range deck {
		// build query etc
	}

	// simulate hands/prizes
	for i := 0; i < numberOfHands; i++ {
		//shuffle and deal
	}

	// return result + error
	return nil
}
