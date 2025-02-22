package whisper_test

import (
	"srtsync/internal/adapter/whisper"
	"srtsync/testutils"
	"testing"
)

func TestInstance(t *testing.T) {
	config, err := testutils.LoadEnv()

	if err != nil {
		t.Error("Error loading config")
		return
	}

	apiKey := config.OpenAIAPIKey

	if apiKey == "" {
		t.Error("API key is empty")
		return
	}

	whisperService := whisper.NewWhisperService(apiKey)

	if whisperService == nil {
		t.Error("Error creating whisper service")
	}
}
