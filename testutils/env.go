package testutils

import (
	"os"
	"regexp"
	"srtsync/config"
)

const projectDirName = "stt-srt-sync"

func GetEnvPath() string {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	return string(rootPath) + `/.env`
}

func LoadEnv() (*config.Config, error) {
	return config.LoadConfig(GetEnvPath())
}
