package models

type MessageReq struct {
	ChatID int    `json:"chat_id"`
	UserId int    `json:"user_id"`
	Text   string `json:"text"`
}

type MessageRes struct {
	Id        int    `json:"id"`
	ChatID    int    `json:"chat_id"`
	UserId    int    `json:"user_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Messages struct {
	Id        int    `json:"id"`
	ChatID    int    `json:"chat_id"`
	UserId    int    `json:"user_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	User      User   `json:"user"`
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Image     string `json:"image"`
}
