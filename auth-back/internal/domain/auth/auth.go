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
	Token    string `json:"jwt"`
}

type UsersJson struct {
	Users map[uint64]AuthUserReq `json:"users"`
}

// Для проверки есть ли пользователь мы должны получить GET запрос с Id нужного пользователя
type ValidUserJsonReq struct {
	Id uint64 `json:"user_id"`
}
