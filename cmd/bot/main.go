package main

import (
	"log"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	// "github.com/zhashkevych/go-pocket-sdk"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6674555428:AAEXURglbwdnw3UFKuys0JSmD6W8KvKGTao")
	if err != nil {
		log.Fatal(err)
	}

	// Создание запуск бота
	bot.Debug = true
	// Надо вставить ключ с GetPocket
	// pocketClient, err := pocket.NewClient("109274-d1ae0275f274205aedbf2ba")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	telegramBot := telegram.NewBot(bot, "http://localhost")
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}

	// gpt.Ctx = context.Background()
	// gpt.Client = gpt3.NewClient(gpt.API_KEY)
	// log.Println("ЭТО ЗДЕСЬ: ", gpt.Ctx, gpt.Client, gpt.API_KEY)

}
