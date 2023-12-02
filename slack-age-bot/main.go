package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/slack-io/slacker"
)

func main() {
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	bot.AddCommand(&slacker.CommandDefinition{
		Command: "hello",
		Handler: func(ctx *slacker.CommandContext) {
			ctx.Response().Reply("hi!")
		},
		HideHelp: true,
	})

	bot.AddCommand(&slacker.CommandDefinition{
		Command:     "my yob is <year>",
		Description: "yob calculator",
		Handler: func(ctx *slacker.CommandContext) {
			year := ctx.Request().Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				fmt.Println(err)
			}
			currentYear := time.Now().Year()
			age := currentYear - yob
			r := fmt.Sprintf("age is %d", age)
			ctx.Response().Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
