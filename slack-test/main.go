package main

import (
	"context"
	"log"
	"os"

	"github.com/slack-io/slacker"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-bot-token")
	os.Setenv("SLACK_APP_TOKEN", "xapp-app-token")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	bot.AddCommand(&slacker.CommandDefinition{
		Command: "ping",
		Handler: func(ctx *slacker.CommandContext) {
			ctx.Response().Reply("pong")
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
