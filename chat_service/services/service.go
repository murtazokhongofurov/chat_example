package services

import (
	"database/sql"

	"github.com/kafka_example/chat_service/genproto/chat_service"
	"github.com/kafka_example/chat_service/genproto/message"
	"github.com/kafka_example/chat_service/pkg/logger"
	"github.com/kafka_example/chat_service/storage"
)

type ChatService struct {
	chat_service.UnimplementedChatServiceServer
	storage storage.StorageI
	log     logger.Logger
}

func NewChatService(db *sql.DB, log logger.Logger) *ChatService {
	return &ChatService{
		storage: storage.NewStoragePg(db),
		log:     log,
	}
}

type MessageService struct {
	message.UnimplementedMessageServiceServer
	strg storage.StorageI
	log  logger.Logger
}

func NewMessageService(db *sql.DB, log logger.Logger) *MessageService {
	return &MessageService{
		strg: storage.NewStoragePg(db),
		log:  log,
	}
}
