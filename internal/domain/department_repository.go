package domain

import "employees/internal/models"

// Интерфейс репозитория для работы с отделами в базе данных.
type DepartmentRepository interface {
	// Создание нового отдела.
	Create(department *models.Department) (int64, error)

	// Получение отдела по его идентификатору.
	Get(id int64) (models.Department, error)

	// Получение списка всех отделов.
	GetAll() ([]models.Department, error)

	// Обновление данных отдела по его идентификатору.
	Update(id int64, department models.Department) error

	// Удаление отдела по его идентификатору.
	Delete(id int64) error
}
