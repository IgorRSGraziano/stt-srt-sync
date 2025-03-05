package adapter

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type GPTService struct {
	client *openai.Client
}

func NewGPTService(apiKey string) *GPTService {
	client := openai.NewClient(apiKey)
	return &GPTService{client: client}
}

func (g *GPTService) Translate(text *string, targetLanguage string) (*string, error) {
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:       openai.GPT4Turbo,
		Temperature: 0.6,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: "Translate the following text to " + targetLanguage + ":"},
			{Role: openai.ChatMessageRoleUser, Content: *text},
		},
	}

	resp, err := g.client.CreateChatCompletion(ctx, req)

	if err != nil {
		return nil, err
	}

	return &resp.Choices[0].Message.Content, nil
}
