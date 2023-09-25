package models

type UserRequest struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	UserName     string `json:"user_name"`
	Bio          string `json:"bio"`
	Image        string `json:"image"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	RefreshToken string `json:"refresh_token"`
}

type UserResponse struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

type UserLoginResponse struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserName    string `json:"user_name"`
	Bio         string `json:"bio"`
	Email       string `json:"email"`
	Image       string `json:"image"`
	Phone       string `json:"phone"`
	AccessToken string `json:"access_token"`
}

type Users struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	ChatId    int    `json:"chat_id"`
	ChatType  string `json:"chat_type"`
}

type CheckFieldReq struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type CheckFieldRes struct {
	Exists bool `json:"exists"`
}

type RedisSave struct {
	Code        string `json:"code"`
	SendDate    string `json:"send_date"`
	ExpiredDate string `json:"expired_date"`
}
