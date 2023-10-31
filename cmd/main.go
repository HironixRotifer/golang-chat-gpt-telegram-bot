package main

import (
	"log"

	telegram "github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/client"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/token"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6674555428:AAEXURglbwdnw3UFKuys0JSmD6W8KvKGTao")
	if err != nil {
		log.Fatal(err)
		return
	}
	bot.Debug = true

	AuthorizationServer := token.NewAuthorizationServer("http://localhost")
	telegramBot := telegram.NewBot(bot, "http://localhost")

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := AuthorizationServer.Start(); err != nil {
		log.Fatal(err)
	}
}
