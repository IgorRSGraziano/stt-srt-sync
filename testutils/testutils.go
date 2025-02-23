package testutils

import (
	"os"
	"regexp"
	"srtsync/config"
)

const projectDirName = "stt-srt-sync"

func GetEnvPath() string {
	rootPath := GetProjectDir()
	return string(rootPath) + `/.env`
}

func LoadEnv() (*config.Config, error) {
	return config.LoadConfig(GetEnvPath())
}

func GetProjectDir() string {
	re := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	cwd, _ := os.Getwd()
	rootPath := re.Find([]byte(cwd))
	return string(rootPath)
}

func GetTestDataFilePath(fileName string) string {
	rootPath := GetProjectDir()
	return rootPath + "/testutils/_testdata/" + fileName
}
