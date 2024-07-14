package config

import (
	"errors"
	"log"
	"os"
)

type Config struct {
	token       string
	dsn         string
	storageType string
}

func (c *Config) GetToken() string {
	return c.token
}

func (c *Config) GetDsn() string {
	return c.dsn
}

func (c *Config) StorageType() string {
	return c.storageType
}

func InitConfig() (*Config, error) {
	var cfg Config
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Println("TELEGRAM_BOT_TOKEN is not set")
		return nil, errors.New("Не найден токен бота")
	}
	cfg.token = botToken

	dbstroka := os.Getenv("DSN")
	if dbstroka == "" {
		log.Println("DSN stroka net")
		return nil, errors.New("Не найден токен dsn")
	}
	cfg.dsn = dbstroka

	stype := os.Getenv("STORAGE_TYPE")
	if stype == "" {
		log.Println("STORAGE_TYPE stroka net")
		return nil, errors.New("Не найден токен STORAGE_TYPE")
	}
	cfg.storageType = stype
	return &cfg, nil
}
