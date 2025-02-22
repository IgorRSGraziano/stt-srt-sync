package config

import (
	"maps"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
)

const defaultConfigFile = ".env"

func LoadConfig(files ...string) (*Config, error) {
	return loadConfig(files...)
}

func getFieldConfig(field reflect.StructField) fieldConfig {
	config := fieldConfig{}

	for _, flag := range strings.Split(field.Tag.Get("env"), ",") {
		switch flag {
		case "required":
			config.required = true
		default:
			config.env = flag
		}
	}
	return config
}

func loadConfig(files ...string) (*Config, error) {
	env := make(map[string]string)

	if len(files) == 0 {
		files = append(files, defaultConfigFile)
	}

	for _, file := range files {
		envMap, err := godotenv.Read(file)
		if err != nil {
			continue
		}

		maps.Copy(env, envMap)
	}

	config := &Config{}

	fields := reflect.TypeOf(config).Elem()

	for i := range fields.NumField() {
		cfg := getFieldConfig(fields.Field(i))

		if cfg.env == "" {
			continue
		}

		value, ok := env[cfg.env]
		if !ok && cfg.required {
			return nil, ErrRequiredField
		}

		reflect.ValueOf(config).Elem().Field(i).SetString(value)
	}

	return config, nil
}
