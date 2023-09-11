package main

import (
	"fmt"

	"github.com/kafka_example/chat_service/config"
	"github.com/kafka_example/chat_service/kafka/consumer"
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
	kafka, err := consumer.NewKafkaConsumer(cfg, log, strg)
	if err != nil {
		fmt.Println("Error creating consumer: ", logger.Error(err))
		return
	}else{
		fmt.Println("Connected to Kafka sucessfully")
	}
	kafka.ConsumeMessages()

}
