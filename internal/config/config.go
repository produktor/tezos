package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	TelegramBotToken       string
	TelegramWebhookAddress string
	PostgresDSN            string
	MigrationsDSN          string
}

func Init() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	return &Config{
		TelegramBotToken:       os.Getenv("TELEGRAM_BOT_TOKEN"),
		TelegramWebhookAddress: os.Getenv("TELEGRAM_WEBHOOK_ADDRESS"),
		PostgresDSN:            os.Getenv("POSTGRES_DSN"),
		MigrationsDSN:          os.Getenv("MIGRATIONS_DSN"),
	}, nil
}
