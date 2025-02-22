package config_test

import (
	"srtsync/config"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := config.LoadConfig("../.env")

	if err != nil {
		t.Errorf("Error loading config: %s", err)
	}

	if cfg.OpenAIAPIKey == "" {
		t.Errorf("OpenAIAPIKey is empty")
	}
}

func TestNonExistentConfig(t *testing.T) {
	_, err := config.LoadConfig("../.env.example")

	if err == nil {
		t.Error("Error should not be nil")
	}
}
