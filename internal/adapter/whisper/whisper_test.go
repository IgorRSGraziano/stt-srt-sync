package adapter_test

import (
	"fmt"
	"os"
	adapter "srtsync/internal/adapter/whisper"
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

	whisperService := adapter.NewWhisperService(apiKey)

	if whisperService == nil {
		t.Error("Error creating whisper service")
	}
}

func TestTranscribeAudio(t *testing.T) {
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

	whisperService := adapter.NewWhisperService(apiKey)

	if whisperService == nil {
		t.Error("Error creating whisper service")
	}

	lyrics, err := os.ReadFile(testutils.GetTestDataFilePath("lyric.txt"))

	if err != nil {
		fmt.Printf("Error reading lyric file: %v\n", err)

	}

	text, err := whisperService.GenerateSRT(testutils.GetTestDataFilePath("music.mp3"), string(lyrics))

	if err != nil {
		t.Error("Error transcribing audio")
	}

	if text == nil || *text == "" {
		t.Error("Transcribed text is nil")
	}
}
