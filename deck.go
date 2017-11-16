package main

import "math/rand"

suites := [4]string{"spades", "hearts", "diamonds", "clubs"}
ranks := [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "A"}
deck := make([]Card, 52)

for suiteNum, suite := range suites {
	for rankNum, rank := range ranks {
		deck[suiteNum + rankNum] = Card(suite, rank)
	}
}

func draw(num int) {
	cards := make([]Card, num)
	for (int i := 0; i < num; i++) {
		cardIndex = rand.Intn(:len(deck))
		cards.append(deck[cardIndex])
		deck[cardIndex] = deck[len(deck)-1]
		deck = deck[:len(deck)-1]
	}
	return cards
}

func shuffle(cards ...Card) {
	for _, card := range cards {
		deck.append(card)
	}
	list := rand.Perm(len(deck))
	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
}
