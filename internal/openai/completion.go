package openai

import (
	"context"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/logger"
	openai "github.com/sashabaranov/go-openai"
)

const (
	API_KEY = "" //
)

var (
	// type of openai engine
	EngineTypes = "gpt-3.5-turbo-0301"
	client      = openai.NewClient(API_KEY)
)

// GetResponseByQuestion is a function to answer with user question
func GetResponseByQuestionOpenAi(question string) ([]string, error) {
	ctx := context.Background()

	stream, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		// EngineTypes
		Model: openai.GPT3Dot5Turbo16K0613,
		// Maximum words of response
		MaxTokens: 240,
		// Message
		Messages: []openai.ChatCompletionMessage{
			{
				Content: question,
				Role:    "assistant",
			},
		},
		// To summarize, Top-p is a balance between response quality and variety, allowing you to customize the API's behavior to suit your particular use case.
		// TopP: ,
		// Is a hyperparameter used to control the balance between diverse responses and quality.
		Temperature: 0.7,
	})

	if err != nil {
		logger.Error("Error to get response from openai: ", err)
		return nil, err
	}

	resp := stream.Choices
	if err != nil {
		logger.Error("Error to response processing sream recv: ", err)
		return nil, err
	}

	var msg []string
	msg = append(msg, resp[0].Message.Content)

	return msg, nil
}
