package main

import (
	"fmt"
	"strings"
)

type Card struct {
	suite string
	rank  string
	image string
}

func (c Card, ascii bool) String() string {
	suite = strings.Title(c.suite)
	if ascii {
		return c.rank + suite
	} else {
		suiteEmoji := map[string]string{"Spades": "♠", "Hearts": "♥", "Diamonds": "♦", "Clubs": "♣"}
		return c.rank + suiteEmoji[suite]
	}
}
