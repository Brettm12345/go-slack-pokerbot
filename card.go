package main

import s "strings"


// Card represents a playing card in a deck.
type Card struct {
	suite string
	rank  string
	image string
}

func (c Card) String() string {
	suiteEmoji := map[string]string{"Spades": "♠", "Hearts": "♥", "Diamonds": "♦", "Clubs": "♣"}
	return c.rank + suiteEmoji[s.Title(c.suite)]
}

func (c Card) ascii() string {
	return c.rank + s.Title(c.suite)
}
