package domain

import "employees/internal/models"

// Интерфейс репозитория для работы с пользователями в базе данных.
type UserRepository interface {
	// Создание нового пользователя.
	Create(user *models.User) (int64, error)

	// Получение пользователя по его идентификатору.
	Get(username string) (models.User, error)
}
