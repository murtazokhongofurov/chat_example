package storage

import (
	"database/sql"
	"github.com/kafka_example/chat_service/storage/postgres"
	"github.com/kafka_example/chat_service/storage/repo"
)

type StorageI interface {
	ChatApp() repo.ChatService
}

type StoragePg struct {
	Db       *sql.DB
	chatrepo repo.ChatService
}

func NewStoragePg(db *sql.DB) *StoragePg {
	return &StoragePg{
		Db:       db,
		chatrepo: postgres.NewStorage(db),
	}
}

func (s StoragePg) ChatApp() repo.ChatService {
	return s.chatrepo
}
