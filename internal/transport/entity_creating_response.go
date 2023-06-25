package transport

// Струтура ответа на запрос создания сущности в БД.
type EntityCreatingResponse struct {
	// Идентификатор сущности.
	Id int64 `json:"id"`
}
