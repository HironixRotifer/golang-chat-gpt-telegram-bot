package telegram

import (
	"log"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// constants with commands to bot
const (
	commandStart          = "start"    // command to start bot
	commandAccount        = "account"  // command to get account info
	commandSettings       = "settings" // command to set type of bot
	commandHelp           = "help"     // command to get help list all commands
	commandGeneratedImage = "genimg"   // command to generate image by keywords
)

type Bot struct {
	bot             *tgbotapi.BotAPI
	tokenRepository repository.TokenRepository
	redirectURL     string
}

// Функция создания нового бота
func NewBot(bot *tgbotapi.BotAPI, tr repository.TokenRepository, redirectURL string) *Bot {
	return &Bot{bot: bot, tokenRepository: tr, redirectURL: redirectURL}
}

// Метод запуска бота
func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		return err
	}

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}

		b.handleMessage(update.Message)
	}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
