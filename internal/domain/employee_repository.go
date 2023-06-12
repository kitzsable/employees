package domain

import "employees/internal/models"

// Интерфейс репозитория для работы с сотрудниками в базе данных.
type EmployeeRepository interface {
	// Создание нового сотрудника.
	Create(employee *models.Employee) (int64, error)

	// Получение сотрудника по его идентификатору.
	Get(id int64) (models.Employee, error)

	// Установка отдела для сотрудника по их идентификаторам.
	SetDepartment(employeeId int64, departmentId int64) error

	// Получение списка всех сотрудников.
	GetAll() ([]models.Employee, error)

	// Обновление данных сотрудника по его идентификатору.
	Update(id int64, employee models.Employee) error

	// Удаление сотрудника по его идентификатору.
	Delete(id int64) error
}
