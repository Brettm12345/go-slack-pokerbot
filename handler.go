package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/nlopes/slack"
)

// interactionHandler handles interactive message response.
type interactionHandler struct {
	slackClient       *slack.Client
	verificationToken string
}

var players []slack.User
var ready []slack.User

func (h interactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("[ERROR] Invalid method: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR] Failed to read request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonStr, err := url.QueryUnescape(string(buf)[8:])
	if err != nil {
		log.Printf("[ERROR] Failed to unespace request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var message slack.AttachmentActionCallback
	if err := json.Unmarshal([]byte(jsonStr), &message); err != nil {
		log.Printf("[ERROR] Failed to decode json message from slack: %s", jsonStr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Only accept message from slack with valid token
	if message.Token != h.verificationToken {
		log.Printf("[ERROR] Invalid token: %s", message.Token)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	action := message.Actions[0]
	switch action.Name {
	case actionNewGame:
		if containsPlayer(ready, message.User) {
			ready = removePlayer(ready, message.User)
		}
		if containsPlayer(players, message.User) {
			players = removePlayer(players, message.User)
		}
		originalMessage := message.OriginalMessage
		originalMessage.Attachments[0].Text = fmt.Sprintf(listPlayers(players, ready))
		originalMessage.Attachments[0].Actions = []slack.AttachmentAction{
			{
				Name:  actionAddPlayer,
				Text:  "Join game",
				Type:  "button",
				Value: "start",
			},
		}

		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&originalMessage)
		return

	case actionAddPlayer:
		if containsPlayer(ready, message.User) {
			ready = removePlayer(ready, message.User)
		}
		if !containsPlayer(players, message.User) {
			players = append(players, message.User)
		}
		originalMessage := message.OriginalMessage

		originalMessage.Attachments[0].Text = fmt.Sprintf(listPlayers(players, ready))
		originalMessage.Attachments[0].Actions = []slack.AttachmentAction{
			{
				Name:  actionPlayerReady,
				Text:  "Ready",
				Type:  "button",
				Value: "ready",
				Style: "primary",
			},
			{
				Name:  actionNewGame,
				Text:  "Leave game",
				Type:  "button",
				Value: "start",
				Style: "danger",
			},
		}
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&originalMessage)
		return

	case actionPlayerReady:
		if !containsPlayer(ready, message.User) {
			ready = append(ready, message.User)
		}
		originalMessage := message.OriginalMessage

		text := listPlayers(players, ready)
		originalMessage.Attachments[0].Text = fmt.Sprintf(text)
		originalMessage.Attachments[0].Actions = []slack.AttachmentAction{
			{
				Name:  actionAddPlayer,
				Text:  "Cancel",
				Type:  "button",
				Value: "Cancel",
			},
			{
				Name:  actionNewGame,
				Text:  "Leave game",
				Type:  "button",
				Value: "start",
				Style: "danger",
			},
		}
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&originalMessage)
		return

	default:
		log.Printf("[ERROR] ]Invalid action was submitted: %s", action.Name)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func listPlayers(players []slack.User, ready []slack.User) string {
	var output string
	for _, player := range players {
		if containsPlayer(ready, player) {
			output += ":heavy_check_mark: "
		} else {
			output += ":x: "
		}
		output += "<@" + player.ID + ">\n"
	}
	return output
}

func containsPlayer(players []slack.User, user slack.User) bool {
	for _, player := range players {
		if player.Name == user.Name {
			return true
		}
	}
	return false
}

func removePlayer(players []slack.User, user slack.User) []slack.User {
	fmt.Println(len(players))
	if len(players) == 1 {
		return []slack.User{}
	}
	for index, player := range players {
		if user == player {
			if index == len(players) {
				return players[0 : index-1]
			}
			return append(players[:index], players[index+1])
		}
	}
	return players
}
