package pokerbot

import "strings"

// Card represents a playing card in a deck.
type Card struct {
	suite string
	rank  string
}

func (c Card) message() string {
	suiteEmoji := map[string]string{"Spades": "♠", "Hearts": "♥", "Diamonds": "♦", "Clubs": "♣"}
	return c.rank + suiteEmoji[strings.Title(c.suite)]
}

func (c Card) ascii() string {
	return c.rank + " " + strings.Title(c.suite)
}

func (c Card) image() string {
	return strings.Join([]string{"cards/", strings.Title(c.rank), c.suite[:1], ".png"}, "")
}
