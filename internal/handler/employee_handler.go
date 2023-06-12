package handler

import (
	"employees/internal/domain"
	"employees/internal/models"
	"employees/internal/transport"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Обработчик запросов, связанных с сотрудниками.
type EmployeeHandler struct {
	employeeRepository domain.EmployeeRepository
}

// Создание обработчика запросов, связанных с сотрудниками.
func NewEmployeeHandler(employeeRepository domain.EmployeeRepository) EmployeeHandler {
	return EmployeeHandler{
		employeeRepository: employeeRepository,
	}
}

// Создание нового сотрудника.
func (handler EmployeeHandler) CreateEmployee(context *gin.Context) {
	errorMessage := "failed to create employee"

	var employee models.Employee

	if err := context.BindJSON(&employee); err != nil {
		HandleError(context, errorMessage, ErrRequestBodyIncorrectFormat)
		return
	}

	id, err := handler.employeeRepository.Create(&employee)
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// Получение данных сотрудника.
func (handler EmployeeHandler) GetEmployee(context *gin.Context) {
	errorMessage := "failed to get employee"

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		HandleError(context, errorMessage, ErrQueryParamsIncorrectFormat)
		return
	}

	employee, err := handler.employeeRepository.Get(int64(id))
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}

	context.JSON(http.StatusOK, employee)
}

// Установка отдела для сотрудника.
func (handler EmployeeHandler) SetEmployeeDepartment(context *gin.Context) {
	errorMessage := "failed to set department"

	var query transport.SetDepartmentQuery

	if err := context.BindJSON(&query); err != nil {
		HandleError(context, errorMessage, ErrRequestBodyIncorrectFormat)
		return
	}

	err := handler.employeeRepository.SetDepartment(query.EmployeeId, query.DepartmentId)
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}

	context.String(http.StatusOK, "department set")
}

// Получение списка всех сотрудников.
func (handler EmployeeHandler) GetAllEmployees(context *gin.Context) {
	errorMessage := "failed to get all employees"

	employees, err := handler.employeeRepository.GetAll()
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}

	context.JSON(http.StatusOK, employees)
}

// Обновление данных сотрудника.
func (handler EmployeeHandler) UpdateEmployee(context *gin.Context) {
	errorMessage := "failed to update employee"

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		HandleError(context, errorMessage, ErrQueryParamsIncorrectFormat)
		return
	}

	var employee models.Employee

	if err := context.BindJSON(&employee); err != nil {
		HandleError(context, errorMessage, ErrRequestBodyIncorrectFormat)
		return
	}

	err = handler.employeeRepository.Update(int64(id), employee)
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}

	context.String(http.StatusOK, "employee updated")
}

// Удаление сотрудника.
func (handler EmployeeHandler) DeleteEmployee(context *gin.Context) {
	errorMessage := "failed to delete department"

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		HandleError(context, errorMessage, ErrQueryParamsIncorrectFormat)
		return
	}

	err = handler.employeeRepository.Delete(int64(id))
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}

	context.String(http.StatusOK, "employee deleted")
}
