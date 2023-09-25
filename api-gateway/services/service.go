package services

import (
	"fmt"
	"github.com/kafka_example/api-gateway/config"
	"github.com/kafka_example/api-gateway/genproto/chat_service"
	"github.com/kafka_example/api-gateway/genproto/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	ChatService() chat_service.ChatServiceClient
	MessageService() message.MessageServiceClient
}

type serviceManager struct {
	chatService    chat_service.ChatServiceClient
	messageService message.MessageServiceClient
}

func NewServiceManager(cfg config.Config) (ServiceManagerI, error) {
	chatService, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.ChatServiceHost, cfg.CHatServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error dial chat service: host: localhost and port: 9111, err: %v", err)
	}

	messageService, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.ChatServiceHost, cfg.CHatServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error dial chat service: host: localhost and port: 9111, err: %v", err)
	}

	return &serviceManager{
		chatService:    chat_service.NewChatServiceClient(chatService),
		messageService: message.NewMessageServiceClient(messageService),
	}, nil

}

func (s *serviceManager) ChatService() chat_service.ChatServiceClient {
	return s.chatService
}

func (s *serviceManager) MessageService() message.MessageServiceClient {
	return s.messageService
}
