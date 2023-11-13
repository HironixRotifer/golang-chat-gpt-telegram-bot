package client

import (
	"strings"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/gpt3"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	// ok = "👌"
	oh = "🫢"
	// lw = "🫶"
	ct = "😺"
)

// handleMessage is a handle function to send a bot message
// Exampe: "Hi, what`s up?"
func (b *Bot) handleMessage(message *tgbotapi.Message) {
	// create a new temp message
	msgTemp := tgbotapi.NewMessage(message.Chat.ID, "Please wait while I process your question..."+ct)
	id, _ := b.bot.Send(msgTemp)

	// get response from chat-gpt3
	response, err := gpt3.GetResponseByQuestion(gpt3.Ctx, message.Text)
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

// // handleMessage is a handle function to send image
// func (b *Bot) handleGenerateImage(message *tgbotapi.Message) {

// }

// handleCallbackQuery is a handle function by getting data with query TODO: Update
func (b *Bot) handleCallbackQuery(query *tgbotapi.CallbackQuery) {
	// if query.Data == "gpt-3.5-turbo" {
	gpt3.GPType = query.Data // add to require
	// }

	deleteMsg := tgbotapi.NewDeleteMessage(query.Message.Chat.ID, query.Message.MessageID)
	b.bot.Send(deleteMsg)
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
