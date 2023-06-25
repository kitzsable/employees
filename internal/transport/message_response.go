package transport

// Струтура ответа на запрос с сообщением.
type MessageResponse struct {
	// Сообщение.
	Message string `json:"message"`
}
