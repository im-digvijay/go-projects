package main

import (
	"discord-ping/bot"
	"discord-ping/config"
	"fmt"
)

func main() {
	apiConfig, err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start(apiConfig)

	<-make(chan struct{})
	return
}
