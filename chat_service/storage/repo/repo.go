package repo

import (
	pbch "github.com/kafka_example/chat_service/genproto/chat_service"
	pbm "github.com/kafka_example/chat_service/genproto/message"
)

type ChatService interface {
	AddChat(*pbch.ChatReq) (*pbch.ChatRes, error)
	AddPrivateChat(*pbch.PrivateChatReq) (*pbch.PrivateChatRes, error)

	AddUser(*pbch.UserReq) (*pbch.UserRes, error)
	FindUser(*pbch.UserId) (*pbch.UserRes, error)
	RemoveUser(*pbch.UserId) error
	Update(*pbch.UserRes) (*pbch.UserRes, error)
}

type MessageService interface {
	AddMessage(*pbm.MessageReq) (*pbm.MessageRes, error)
	FindMessage(*pbm.MessageId) (*pbm.MessageRes, error)
}
