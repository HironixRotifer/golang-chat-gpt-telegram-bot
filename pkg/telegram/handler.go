package telegram

import (
	"strings"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/gpt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	// ok = "👌"
	oh = "🫢"
	// lw = "🫶"
	ct = "😺"
)

// handleCommand is a handle function to send a command for bot
// want: telegram message by "/start"
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

// handleMessage is a handle function to send a message for bot
// Exampe: "Hi, what`s up?"
func (b *Bot) handleMessage(message *tgbotapi.Message) {
	// create a new temp message
	msgTemp := tgbotapi.NewMessage(message.Chat.ID, "Please wait while I process your question..."+ct)
	id, _ := b.bot.Send(msgTemp)

	// get response from chat-gpt3
	response, err := gpt.GetResponse(gpt.Client, gpt.Ctx, message.Text)
	if err != nil {
		message.Text = err.Error()
	}
	message.Text = strings.Join(response, " ")

	// check the response for an empty value
	if message.Text == "" {
		message.Text = "Please try again" + oh
	}

	// remove a temp message
	deleteMsg := tgbotapi.NewDeleteMessage(message.Chat.ID, id.MessageID)
	b.bot.Send(deleteMsg)
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	b.bot.Send(msg)
}

// handleStrartCommand is handle function to start a bot
func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "what's up? my name is Rosetta. i could help you with anything, but only when i feel like it.")
	_, err := b.bot.Send(msg)
	return err
}

// handleUnknownCommand is a handle function by getting unknown command
// Send a message: "I don`t know this command ;(" if the command is not in the list
func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don`t know this command ;(")
	_, err := b.bot.Send(msg)
	return err
}
