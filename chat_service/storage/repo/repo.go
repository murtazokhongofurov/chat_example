package repo

import "github.com/kafka_example/chat_service/storage/models"



type (
	ChatService interface {
		SavePerson(person models.Message) error 
	}
)
