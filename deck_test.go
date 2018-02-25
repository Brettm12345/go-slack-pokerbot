package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeck(t *testing.T) {
	deck := NewDeck()
	assert.Equal(t, len(deck.cards), 52, "The deck was not successfully created")
	cards := deck.Draw(5)
	assert.Equal(t, len(cards), 52-len(deck.cards), "Cards were not removed from the deck")
	deck.Shuffle(cards...)
	assert.Equal(t, len(deck.cards), 52, "Cards were not added back to the deck")
}
