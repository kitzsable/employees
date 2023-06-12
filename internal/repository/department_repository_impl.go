package repository

import (
	"database/sql"
	"employees/internal/models"
	"fmt"
)

// Реализация репозиторий для работы с отделами в базе данных.
type DepartmentRepositoryImpl struct {
	database *sql.DB
}

// Функция создания нового репозитория для работы с отделами в базе данных.
func NewDepartmentRepositoryImpl(database *sql.DB) DepartmentRepositoryImpl {
	return DepartmentRepositoryImpl{
		database: database,
	}
}

func (repository DepartmentRepositoryImpl) Create(department *models.Department) (int64, error) {
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", depatmentsTable)

	row := repository.database.QueryRow(query, department.Name)
	if err := row.Scan(&id); err != nil {
		return id, fmt.Errorf("cannot insert department to database: %v", err)
	}

	return id, nil
}

func (repository DepartmentRepositoryImpl) Get(id int64) (models.Department, error) {
	var department models.Department
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE id = $1", depatmentsTable)

	row := repository.database.QueryRow(query, id)
	if err := row.Scan(&department.Id, &department.Name); err != nil {
		if err == sql.ErrNoRows {
			return department, ErrRecordNotFound
		}
		return department, fmt.Errorf("cannot read department with id = %d from database: %v", id, err)
	}
	return department, nil
}

func (repository DepartmentRepositoryImpl) GetAll() ([]models.Department, error) {
	var departments []models.Department
	query := fmt.Sprintf("SELECT id, name FROM %s", depatmentsTable)

	rows, err := repository.database.Query(query)
	if err != nil {
		return departments, fmt.Errorf("cannot read departments from database: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var department models.Department
		if err := rows.Scan(&department.Id, &department.Name); err != nil {
			return departments, fmt.Errorf("cannot read departments from database: %v", err)
		}
		departments = append(departments, department)
	}
	if err := rows.Err(); err != nil {
		return departments, fmt.Errorf("cannot read departments from database: %v", err)
	}
	return departments, nil
}

func (repository DepartmentRepositoryImpl) Update(id int64, department models.Department) error {
	exist, err := repository.checkRecordExist(id)

	if err != nil {
		return err
	}

	if !exist {
		return ErrRecordNotFound
	}

	query := fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2", depatmentsTable)

	_, err = repository.database.Exec(query, department.Name, id)
	if err != nil {
		return fmt.Errorf("cannot update department with id = %d from database: %v", id, err)
	}

	return nil
}

func (repository DepartmentRepositoryImpl) Delete(id int64) error {
	exist, err := repository.checkRecordExist(id)

	if err != nil {
		return err
	}

	if !exist {
		return ErrRecordNotFound
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", depatmentsTable)

	_, err = repository.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("cannot delete department with id = %d from database: %v", id, err)
	}

	return nil
}

func (repository DepartmentRepositoryImpl) checkRecordExist(id int64) (bool, error) {
	var exist bool
	employeeExistQuery := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE id = $1)", depatmentsTable)

	row := repository.database.QueryRow(employeeExistQuery, id)
	if err := row.Scan(&exist); err != nil {
		return false, fmt.Errorf("cannot check department exists in database: %v", err)
	}

	return exist, nil
}
