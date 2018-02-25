package main

import "github.com/nlopes/slack"

// Game is a struct containing all of the game information
type Game struct {
	players []Player
	deck    *Deck
}

// AddPlayers adds a players to an existing poker game
func (g *Game) AddPlayers(users ...*slack.User) {
	for i := range users {
		g.players = append(g.players, Player{users[i], g.deck.Draw(2)})
	}
}

// NewGame creates a new poker game
func NewGame(users ...*slack.User) *Game {
	game := new(Game)
	game.deck = NewDeck()
	game.AddPlayers(users...)
	return game
}
