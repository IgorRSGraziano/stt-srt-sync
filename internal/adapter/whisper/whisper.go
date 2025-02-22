package whisper

import (
	"srtsync/internal/core/stt"

	"github.com/sashabaranov/go-openai"
)

type WhisperService struct {
	client *openai.Client
}

func NewWhisperService(apiKey string) *WhisperService {
	client := openai.NewClient(apiKey)
	return &WhisperService{client: client}
}

func (w *WhisperService) TranscribeAudio(audioPath string) ([]stt.Text, error) {
	return nil, nil
}
