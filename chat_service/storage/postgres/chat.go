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


