package config

type Config struct {
	OpenAIAPIKey string `env:"OPENAI_API_KEY,required"`
}

type fieldConfig struct {
	env      string
	required bool
}
