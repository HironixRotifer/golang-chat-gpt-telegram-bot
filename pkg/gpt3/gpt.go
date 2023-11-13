package gpt3

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	"github.com/PullRequestInc/go-gpt3"
	openai "github.com/sashabaranov/go-openai"
)

// API key from chat-gpt3
const (
	API_KEY = ""
)

// vars for GetResponse function
var (
	client = gpt3.NewClient(API_KEY)
	Ctx    = context.Background()
	GPType = "text-davinci-003"
	c      = openai.NewClient(API_KEY)
)

type NullWriter int

func (NullWriter) Write([]byte) (int, error) { return 0, nil }

// GetResponse is a function to answer with user question
func GetResponseByQuestion(ctx context.Context, question string) ([]string, error) {
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

// func GetResponseByQuestion(ctx context.Context, question string) ([]string, error) {
// 	message := []openai.ChatCompletionMessage{
// 		{
// 			Content: question,
// 			Role:    "assistant",
// 		},
// 	}
// 	var msg []string

// 	log.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA: ", message)

// 	stream, err := c.CreateChatCompletionStream(ctx, openai.ChatCompletionRequest{
// 		Model:       openai.GPT3Dot5Turbo0301,
// 		Messages:    message,
// 		Stop:        []string{"."},
// 		MaxTokens:   3000,
// 		Temperature: 0,
// 	})
// 	if err != nil {
// 		log.Printf("Error to get question: %v", err)
// 		return nil, err
// 	}

// 	resp, err := stream.Recv()
// 	if err != nil {
// 		log.Printf("Error to get question: %v", err)
// 		return nil, err

// 	}
// 	log.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA: ", resp)

// 	stream.Close()

// 	log.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA: ", msg)

// 	msg = append(msg, resp.Choices[0].Delta.Content)

// 	return msg, nil

// }

// TODO: оптимизация строк
// GenerateImageResponse is a function to generate image with keywords
func GenerateImageResponse(ctx context.Context, prompt string) (*os.File, error) {
	req := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize1024x1024,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
	}

	resp, err := c.CreateImage(Ctx, req)
	if err != nil {
		log.Printf("Error to generate image: %v", err) // add logger
	}

	b, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)
	if err != nil {
		log.Printf("Error to generate image: %v", err) // add logger
	}

	f, err := os.Create("image.png")
	if err != nil {
		log.Printf("Error to generate image: %v", err) // add logger
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		log.Printf("Error to generate image: %v", err) // add logger
	}

	return f, nil
}
