package telegram

import (
	"strings"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/gpt3"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/models"
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
	response, err := gpt3.GetResponse(gpt3.Client, gpt3.Ctx, message.Text)
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
	msg := tgbotapi.NewMessage(message.Chat.ID, models.StartCommandString)
	_, err := b.bot.Send(msg)
	return err
}

// handleAccountCommand is handle function by getting info about user account
func (b *Bot) handleAccountCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, models.AccountCommandString)
	_, err := b.bot.Send(msg)
	return err
}

// handleSettingCommand is a handle function by change type of chat-GPT
func (b *Bot) handleSettingCommand(message *tgbotapi.Message) error {
	// create a buttons
	var row1 = []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("gpt-3.5-turbo-16k", "gpt-3.5-turbo"),
		tgbotapi.NewInlineKeyboardButtonData("gpt-3.5-turbo-instruct", "gpt-3.5-turbo-instruct"),
	}
	var row2 = []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("gpt-3.5-turbo-16k", "gpt-3.5-turbo-16k"),
		tgbotapi.NewInlineKeyboardButtonData("gpt-4", "gpt-4"),
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(row1, row2)

	// send message with buttons
	msg := tgbotapi.NewMessage(message.Chat.ID, models.SettingCommandString) // todo: Вынести в отдельную функцию
	msg.ReplyMarkup = keyboard
	_, err := b.bot.Send(msg)

	return err
}

// handleCallbackQuery is a handle function by getting data with query
func (b *Bot) handleCallbackQuery(query *tgbotapi.CallbackQuery) {
	// if query.Data == "gpt-3.5-turbo" {
	gpt3.GPType = query.Data // add to require
	// }

	deleteMsg := tgbotapi.NewDeleteMessage(query.Message.Chat.ID, query.Message.MessageID)
	b.bot.Send(deleteMsg)
}

// handleGenerateImageCommand TODO:
func (b *Bot) handleGenerateImageCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, models.GenerateImageCommandString)
	_, err := b.bot.Send(msg)
	return err
}

// handleHelpCommand is a handle function by getting info about the commands
func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, models.HelpCommandString)
	_, err := b.bot.Send(msg)
	return err
}

// handleUnknownCommand is a handle function by getting unknown command
func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, models.UnknownCommandString)
	_, err := b.bot.Send(msg)
	return err
}

// func deleteInlineButtons(c *Client, userID int64, msgID int, sourceText string) error {
// 	msg := tgbotapi.NewEditMessageText(userID, msgID, sourceText)
// 	_, err := c.client.Send(msg)
// 	if err != nil {
// 		logger.Error("Ошибка отправки сообщения", "err", err)
// 		return errors.Wrap(err, "client.Send remove inline-buttons")
// 	}
// 	return nil
// }
