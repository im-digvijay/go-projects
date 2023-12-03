package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-bot-token")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channels := []string{os.Getenv("CHANNEL_ID")}
	files := []string{"stocks.csv"}

	for i := 0; i < len(files); i++ {
		params := slack.FileUploadParameters{
			Channels: channels,
			File:     files[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URLPrivate)
	}
}
