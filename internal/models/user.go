package models

// Пользователь.
type User struct {
	// Идентификатор пользователя.
	Id int64 `json:"-"`

	// Имя поьзователя.
	Username string `json:"username"`

	// Пароль пользователя.
	Password string `json:"password"`

	// Пароль пользователя в хэшированном виде.
	PasswordHash string `json:"-"`
}
