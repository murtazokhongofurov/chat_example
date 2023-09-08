package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/kafka_example/chat_service/config"
	"github.com/kafka_example/chat_service/kafka"
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
	_ = storage.NewStoragePg(connDb)
	kafkaConn, close, err := kafka.NewKafkaReader(cfg, connDb)
	if err != nil {
		log.Error("Error connection to kafka: ", logger.Error(err))
	}
	defer close()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		kafkaConn.Reads().Consume()
		wg.Done()
	}()
	if err := http.ListenAndServe(":9111", nil); err != nil {
		fmt.Println("error listen :9111", err)
	}
}
