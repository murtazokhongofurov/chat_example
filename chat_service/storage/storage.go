package storage

import (
	"database/sql"

	"github.com/kafka_example/chat_service/storage/postgres"
	"github.com/kafka_example/chat_service/storage/repo"
)

type StorageI interface {
	ChatApp() repo.ChatService
	Message() repo.MessageService
}

type StoragePg struct {
	Db          *sql.DB
	chatrepo    repo.ChatService
	messagerepo repo.MessageService
}

func NewStoragePg(db *sql.DB) *StoragePg {
	return &StoragePg{
		Db:          db,
		chatrepo:    postgres.NewStorage(db),
		messagerepo: postgres.NewStorage(db),
	}
}

func (s StoragePg) ChatApp() repo.ChatService {
	return s.chatrepo
}

func (s StoragePg) Message() repo.MessageService {
	return s.messagerepo
}
