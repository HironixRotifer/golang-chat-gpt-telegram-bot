package client

import (
	"fmt"
	"io"
	"net/http"
	"os"

	ai "github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/openai"
	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sashabaranov/go-openai"
)

var (
	// ok = "👌"
	// oh = "🫢"
	// lw = "🫶"
	ct = "😺"
	sc = "😐"
)

// handleMessage is a handle function to send a bot message
func (b *Bot) handleMessage(message *tgbotapi.Message) {
	// create a new temp message
	id := b.newTempMessage(message.Chat.ID, "Please wait while I process your question...")

	fmt.Println("EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE ", message.From.FirstName)
	fmt.Println(message.From.ID)
	fmt.Println(message.From.ID)

	// get response from chat-gpt
	resp, err := ai.GetResponseByQuestionOpenAi(message.Text, message.From.FirstName)
	if err != nil {
		resp = err.Error()
	}

	// check the response for an empty value
	if resp == "" {
		resp = "Please try again"
	}

	// remove a temp message
	b.deleteTempMessage(message.Chat.ID, id)

	msg := tgbotapi.NewMessage(message.Chat.ID, resp)
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

	msg = tgbotapi.NewMessage(message.Chat.ID, ai.SpeechToText(audioFile.Name()))
	b.bot.Send(msg)

	return nil
}

// handleCallbackQuery is a handle function by getting data with query TODO: Update
func (b *Bot) handleCallbackQuery(query *tgbotapi.CallbackQuery) {

	switch query.Data {
	case "gpt-3.5-turbo-0301":
		ai.EngineTypes = openai.GPT3Dot5Turbo16K0613
	case "gpt-3.5-turbo-16k":
		ai.EngineTypes = openai.GPT3Dot5Turbo16K
	case "code-davinci-002":
		ai.EngineTypes = openai.CodexCodeDavinci002
	case "gpt-4":
		ai.EngineTypes = openai.GPT4
	default:
		ai.EngineTypes = openai.GPT3Dot5Turbo16K0613
	}

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
func (b *Bot) deleteTempMessage(chatID int64, messageID int) {
	deleteMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
	b.bot.Send(deleteMsg)
}

// TODO:
func (b *Bot) newTempMessage(chatID int64, messageText string) int {
	msg := tgbotapi.NewMessage(chatID, messageText)
	id, _ := b.bot.Send(msg)

	return id.MessageID
}
