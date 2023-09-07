package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	HttpPort   string
	DbUser     string
	DbPassword string
	DbPort     string
	DbHost     string
	DbName     string
	KafkaPort  string
	KafkaHost  string
	KafkaTopic string
	Partition  int
	LogLevel   bool
}

func Load() Config {
	c := Config{}
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Error read .env: ", err)
	}

	c.HttpPort = cast.ToString(getDefaultKey("HTTP_PORT", "3333"))
	c.DbHost = cast.ToString(getDefaultKey("DB_HOST", "host"))
	c.DbPort = cast.ToString(getDefaultKey("DB_PORT", "5432"))
	c.DbUser = cast.ToString(getDefaultKey("DB_USER", "username"))
	c.DbPassword = cast.ToString(getDefaultKey("DB_PASSWORD", "password"))
	c.DbName = cast.ToString(getDefaultKey("DB_NAME", "db_name"))
	c.KafkaPort = cast.ToString(getDefaultKey("KAFKA_PORT", "9092"))
	c.KafkaHost = cast.ToString(getDefaultKey("KAFKA_HOST", "localhost"))
	c.KafkaTopic = cast.ToString(getDefaultKey("KAFKA_TOPIC", "TOPIC"))
	c.Partition = cast.ToInt(getDefaultKey("PARTITION", 0))
	c.LogLevel = cast.ToBool(getDefaultKey("LOG_LEVEL", true))

	return c

}

func getDefaultKey(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
