package services

import (
	"context"

	pbch "github.com/kafka_example/chat_service/genproto/chat_service"
	"github.com/kafka_example/chat_service/pkg/logger"
	"google.golang.org/grpc/codes"
)

func (s *ChatService) AddChat(ctx context.Context, req *pbch.ChatReq) (*pbch.ChatRes, error) {
	res, err := s.storage.ChatApp().AddChat(req)
	if err != nil {
		s.log.Info(codes.Internal.String(), logger.Any("error while adding chat: ", logger.Error(err)))
		return &pbch.ChatRes{}, err
	}
	return res, nil
}
