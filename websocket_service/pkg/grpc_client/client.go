package grpcclient

import (
	"fmt"

	"github.com/kafka_example/websocket_service/config"
	pbch "github.com/kafka_example/websocket_service/genproto/chat_service"
	pbm "github.com/kafka_example/websocket_service/genproto/message"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientI interface {
	ChatService() pbch.ChatServiceClient
	MessageService() pbm.MessageServiceClient
}

type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (GrpcClientI, error) {
	connChatService, err := grpc.Dial("localhost:9111", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error dial chat service: host: localhost and port: 9111, err: %v", err)
	}
	connMessageService, err := grpc.Dial("localhost:9111", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error dial chat service: host: localhost and port: 9111, err:%v", err)
	}
	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"chat_service":    pbch.NewChatServiceClient(connChatService),
			"message_service": pbm.NewMessageServiceClient(connMessageService),
		},
	}, nil
}

func (g *GrpcClient) ChatService() pbch.ChatServiceClient {
	return g.connections["chat_service"].(pbch.ChatServiceClient)
}

func (g *GrpcClient) MessageService() pbm.MessageServiceClient {
	return g.connections["message_service"].(pbm.MessageServiceClient)
}
