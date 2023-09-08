package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

const (
	FileName = "app.config.yml"
)

var (
	AppConfig = &AppConfigType{}
)

func Init() error {
	var err error

	b, err := os.ReadFile(FileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, AppConfig)
	if err != nil {
		return err
	}

	return nil
}
