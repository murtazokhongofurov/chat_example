package postgres

import pbch "github.com/kafka_example/chat_service/genproto/chat_service"

func (r storagePg) AddChat(req *pbch.ChatReq) (*pbch.ChatRes, error) {
	res := pbch.ChatRes{}
	query := `
	INSERT INTO chats(chat_type) VALUES($1) RETURNING id, chat_type`
	err := r.db.QueryRow(query, req.ChatType).Scan(&res.Id, &res.ChatType)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r storagePg) AddPrivateChat(req *pbch.PrivateChatReq) (*pbch.PrivateChatRes, error) {
	res := pbch.PrivateChatRes{}
	err := r.db.QueryRow(`INSERT INTO user_chat(user_id, chat_id) VALUES($1, $2) 
	RETURNING id, user_id, chat_id, created_at`, req.UserId, req.ChatId).
		Scan(&res.Id, &res.UserId, &res.ChatId, &res.CreatedAt)
	if err != nil {
		return &pbch.PrivateChatRes{}, err
	}
	return &res, nil
}
