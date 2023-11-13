package main

import (
	"log"
	"os"

	telegram "github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/client"
	// "github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/logger"
	db "github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/database"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/models"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/token"

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

	config := &db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := db.NewConnection(config)
	if err != nil {
		log.Fatal("could not load the database ")
	}
	err = models.MigrateUsers(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := AuthorizationServer.Start(); err != nil {
		log.Fatal(err)
	}
}
