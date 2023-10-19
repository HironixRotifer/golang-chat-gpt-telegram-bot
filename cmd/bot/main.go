package main

import (
	"log"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6674555428:AAEXURglbwdnw3UFKuys0JSmD6W8KvKGTao")
	if err != nil {
		log.Fatal(err)
	}

	// Создание запуск бота
	bot.Debug = true
		// Надо вставить ключ с GetPocket
	pocketClient,  err:= pocket.NewClient("")
	if err != nil {
		log.Fatal(err)
	}
	telegramBot := telegram.NewBot(bot, pocketClient, redirectURL: "http://localhost")
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
