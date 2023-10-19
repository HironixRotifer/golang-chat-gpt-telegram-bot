package telegram

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) generateAuthorizationLink(chatId int64) (string, error) {
	redirectURL := b.generateRedirectURL(chatID)
	requestToken, err := b.pocketClient.GetRequestToken(context.Background(), redirectURL)
	if err != nil {
		return "", err
	}

	return b.pocketClient.GetAuthorizationURL(requestToken, redirectURL)
	// authorizationLink, err := b.generatePocketLink()
	// if err != nil {
	// 	return "", err
	// }
}

func (b *Bot) generateRedirectURL(chatID int64) string{
	return fmt.Sprintf("%s?chat_id=%s", redirectURL, chatID)
}
