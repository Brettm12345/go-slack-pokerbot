package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCard(t *testing.T) {
	card := Card{rank: "A", suite: "spades"}
	assert.Equal(t, card.Message(), "A♠", "Card string conversion error")
	assert.Equal(t, card.Image(), "cards/As.png", "Failed to get card image")
}
