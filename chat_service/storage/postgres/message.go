package postgres

import pbm "github.com/kafka_example/chat_service/genproto/message"

func (r storagePg) AddMessage(req *pbm.MessageReq) (*pbm.MessageRes, error) {
	res := pbm.MessageRes{}
	query := `
	INSERT INTO 
		messages(user_id, chat_id, text) 
	VALUES 
		($1, $2, $3) RETURNING id, user_id, chat_id, text, created_at, updated_at`
	err := r.db.QueryRow(query, req.UserId, req.ChatId, req.MessageText).
		Scan(&res.Id, &res.UserId, &res.ChatId, &res.MessageText, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return &pbm.MessageRes{}, err
	}
	return &res, nil
}

func (r storagePg) FindMessage(req *pbm.MessageId) (*pbm.MessageRes, error) {
	res := pbm.MessageRes{}
	err := r.db.QueryRow(`
	SELECT 
		id, user_id, chat_id, message_text, created_at, updated_at 
	FROM 
		messages WHERE id=$1`, req.MessageId).
		Scan(&res.Id, &res.UserId, &res.ChatId, &res.MessageText, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return &pbm.MessageRes{}, err
	}

	return &res, nil
}
