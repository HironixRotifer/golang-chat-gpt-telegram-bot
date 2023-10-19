package main

import (
	"log"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6674555428:AAEXURglbwdnw3UFKuys0JSmD6W8KvKGTao")
	if err != nil {
		log.Panic(err)
	}

	// Создание запуск бота
	bot.Debug = true
	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}
