package client

import (
	"log"
	"time"

	models "github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/model"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/openai"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// constants with commands to bot
const (
	commandStart          = "start"         // command to start bot
	commandAccount        = "account"       // command to get account info
	commandSettings       = "setting"       // command to set type of bot
	commandGeneratedImage = "generateimage" // command to generate image by keywords
	commandHelp           = "help"          // command to get help list all commands
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

// handleGenerateImageCommand TODO: переделать архитектуру хранения
func (b *Bot) handleGenerateImageCommand(message *tgbotapi.Message) error {
	keywords := message.CommandArguments()
	if keywords == "" {
		msgTemp := tgbotapi.NewMessage(message.Chat.ID, "please write a description for your picture:\n /generateimage funny gopher")
		b.bot.Send(msgTemp)
		return nil
	}

	msgTemp := tgbotapi.NewMessage(message.Chat.ID, "Please wait while I generate your image..."+ct)
	id, _ := b.bot.Send(msgTemp)

	photo, err := openai.GenerateImageResponse(keywords)
	if err != nil {
		log.Printf("Failed to send Photo: %v", err)
	}

	if photo == nil {
		msgTemp := tgbotapi.NewMessage(message.Chat.ID, "I don't like creating this "+sc)
		b.bot.Send(msgTemp)
		deleteMsg := tgbotapi.NewDeleteMessage(message.Chat.ID, id.MessageID)
		b.bot.Send(deleteMsg)
		return nil
	}

	photoConfig := tgbotapi.NewPhotoUpload(message.Chat.ID, "img/image.png")
	deleteMsg := tgbotapi.NewDeleteMessage(message.Chat.ID, id.MessageID)
	b.bot.Send(deleteMsg)

	_, err = b.bot.Send(photoConfig)
	if err != nil {
		log.Printf("Failed to send Photo: %v", err)
	}

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
		tgbotapi.NewInlineKeyboardButtonData("gpt-3.5-turbo-0301", "gpt-3.5-turbo-0301"),
		tgbotapi.NewInlineKeyboardButtonData("code-davinci-002", "code-davinci-002"),
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

// func (b *Bot) test() {
// 	b.bot.GetChatAdministrators()

// 	b.bot.GetMe()

// 	b.bot.RestrictChatMember()

// func (b *Bot) PromoteChatMember() {

// 	b.bot.PromoteChatMember() // добавляет права администратора пользователю

// }

func (b *Bot) RestrictUser(chatID int64, userID int) {
	restrictionConfig := tgbotapi.RestrictChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			ChatID: chatID,
			UserID: userID,
		},
		UntilDate: time.Now().Add(time.Hour * 24).Unix(), // блокировка на 24 часа
	}
	_, err := b.bot.RestrictChatMember(restrictionConfig)
	if err != nil {
		log.Println("Error restricting chat member:", err)
	} else {
		msg := tgbotapi.NewMessage(chatID, "Пользователь заблокирован на 24 часа.")
		b.bot.Send(msg)
	}
}
