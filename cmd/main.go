package main

import (
	"log"

	telegram "github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/app"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/token"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/logger"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	logger.Info("Application is started")

	bot, err := tgbotapi.NewBotAPI("6674555428:AAEXURglbwdnw3UFKuys0JSmD6W8KvKGTao")
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Debug = true
	AuthorizationServer := token.NewAuthorizationServer("http://localhost")
	telegramBot := telegram.NewBot(bot, "http://localhost")

	// config := &db.Config{
	// 	Host:     os.Getenv("DB_HOST"),
	// 	Port:     os.Getenv("DB_PORT"),
	// 	User:     os.Getenv("DB_USER"),
	// 	Password: os.Getenv("DB_PASS"),
	// 	DBName:   os.Getenv("DB_DBNAME"),
	// 	SSLMode:  os.Getenv("DB_SSLMODE"),
	// }

	// db, err := db.NewConnection(config)
	// if err != nil {
	// 	log.Fatal("could not load the database ")
	// }
	// err = models.MigrateUsers(db)
	// if err != nil {
	// 	log.Fatal("could not migrate db")
	// }

	go func() {
		if err := telegramBot.Start(); err != nil {
			logger.Error("Failed to start telegram bot", err)
		}
	}()

	if err := AuthorizationServer.Start(); err != nil {
		logger.Error("Failed to start authorization server", err)
	}

	logger.Info("Application has terminated")
}
