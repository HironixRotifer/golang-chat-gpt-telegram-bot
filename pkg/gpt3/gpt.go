package gpt3

import (
	"context"

	"github.com/PullRequestInc/go-gpt3"
)

// API key from chat-gpt3
const (
	API_KEY = ""
)

// vars for GetResponse function
var (
	Client = gpt3.NewClient(API_KEY)
	Ctx    = context.Background()
	GPType = "text-davinci-003"
)

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func GetResponse(client gpt3.Client, ctx context.Context, question string) ([]string, error) {
	var msg []string
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},
		MaxTokens:   gpt3.IntPtr(3000),
		Temperature: gpt3.Float32Ptr(0),
		Stop:        []string{"."},
	}, func(resp *gpt3.CompletionResponse) {
		msg = append(msg, resp.Choices[0].Text)
		// log.Println("TEXT3: ", msg)
	})
	// log.Println("TEXT4: ", msg)

	if err != nil {
		return nil, err
	}

	return msg, nil
}