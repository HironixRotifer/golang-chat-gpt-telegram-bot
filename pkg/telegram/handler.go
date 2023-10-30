package telegram

import (
	"strings"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/gpt"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// constants with commands to bot
const (
	commandStart          = "start"   // command to start bot
	commandAccount        = "account" // command to get account info
	commandSettings       = "setting" // command to set type of bot
	commandGeneratedImage = "genimg"  // command to generate image by keywords
	commandHelp           = "help"    // command to get help list all commands
)

var (
	// ok = "👌"
	oh = "🫢"
	// lw = "🫶"
	ct = "😺"
)

// handleCommand is a handle function to send a command for bot
func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	case commandAccount:
		return b.handleAccountCommand(message)
	case commandSettings:
		return b.handleSettingCommand(message)
	case commandGeneratedImage:
		return b.handleGenerateImageCommand(message)
	case commandHelp:
		return b.handleHelpCommand(message)
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

// handleAccountCommand is handle function by getting info about user account
func (b *Bot) handleAccountCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "todo:")
	_, err := b.bot.Send(msg)
	return err
}

// handleSettingCommand is a handle function by change type of chat-GPT
func (b *Bot) handleSettingCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "todo:")

	// создаем кнопки
	btn1 := tgbotapi.NewInlineKeyboardButtonData("gpt-3.5-turbo", "btn1")
	btn2 := tgbotapi.NewInlineKeyboardButtonData("gpt-3.5-turbo-instruct", "gpt-3.5-turbo-instruct")
	btn3 := tgbotapi.NewInlineKeyboardButtonData("gpt-3.5-turbo-16k", "btn3")
	btn4 := tgbotapi.NewInlineKeyboardButtonData("gpt-4", "gpt-4")
	row1 := tgbotapi.NewInlineKeyboardRow(btn1, btn2)
	row2 := tgbotapi.NewInlineKeyboardRow(btn3, btn4)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2)

	// отправляем сообщение с кнопками // Запихнуть в файлик строки
	msg = tgbotapi.NewMessage(message.Chat.ID, models.SettingCommandString)

	msg.ReplyMarkup = keyboard

	// Удаление сообщения
	// изменение переменной
	// поместить строки в файл

	_, err := b.bot.Send(msg)
	return err
}

// handleGenerateImageCommand TODO:
func (b *Bot) handleGenerateImageCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "todo:")
	_, err := b.bot.Send(msg)
	return err
}

// handleHelpCommand is a handle function by getting info about the commands
func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "todo:")
	_, err := b.bot.Send(msg)
	return err
}

// handleUnknownCommand is a handle function by getting unknown command
func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "I don`t know this command ;(")
	_, err := b.bot.Send(msg)
	return err
}
