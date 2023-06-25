package transport

// Струтура ответа на запрос аутентификации.
type SignInResponse struct {
	// Токен авторизации.
	Token string `json:"token"`
}
