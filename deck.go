package pokerbot

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
	for _, suite := range deck.suites {
		for _, rank := range deck.ranks {
			card := Card{suite, rank, rank + suite[0:1] + ".png"}
			deck.cards = append(deck.cards, card)
		}
	}
}

func draw(d Deck, num int) []Card {
	cards := make([]Card, num)
	for i := 0; i < num; i++ {
		cardIndex := rand.Intn(len(d.cards))
		cards = append(cards, d.cards[cardIndex])
		d.cards[cardIndex] = d.cards[len(d.cards)-1]
		d.cards = d.cards[:len(d.cards)-1]
	}
	return cards
}

func shuffle(d Deck, cards ...Card) Deck {
	for _, card := range cards {
		d.cards = append(d.cards, card)
	}
	for i := range d.cards {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
	return d
}
