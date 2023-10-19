package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/shashkevych/go-pocket-sdk"
)

type Bot struct {
	bot          *tgbotapi.BotAPI
	pocketClient *pocket.Client
	redirectURL  string
}

// Функция создания нового бота
func NewBot(bot *tgbotapi.BotAPI, pocketClient *pocket.Client, redirectURL string) *Bot {
	return &Bot{bot: bot, pocketClient: pocketClient, redirectURL: redirectURL}
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

const (
	commandStart = "start"
	// https://youtu.be/fwkHaJkjO04?list=PLbTTxxr-hMmzSGTsO5mdYLrvKY-RZFanp&t=1634 посмотри что нужно вставить
	replyStartTemplate = ""
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.bot.Send(msg)
}

func (b *Bot) handleStartCommand(message *tgbotapi.Mesage) error {
	authLink, err := b.generateAuthrizationLink(message.Chat.ID)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID,
	fmt.Sprintf(replyStartTemplate, authLink))

	_, err = b.bot.Send(msg)
	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Mesage) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")

	_, err := b.bot.Send(msg)
	return err
}
