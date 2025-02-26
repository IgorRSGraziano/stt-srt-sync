package adapter

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type WhisperService struct {
	client *openai.Client
}

func NewWhisperService(apiKey string) *WhisperService {
	client := openai.NewClient(apiKey)
	return &WhisperService{client: client}
}

func (w *WhisperService) GenerateSRT(audioPath string, lyric *string) (*string, error) {
	ctx := context.Background()

	var prompt string

	if lyric != nil {
		prompt = *lyric
	} else {
		prompt = ""
	}

	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: audioPath,
		Format:   openai.AudioResponseFormatSRT,
		Prompt:   prompt,
	}

	resp, err := w.client.CreateTranscription(ctx, req)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return nil, err
	}

	return &resp.Text, nil
}
