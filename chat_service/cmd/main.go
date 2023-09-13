package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/kafka_example/chat_service/config"
	pbch "github.com/kafka_example/chat_service/genproto/chat_service"
	"github.com/kafka_example/chat_service/genproto/message"
	"github.com/kafka_example/chat_service/kafka"
	"github.com/kafka_example/chat_service/pkg/db"
	"github.com/kafka_example/chat_service/pkg/logger"
	"github.com/kafka_example/chat_service/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	cfg := config.Load()
	log := logger.New("debug", "chatapp")
	connDb, err := db.ConnectToDb(cfg)
	if err != nil {
		log.Error("error connection postgres: ", logger.Error(err))
	}
	lis, err := net.Listen("tcp", ":"+cfg.HttpPort)
	if err != nil {
		log.Error(err.Error())
	}
	defer lis.Close()

	_ = services.NewChatService(connDb, log)
	_ = services.NewMessageService(connDb, log)
	s := grpc.NewServer()

	reflection.Register(s)

	pbch.RegisterChatServiceServer(s, pbch.UnimplementedChatServiceServer{})
	message.RegisterMessageServiceServer(s, message.UnimplementedMessageServiceServer{})
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
	fmt.Println("Server is listening on port: ", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: ", logger.Error(err))
	}
}
