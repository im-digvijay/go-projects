package bot

import (
	"discord-ping/config"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var BotID string

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "pong")
	}
}

func Start(apiConfig config.ApiConfigData) {
	bot, err := discordgo.New("Bot " + apiConfig.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := bot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	BotID = user.ID

	bot.AddHandler(messageHandler)

	err = bot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running!")
}
