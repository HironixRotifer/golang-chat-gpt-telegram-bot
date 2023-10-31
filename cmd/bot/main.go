package main

import (
	"log"

	telegram "github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/client"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/db"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/helpers/dbutils"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/logger"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/messages"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/tasks/reportserver"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	_ "github.com/lib/pq"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6674555428:AAEXURglbwdnw3UFKuys0JSmD6W8KvKGTao")
	if err != nil {
		log.Fatal(err)
	}
	dbconn, err := dbutils.NewDBConnect(connectionStringDB)
	if err != nil {
		logger.Fatal("Ошибка подключения к базе данных:", "err", err)
	}
	// БД информации пользователей.
	userStorage := db.NewUserStorage(dbconn, 0)

	bot.Debug = true
	telegramBot := telegram.NewBot(bot, "http://localhost")

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	msgModel := messages.New(userStorage)

	reportserver.StartReportServer(msgModel)

}
