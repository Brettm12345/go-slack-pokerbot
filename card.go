package pokerbot

import "strings"

// Card represents a playing card in a deck.
type Card struct {
	suite string
	rank  string
}

// Message returns the formated slack message string
func (c Card) Message() string {
	suiteEmoji := map[string]string{"Spades": "♠", "Hearts": "♥", "Diamonds": "♦", "Clubs": "♣"}
	return c.rank + suiteEmoji[strings.Title(c.suite)]
}

// Ascii returns the ascii string for the card
func (c Card) Ascii() string {
	return c.rank + " " + strings.Title(c.suite)
}

// Image returns the filepath of the image associated with the card
func (c Card) Image() string {
	return strings.Join([]string{"cards/", strings.Title(c.rank), c.suite[:1], ".png"}, "")
}
