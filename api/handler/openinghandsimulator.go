package handler

import (
	"bytes"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (i impl) ProcessOpeningHandSimulator(c *fiber.Ctx) error {
	// retrieve data from request body
	var (
		decklistData []byte
		// cardCache     map[string]string
		numberOfHands int
	)

	// prepare decklist
	deckString := bytes.NewBuffer(decklistData).String()
	preparedDecklist := strings.Split(deckString, "\n")

	// parse decklist
	deck, err := i.ctrl.ParseDecklist(preparedDecklist)
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
