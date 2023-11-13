package gpt3

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/internal/logger"
	"github.com/PullRequestInc/go-gpt3"
	openai "github.com/sashabaranov/go-openai"
)

// CodexCodeDavinci002 для кодинга TODO: добавить в список используемых машин

// API key from openai
const (
	API_KEY = ""
)

var (
	client = gpt3.NewClient(API_KEY) // TODO: remove
	// type of openai engine
	EngineTypes = "text-davinci-003"
	c           = openai.NewClient(API_KEY) // TODO: remove
)

// GetResponseByQuestion is a function to answer with user question
func GetResponseByQuestion(question string) ([]string, error) {
	var msg []string

	ctx := context.Background()
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},
		MaxTokens:   gpt3.IntPtr(3000),
		Temperature: gpt3.Float32Ptr(0),
		Stop:        []string{"."},
	}, func(resp *gpt3.CompletionResponse) {
		msg = append(msg, resp.Choices[0].Text)
	})

	if err != nil {
		logger.Error("Error to get response from openai: ", err)
		return nil, err
	}

	return msg, nil
}

func GetResponseByQuestionOpenAi(question string) ([]string, error) {
	// Personalidade

	// openai.ChatMessageRoleAssistant

	ctx := context.Background()
	message := []openai.ChatCompletionMessage{
		{
			Content: question,
			Role:    "assistant",
		},
	}

	log.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA: ", message)

	stream, err := c.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo0301, // EngineTypes
		Messages:    message,
		Stop:        []string{"."},
		MaxTokens:   3000,
		Temperature: 0,
	})
	if err != nil {
		logger.Error("Error to get response from openai: ", err)
		return nil, err
	}

	resp, err := stream.Recv()
	if err != nil {
		logger.Error("Error to response processing sream recv: ", err)
		return nil, err

	}
	log.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA: ", resp)

	stream.Close()

	var msg []string
	log.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA: ", msg)

	msg = append(msg, resp.Choices[0].Delta.Content)

	return msg, nil
}

// TODO: оптимизация строк
// GenerateImageResponse is a function to generate image with keywords
func GenerateImageResponse(prompt string) (*os.File, error) {
	ctx := context.Background()
	req := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize1024x1024,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
	}

	resp, err := c.CreateImage(ctx, req)
	if err != nil {
		logger.Error("Error to generate image: ", err)
	}

	b, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)
	if err != nil {
		logger.Error("Error to decode string returns the bytes represented by the base64 string: ", err)
	}

	f, err := os.Create("image.png")
	if err != nil {
		logger.Error("Error to creates the named file: ", err)
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		logger.Error("Error to returns the number of bytes written: ", err)
	}

	return f, nil
}
