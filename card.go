package main

import (
	"fmt"
	"strings"
)

// Card represents a playing card in a deck.
type Card struct {
	suite string
	rank  string
	image string
}

func (c Card, ascii bool) String() string {
	suite = strings.Title(c.suite)
	if ascii {
		return c.rank + suite
	}
	suiteEmoji := map[string]string{"Spades": "♠", "Hearts": "♥", "Diamonds": "♦", "Clubs": "♣"}
	return c.rank + suiteEmoji[suite]
}
