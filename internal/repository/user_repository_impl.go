package repository

import (
	"database/sql"
	"employees/internal/models"
	"fmt"
)

// Реализация репозитория для работы с пользователями в базе данных.
type UserRepositoryImpl struct {
	database *sql.DB
}

// Функция создания нового репозитория для работы с пользователями в базе данных.
func NewUserRepositoryImpl(database *sql.DB) UserRepositoryImpl {
	return UserRepositoryImpl{
		database: database,
	}
}

func (repository UserRepositoryImpl) Create(user *models.User) (int64, error) {
	var id int64 = 0
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) values ($1, $2) RETURNING id", usersTable)

	row := repository.database.QueryRow(query, user.Username, user.PasswordHash)
	if err := row.Scan(&id); err != nil {
		return id, fmt.Errorf("cannot insert user: %v", err)
	}

	return id, nil
}

func (repository UserRepositoryImpl) Get(username string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id, username, password_hash FROM %s WHERE username = $1", usersTable)

	row := repository.database.QueryRow(query, username)
	if err := row.Scan(&user.Id, &user.Username, &user.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return user, ErrRecordNotFound
		}
		return user, fmt.Errorf("cannot read user from database: %v", err)
	}
	return user, nil
}
