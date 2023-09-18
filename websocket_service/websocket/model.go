package websocket

type Broadcast struct {
	Type string `json:"type"`
	Users []string `json:"users"`
}

type Response struct {
	Message string `json:"message"`
	Type string `json:"type"`
	Users []string `json:"users"`
}

type WebscoketMessage struct {
	Type string           `json:"type"`
	Conn SocketConnection `json:"conn"`
	Chat_id int `json:"chat_id"`
	User_id int `json:"user_id"`
	Message string `json:"message"`
}
