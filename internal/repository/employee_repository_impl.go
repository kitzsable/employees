package repository

import (
	"database/sql"
	"employees/internal/models"
	"fmt"
)

// Реализация репозитория для работы с сотрудниками в базе данных.
type EmployeeRepositoryImpl struct {
	database *sql.DB
}

// Функция создания нового репозитория для работы с сотрудниками в базе данных.
func NewEmployeeRepositoryImpl(database *sql.DB) EmployeeRepositoryImpl {
	return EmployeeRepositoryImpl{
		database: database,
	}
}

func (repository EmployeeRepositoryImpl) Create(employee *models.Employee) (int64, error) {
	var id int64 = 0
	query := fmt.Sprintf("INSERT INTO %s (name, sex, age, salary) VALUES ($1, $2, $3, $4) RETURNING id", employeesTable)

	row := repository.database.QueryRow(query, employee.Name, employee.Sex, employee.Age, employee.Salary)
	if err := row.Scan(&id); err != nil {
		return id, fmt.Errorf("cannot insert employee to database: %v", err)
	}

	return id, nil
}

func (repository EmployeeRepositoryImpl) Get(id int64) (models.Employee, error) {
	var employee models.Employee
	query := fmt.Sprintf(`SELECT e.id, e.name, e.sex, e.age, e.salary, d.id, d.name FROM %s e 
	JOIN %s d ON e.department_id = d.id 
	WHERE e.id = $1`, employeesTable, depatmentsTable)

	row := repository.database.QueryRow(query, id)
	if err := row.Scan(&employee.Id, &employee.Name, &employee.Sex, &employee.Age, &employee.Salary, &employee.Department.Id, &employee.Department.Name); err != nil {
		if err == sql.ErrNoRows {
			return employee, ErrRecordNotFound
		}
		return employee, fmt.Errorf("cannot read employee with id = %d from database: %v", id, err)
	}
	return employee, nil
}

func (repository EmployeeRepositoryImpl) SetDepartment(employeeId int64, departmentId int64) error {
	var exist bool
	departmentExistQuery := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE id = $1)", depatmentsTable)
	setDepartmentQuery := fmt.Sprintf("UPDATE %s SET department_id = $1 WHERE id = $2", employeesTable)

	row := repository.database.QueryRow(departmentExistQuery, departmentId)
	if err := row.Scan(&exist); err != nil {
		return fmt.Errorf("cannot check department exists in database: %v", err)
	}

	if !exist {
		return ErrRecordNotFound
	}

	_, err := repository.database.Exec(setDepartmentQuery, departmentId, employeeId)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrRecordNotFound
		}
		return fmt.Errorf("cannot update employee with id = %d in database: %v", employeeId, err)
	}

	return nil
}

func (repository EmployeeRepositoryImpl) GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	query := fmt.Sprintf(`SELECT e.id, e.name, e.sex, e.age, e.salary, d.id, d.name FROM %s e 
	LEFT JOIN %s d ON e.department_id = d.id`, employeesTable, depatmentsTable)

	rows, err := repository.database.Query(query)
	if err != nil {
		return employees, fmt.Errorf("cannot read employees from database: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var employee models.Employee
		var departmentId sql.NullInt64
		var departmentName sql.NullString

		if err := rows.Scan(&employee.Id, &employee.Name, &employee.Sex, &employee.Age, &employee.Salary, &departmentId, &departmentName); err != nil {
			return employees, fmt.Errorf("cannot read employees from database: %v", err)
		}

		if departmentId.Valid && departmentName.Valid {
			employee.Department.Id = departmentId.Int64
			employee.Department.Name = departmentName.String
		}

		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		return employees, fmt.Errorf("cannot read employees from database: %v", err)
	}
	return employees, nil
}

func (repository EmployeeRepositoryImpl) Update(id int64, employee models.Employee) error {
	exist, err := repository.checkRecordExist(id)

	if err != nil {
		return err
	}

	if !exist {
		return ErrRecordNotFound
	}

	query := fmt.Sprintf("UPDATE %s SET name = $1, sex = $2, age = $3, salary = $4 WHERE id = $5", employeesTable)

	_, err = repository.database.Exec(query, employee.Name, employee.Sex, employee.Age, employee.Salary, id)
	if err != nil {
		return fmt.Errorf("cannot update employee with id = %d in database: %v", id, err)
	}

	return nil
}

func (repository EmployeeRepositoryImpl) Delete(id int64) error {
	exist, err := repository.checkRecordExist(id)

	if err != nil {
		return err
	}

	if !exist {
		return ErrRecordNotFound
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", employeesTable)

	_, err = repository.database.Exec(query, id)
	if err != nil {
		return fmt.Errorf("cannot delete employee with id = %d from database: %v", id, err)
	}

	return nil
}

func (repository EmployeeRepositoryImpl) checkRecordExist(id int64) (bool, error) {
	var exist bool
	employeeExistQuery := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE id = $1)", employeesTable)

	row := repository.database.QueryRow(employeeExistQuery, id)
	if err := row.Scan(&exist); err != nil {
		return false, fmt.Errorf("cannot check employee exists in database: %v", err)
	}

	return exist, nil
}
