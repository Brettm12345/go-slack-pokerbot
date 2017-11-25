package pokerbot

import "math/rand"

// Deck is a slice of every playing card in the virtual deck
type Deck struct {
	cards []Card
}

// Shuffle shuffles the deck and takes cards as an argument to shuffle into the deck
func (d *Deck) Shuffle(cards ...Card) {
	d.cards = append(d.cards, cards...)
	for i := range d.cards {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// Draw returns an arbitrary number of cards and removes them from the deck
func (d *Deck) Draw(num int) []Card {
	d.Shuffle()
	cards := d.cards[len(d.cards)-num:]
	d.cards = d.cards[:len(d.cards)-num]
	return cards
}

// NewDeck creates a new deck
func NewDeck() *Deck {
	deck := new(Deck)
	suites := [4]string{"spades", "hearts", "diamonds", "clubs"}
	ranks := [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	for _, suite := range suites {
		for _, rank := range ranks {
			card := Card{suite, rank}
			deck.cards = append(deck.cards, card)
		}
	}
	return deck
}
