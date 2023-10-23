package gpt

import (
	"context"

	"github.com/PullRequestInc/go-gpt3"
)

// API key from chat-gpt3
const (
	API_KEY = "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"
)

// vars for GetResponse function
var (
	Client = gpt3.NewClient(API_KEY)
	Ctx    = context.Background()
)

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

func GetResponse(client gpt3.Client, ctx context.Context, question string) []string {
	var msg []string
	err := client.CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			question,
		},
		MaxTokens:   gpt3.IntPtr(3000),
		Temperature: gpt3.Float32Ptr(0),
	}, func(resp *gpt3.CompletionResponse) {
		msg = append(msg, resp.Choices[0].Text)
	})

	if err != nil {
		return nil
	}

	return msg
}
