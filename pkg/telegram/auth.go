package telegram

// tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

// func (b *Bot) initAuthorizationProcess(message *tgbotapi.Message) error {
// 	authLink, err := b.generateAuthorizationLink(message.Chat.ID)
// 	if err != nil {

// 		msg := tgbotapi.NewMessage(message.Chat.ID,
// 			fmt.Sprintf(replyStartTemplate, authLink))

// 		_, err = b.bot.Send(msg)
// 		return err
// 	}
// }

// func (b *Bot) getAccesToken(chatID int64) (string, error) {
// 	return b.tokenRepository(chatID, repository.AccesTokens)
// }

// func (b *Bot) generateAuthorizationLink(chatId int64) (string, error) {
// 	redirectURL := b.generateRedirectURL(chatId, b.redirectURL)
// 	requestToken, err := b.pocketClient.GetRequestToken(context.Background(), redirectURL)
// 	if err != nil {
// 		return "", err
// 	}

// 	if err := b.tokenRepository.Save(chatID, requestToken, repository.RequestTokens); err != nil {
// 		return "", err
// 	}

// 	return b.pocketClient.GetAuthorizationURL(requestToken, redirectURL)
// 	// authorizationLink, err := b.generatePocketLink()
// 	// if err != nil {
// 	// 	return "", err
// 	// }
// }

// func (b *Bot) generateRedirectURL(chatID int64, redirectURL string) string {
// 	return fmt.Sprintf("%s?chat_id=%d", redirectURL, chatID)
// }
