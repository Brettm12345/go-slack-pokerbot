package main

import "github.com/nlopes/slack"

// Player is a struct containing all the player details
type Player struct {
	User *slack.User
	Hand []Card
}
