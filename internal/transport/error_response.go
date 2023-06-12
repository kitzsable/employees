package transport

// Струтура ответа на запрос при возникновении ошибки во время его обработки.
type ErrorResponse struct {
	// Сообщение с информацией об ошибке.
	Message string `json:"message"`
}
