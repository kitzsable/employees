package models

// Отдел.
type Department struct {
	// Идентификатор отдела.
	Id int64 `json:"id"`

	// Наименование отдела.
	Name string `json:"name"`
}
