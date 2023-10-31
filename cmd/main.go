package main

import (
	"log"

	telegram "github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/client"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/helpers/dbutils"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/logger"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/token"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/db"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/messages"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/tasks/reportserver"

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

	// БД информации пользователей.
	dbconn, err := dbutils.NewDBConnect(connectionStringDB)
	if err != nil {
		logger.Fatal("Ошибка подключения к базе данных:", "err", err)
	}
	userStorage := db.NewUserStorage(dbconn, "", 0)
	msgModel := messages.New(userStorage)
	reportserver.StartReportServer(msgModel)

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := AuthorizationServer.Start(); err != nil {
		log.Fatal(err)
	}
}
