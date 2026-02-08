package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const config_url = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()

	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(configPath)

	if err != nil {
		return Config{}, err
	}

	defer file.Close()
	config := Config{}

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configPath := filepath.Join(homePath, config_url)

	return configPath, nil
}

func write(config Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		return err
	}

	return nil
}
