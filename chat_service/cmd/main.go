package main

import (
	"github.com/kafka_example/chat_service/config"
	"github.com/kafka_example/chat_service/pkg/db"
	"github.com/kafka_example/chat_service/pkg/logger"
	"github.com/kafka_example/chat_service/storage"
)

func main() {
	cfg := config.Load()
	log := logger.New("debug", "chatapp")
	connDb, err := db.ConnectToDb(cfg)
	if err != nil {
		log.Error("error connection postgres: ", logger.Error(err))
	}
	strg := storage.NewStoragePg(connDb)
	strg = strg

}
