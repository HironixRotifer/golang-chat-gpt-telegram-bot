package openai

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func SpeechToText(filePath string) string {
	ctx := context.Background()
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: filePath,
	}

	resp, err := client.CreateTranscription(ctx, req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return ""
	}

	return resp.Text
}

// func FindMusic(text string) string {
// 	GetResponseByQuestionOpenAi("please, find this track and author" + text)
// }

// func audioCaptions() {
// 	req := openai.AudioRequest{
// 		Model:    openai.Whisper1,
// 		FilePath: os.Args[1],
// 		Format:   openai.AudioResponseFormatSRT,
// 	}
// 	resp, err := client.CreateTranscription(context.Background(), req)
// 	if err != nil {
// 		fmt.Printf("Transcription error: %v\n", err)
// 		return
// 	}
// 	f, err := os.Create(os.Args[1] + ".srt")
// 	if err != nil {
// 		fmt.Printf("Could not open file: %v\n", err)
// 		return
// 	}
// 	defer f.Close()
// 	if _, err := f.WriteString(resp.Text); err != nil {
// 		fmt.Printf("Error writing to file: %v\n", err)
// 		return
// 	}
// }
