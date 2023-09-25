package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	HttpPort        string
	ChatServiceHost string
	CHatServicePort string
	CsvFilePath     string
	AuthFilePath    string
	SigningKey      string
}

func Load() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error .env file loading...", err)
	}

	return Config{
		HttpPort:        os.Getenv("HTTP_PORT"),
		ChatServiceHost: os.Getenv("CHAT_SERVICE_HOST"),
		CHatServicePort: os.Getenv("CHAT_SERVICE_PORT"),
		CsvFilePath:     os.Getenv("CSV_FILE_PATH"),
		AuthFilePath:    os.Getenv("AUTH_FILE_PATH"),
		SigningKey:      os.Getenv("SIGNINGKEY"),
	}
}
