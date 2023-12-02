package client

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/openai"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	// ok = "👌"
	oh = "🫢"
	// lw = "🫶"
	ct = "😺"
	sc = "😐"
)

// handleMessage is a handle function to send a bot message
// Exampe: "Hi, what`s up?"
func (b *Bot) handleMessage(message *tgbotapi.Message) {
	// create a new temp message
	msgTemp := tgbotapi.NewMessage(message.Chat.ID, "Please wait while I process your question..."+ct)
	id, _ := b.bot.Send(msgTemp)

	// get response from chat-gpt3
	response, err := openai.GetResponseByQuestionOpenAi(message.Text)
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

// handleVoiceMessage is hendle func by recive voice messages and answer them
func (b *Bot) handleVoiceMessage(message *tgbotapi.Message) error {
	var msg = tgbotapi.NewMessage(message.Chat.ID, "file is too big")
	defer func() {
		if err := recover(); err != nil {
			b.bot.Send(msg)
			return
		}
	}()

	fileID := message.Voice.FileID
	file, err := b.bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		logger.Error("Error to get voice message", err)
	}

	url := "https://api.telegram.org/file/bot" + b.bot.Token + "/" + file.FilePath
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Создаем новый файл для сохранения аудиосообщения
	audioFile, err := os.Create("audio_message.mp3")
	if err != nil {
		return err
	}
	defer audioFile.Close()

	// Копируем содержимое файла аудиосообщения из HTTP ответа в созданный файл
	_, err = io.Copy(audioFile, resp.Body)
	if err != nil {
		return err
	}

	msg = tgbotapi.NewMessage(message.Chat.ID, openai.SpeechToText(audioFile.Name()))
	b.bot.Send(msg)

	return nil
}

// handleCallbackQuery is a handle function by getting data with query TODO: Update
func (b *Bot) handleCallbackQuery(query *tgbotapi.CallbackQuery) {
	openai.EngineTypes = query.Data // add to require

	deleteMsg := tgbotapi.NewDeleteMessage(query.Message.Chat.ID, query.Message.MessageID)
	b.bot.Send(deleteMsg)
}

// TODO:
func (b *Bot) handleAudioMessage(message *tgbotapi.Message) error {
	audio := message.Audio

	// Получаем информацию о файле аудиосообщения
	fileConfig := tgbotapi.FileConfig{
		FileID: audio.FileID,
	}
	file, err := b.bot.GetFile(fileConfig)
	if err != nil {
		return err
	}

	// Скачиваем файл аудиосообщения
	url := "https://api.telegram.org/file/bot" + b.bot.Token + "/" + file.FilePath
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Создаем новый файл для сохранения аудиосообщения
	audioFile, err := os.Create("audio_message.mp3")
	if err != nil {
		return err
	}
	defer audioFile.Close()

	// Копируем содержимое файла аудиосообщения из HTTP ответа в созданный файл
	_, err = io.Copy(audioFile, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// TODO:
func (b *Bot) DeleteTempMessage(message *tgbotapi.Message) {
	deleteMsg := tgbotapi.NewDeleteMessage(message.Chat.ID, message.MessageID)
	b.bot.Send(deleteMsg)

}

// TODO:
func (b *Bot) NewTempMessage(message *tgbotapi.Message, messageText string) {
	msg := tgbotapi.NewMessage(message.Chat.ID, messageText)
	b.bot.Send(msg)
	// if err != nil {
	// 	logger.Error("Error temp message: ", err)
	// }
}
