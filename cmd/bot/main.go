package main

import (
	"log"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/repository"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/repository/boltdb"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/repository/server"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/telegram"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6674555428:AAEXURglbwdnw3UFKuys0JSmD6W8KvKGTao")
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true

	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}

	tokenRepository := boltdb.NewTokenRepository(db)
	telegramBot := telegram.NewBot(bot, tokenRepository, "http://localhost")
	AuthorizationServer := server.NewAuthorizationServer(tokenRepository, "https://t.me/WebNix_bot")

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := AuthorizationServer.Start(); err != nil {
		log.Fatal(err)
	}
}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("bolt.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.AccesTokens))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(repository.RequestTokens))
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}
