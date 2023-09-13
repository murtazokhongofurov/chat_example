package services

import (
	"context"

	pbm "github.com/kafka_example/chat_service/genproto/message"
)

func (s MessageService) AddMessage(ctx context.Context, req *pbm.MessageReq) (*pbm.MessageRes, error) {
	res, err := s.strg.Message().AddMessage(req)
	if err != nil {
		return &pbm.MessageRes{}, err
	}

	return res, nil
}

func (s MessageService) FindMessage(ctx context.Context, req *pbm.MessageId) (*pbm.MessageRes, error) {
	res, err := s.strg.Message().FindMessage(req)
	if err != nil {
		return &pbm.MessageRes{}, err
	}
	return res, nil
}
