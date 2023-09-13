package services

import (
	"context"

	pbch "github.com/kafka_example/chat_service/genproto/chat_service"
	"github.com/kafka_example/chat_service/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ChatService) AddChat(ctx context.Context, req *pbch.ChatReq) (*pbch.ChatRes, error) {
	res, err := s.storage.ChatApp().AddChat(req)
	if err != nil {
		s.log.Info(codes.Internal.String(), logger.Any("error while adding chat: ", logger.Error(err)))
		return &pbch.ChatRes{}, err
	}
	return res, nil
}

func (s *ChatService) AddUser(ctx context.Context, req *pbch.UserReq) (*pbch.UserRes, error) {
	res, err := s.storage.ChatApp().AddUser(req)
	if err != nil {
		s.log.Error("Error while insert", logger.Any("error insert user", err))
		return &pbch.UserRes{}, status.Error(codes.Internal, "something went wrong, please check user info")
	}
	return res, nil
}

func (s *ChatService) FindUser(ctx context.Context, req *pbch.UserId) (*pbch.UserRes, error) {
	res, err := s.storage.ChatApp().FindUser(req)
	if err != nil {
		s.log.Error("Error while get", logger.Any("error get user", err))
		return &pbch.UserRes{}, status.Error(codes.Internal, "something went wrong, please check chat info")
	}
	return res, nil
}

func (s *ChatService) RemoveUser(ctx context.Context, req *pbch.UserId) (*pbch.Empty, error) {
	err := s.storage.ChatApp().RemoveUser(req)
	if err != nil {
		s.log.Error("Error while delete", logger.Any("error delete user", err))
		return &pbch.Empty{}, status.Error(codes.Internal, "something went wrong, please check user info")
	}
	return &pbch.Empty{}, nil
}

func (s *ChatService) Update(ctx context.Context, req *pbch.UserRes) (*pbch.UserRes, error) {
	res, err := s.storage.ChatApp().Update(req)
	if err != nil {
		s.log.Error("Error while update", logger.Any("error update user", err))
		return &pbch.UserRes{}, status.Error(codes.Internal, "something went wrong, please check user info")
	}
	return res, nil
}

func (s *ChatService) AddPrivateChat(ctx context.Context, req *pbch.PrivateChatReq) (*pbch.PrivateChatRes, error) {
	res, err := s.storage.ChatApp().AddPrivateChat(req)
	if err != nil {
		s.log.Error("Error while update", logger.Any("error update user_chat", err))
		return &pbch.PrivateChatRes{}, status.Error(codes.Internal, "something went wrong, please check user_chat info")
	}
	return res, nil
}
