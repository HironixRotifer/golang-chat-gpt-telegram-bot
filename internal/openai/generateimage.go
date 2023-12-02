package openai

import (
	"context"
	"encoding/base64"
	"os"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/logger"
	openai "github.com/sashabaranov/go-openai"
)

// TODO: Заменить модулем питухона
// GenerateImageResponse is a function to generate image with keywords
func GenerateImageResponse(prompt string) (*os.File, error) {
	ctx := context.Background()
	req := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize1024x1024,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
		User:           openai.GPT3Dot5Turbo0301,
	}

	defer func() {
		if err := recover(); err != nil {
			return

		}
	}()

	resp, err := client.CreateImage(ctx, req)
	if err != nil {
		return nil, err
	}

	b, err := base64.StdEncoding.DecodeString(resp.Data[0].B64JSON)
	if err != nil {
		logger.Error("Error to decode string returns the bytes represented by the base64 string: ", err)
		return nil, err

	}

	f, err := os.Create("img/image.png")
	if err != nil {
		logger.Error("Error to creates the named file: ", err)
		return nil, err

	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		logger.Error("Error to returns the number of bytes written: ", err)
		return nil, err

	}

	return f, nil
}
