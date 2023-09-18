package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kafka_example/websocket_service/config"
	grpcclient "github.com/kafka_example/websocket_service/pkg/grpc_client"
	"github.com/kafka_example/websocket_service/pkg/logger"
	"github.com/kafka_example/websocket_service/websocket"
)

func main() {
	r := gin.Default()
	cfg := config.Load()
	log := logger.New("debug", "chatapp")

	grpcConn, err := grpcclient.New(cfg)
	if err != nil {
		log.Error(err.Error())
	}
/*	kafka, close, err := kafka.NewKafka(cfg, log)
	if err != nil {
		log.Info("error kafka: ", logger.Error(err))
	}
	defer close()
*/
	fmt.Println("Port:", cfg.SocketPort)
	websocket.Run(r, grpcConn)
	if err := r.Run(":8081"); err != nil {
		log.Fatal("Error listen", logger.String("port: ", cfg.SocketPort))
	}
}
