package client

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	redirectURL string
}

// Create a new bot
func NewBot(bot *tgbotapi.BotAPI, redirectURL string) *Bot {
	return &Bot{bot: bot, redirectURL: redirectURL}
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

// handleUpdates updates chat, receives text and commands
func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {

		if update.Message.Voice != nil {
			b.handleVoiceMessage(update.Message)
			continue
		}

		// if update.Message.Audio != nil {
		// 	b.handleAudioMessage(update.Message)
		// 	continue
		// }

		if update.CallbackQuery != nil { // to handle a user's click on inline buttons
			b.handleCallbackQuery(update.CallbackQuery)
			continue
		}
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
