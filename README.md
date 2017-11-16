# Slack interactive poker bot in Golang [![Go Report Card](https://goreportcard.com/badge/github.com/Brettm12345/go-slack-pokerbot)](https://goreportcard.com/report/github.com/Brettm12345/go-slack-pokerbot)

This is a slack bot to dealing and managing Texas Hold Em' games. We intend to use [slack interactive messages](https://slackhq.com/get-more-done-with-message-buttons-5fa5b283a59?ref=producthunt) to make it easier for the player to respond to the bot.

## Usage

To run this bot, you need to set the following env vars,

```bash
export BOT_ID=""             // you can get this after create a bot user (via slack app management console)
export BOT_TOKEN="xoxb-***"      // you can get this after create a bot user (via slack app management console)
export VERIFICATION_TOKEN="***"  // you can get this after enable interactive message (via slack app management console)
export CHANNEL_ID="C***"         // bot reacts only this channel
```

To run this,

```bash
$ dep ensure
$ go build -o bot && ./bot
```

To run this local, use `ngrok` (See more about it [here](https://api.slack.com/tutorials/tunneling-with-ngrok)) and set it for interactive message requests endpoint.

## Disclaimer
Please note this still a work in progress. We would love any contributions but the project is not yet functional.

## Thank You
- https://github.com/tcnksm/go-slack-interactive
- https://github.com/CharlieHess/slack-poker-bot
