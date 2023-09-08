package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kafka_example/websocket_service/config"
	"github.com/kafka_example/websocket_service/kafka"
	"github.com/kafka_example/websocket_service/pkg/logger"
	"github.com/kafka_example/websocket_service/websocket"
)

func main() {
	r := gin.Default()
	cfg := config.Load()
	log := logger.New("debug", "chatapp")

	kafka, close, err := kafka.NewKafka(cfg, log)
	if err != nil {
		log.Info("error kafka: ", logger.Error(err))
	}
	defer close()
	websocket.Run(r, &kafka)
	if err := r.Run(cfg.SocketPort); err != nil {
		log.Fatal("Error listen", logger.String("port: ", cfg.SocketPort))
	}
}
