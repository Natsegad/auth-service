package auth

// Ответ от сервера в виде json
type AuthUserRes struct {
	Token string `json:"jwt"`
	Id    uint64 `json:"id"`
}

// Запрос для авторизации пользователя
type AuthUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Id       uint64 `json:"id"`
}
