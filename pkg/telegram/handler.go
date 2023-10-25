package telegram

import (
	"log"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/gpt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Функция отправки команды
func (b *Bot) handleCommand(message *tgbotapi.Message) error {

	switch message.Command() {
	case commandStart:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

// Функция отправки сообщения
func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	var s string

	// Проверка текста на пустоту
	if message.Text == "" {
		message.Text = "Text is empty"
	}
	log.Println("TEXT: ", message.Text)

	response := gpt.GetResponse(gpt.Client, gpt.Ctx, message.Text)
	for _, v := range response {
		s += v
	}

	// response := gpt.Test(gpt.Client, gpt.Ctx, message.Text)

	message.Text = s
	if message.Text == "" {
		message.Text = "Повторите попытку"
	}
	log.Println("TEXT2: ", message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)

	b.bot.Send(msg)

	if message.Text == "" {
		message.Text = "Повторите попытку"
	}
}

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	// authLink, err := b.generateAuthorizationLink(message.Chat.ID)
	// if err != nil {
	// 	return err
	// }
	msg := tgbotapi.NewMessage(message.Chat.ID, "hello")

	_, err := b.bot.Send(msg)
	return err
}

// Функция неизвестной комманды
// Возвращает "Я не знаю такой команды" если команды нет в списке
func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")

	_, err := b.bot.Send(msg)
	return err
}
