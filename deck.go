package main

import "math/rand"


// Deck is a slice of every playing card in the virtual deck
type Deck struct {
	cards  []Card
	suites [4]string
	ranks  [13]string
}


func init() {
	deck := new(Deck)
	deck.suites[0] = "spades"
	deck.suites[1] = "hearts"
	deck.suites[2] = "diamonds"
	deck.suites[3] = "clubs"
	deck.ranks[0] = "2"
	deck.ranks[1] = "3"
	deck.ranks[2] = "4"
	deck.ranks[3] = "5"
	deck.ranks[4] = "6"
	deck.ranks[5] = "7"
	deck.ranks[6] = "8"
	deck.ranks[7] = "9"
	deck.ranks[8] = "10"
	deck.ranks[9] = "J"
	deck.ranks[10] = "Q"
	deck.ranks[11] = "K"
	deck.ranks[12] = "A"
	for suiteNum, suite := range deck.suites {
		for rankNum, rank := range deck.ranks {
			deck.cards.append(Card(suite, rank))
		}
	}
}

func draw(d Deck, num int) []Card {
	cards := make([]Card, num)
	for i := 0; i < num; i++ {
		cardIndex = rand.Intn(len(deck))
		cards.append(deck[cardIndex])
		deck[cardIndex] = deck[len(deck)-1]
		deck = deck[:len(deck)-1]
	}
	return cards
}

func shuffle(d Deck, cards ...Card) Deck {
	for _, card := range cards {
		d.cards.append(card)
	}
	for i := range d.cards {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}
