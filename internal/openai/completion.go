package openai

import (
	"context"
	"strings"
	"time"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/logger"
	openai "github.com/sashabaranov/go-openai"
)

var (
	mapMsg = make(map[string]openai.ChatCompletionMessage, 10)
)

// GetResponseByQuestion is a function to answer with user question
func GetResponseByQuestionOpenAi(question string, FirstName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	question = strings.Replace(question, "\n", "", -1)
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		// EngineTypes
		Model: openai.GPT3Dot5Turbo16K0613,
		// Maximum words of response
		MaxTokens: 240,
		// Message
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: mapMsg[FirstName].Content,
			},
			{
				Role:    openai.ChatMessageRoleAssistant,
				Content: question,
			},
		},
		// To summarize, Top-p is a balance between response quality and variety, allowing you to customize the API's behavior to suit your particular use case.
		// TopP: ,
		// Is a hyperparameter used to control the balance between diverse responses and quality.
		Temperature: 0.7,
	})

	if err != nil {
		logger.Error("Error to get response from openai: ", err)
		return "", err
	}

	content := resp.Choices[0].Message.Content

	mapMsg[FirstName] = openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	}

	// resp := stream.Choices
	// if err != nil {
	// 	logger.Error("Error to response processing sream recv: ", err)
	// 	return nil, err
	// }

	// var msg []string
	// var msg string

	// msg1 = strings.Join(msg, " ")
	// msg = append(msg, resp[0].Message.Content)

	return content, nil
}
