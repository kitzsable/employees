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

// @Summary Создание нового сотрудника.
// @Security AuthorizationKey
// @Tags employee
// @Description Создание нового сотрудника.
// @ID createEmployee
// @Accept json
// @Produce json
// @Param input body models.Employee true "Данные нового сотрудника."
// @Success 200 {object} transport.EntityCreatingResponse
// @Failure 400,401,500 {object} transport.ErrorResponse
// @Router /api/employee/createEmployee [post]
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

	context.JSON(http.StatusOK, transport.EntityCreatingResponse{Id: id})
}

// @Summary Получение данных сотрудника.
// @Security AuthorizationKey
// @Tags employee
// @Description Получение данных сотрудника.
// @ID getEmployee
// @Produce json
// @Param id path integer true "Идентификатор сотрудника."
// @Success 200 {object} models.Employee
// @Failure 400,401,500 {object} transport.ErrorResponse
// @Router /api/employee/getEmployee/{id} [get]
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

// @Summary Установка отдела для сотрудника.
// @Security AuthorizationKey
// @Tags employee
// @Description Установка отдела для сотрудника.
// @ID setDepartment
// @Accept json
// @Produce json
// @Param input body transport.SetDepartmentQuery true "Данные установки отдела пользователю."
// @Success 200 {object} transport.MessageResponse "Сообщение об успешной установке."
// @Failure 400,401,500 {object} transport.ErrorResponse
// @Router /api/employee/setDepartment [put]
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

	context.JSON(http.StatusOK, transport.MessageResponse{Message: "department set"})
}

// @Summary Получение списка всех сотрудников.
// @Security AuthorizationKey
// @Tags employee
// @Description Получение списка всех сотрудников.
// @ID getAllEmployees
// @Produce json
// @Success 200 {object} []models.Employee
// @Failure 400,401,500 {object} transport.ErrorResponse
// @Router /api/employee/getAllEmployees [get]
func (handler EmployeeHandler) GetAllEmployees(context *gin.Context) {
	errorMessage := "failed to get all employees"

	employees, err := handler.employeeRepository.GetAll()
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}

	context.JSON(http.StatusOK, employees)
}

// @Summary Обновление данных сотрудника.
// @Security AuthorizationKey
// @Tags employee
// @Description Обновление данных сотрудника.
// @ID updateEmployee
// @Accept json
// @Produce json
// @Param id path integer true "Идентификатор сотрудника."
// @Param input body models.Employee true "Новые данные сотрудника."
// @Success 200 {object} transport.MessageResponse "Сообщение об успешном обновлении."
// @Failure 400,401,500 {object} transport.ErrorResponse
// @Router /api/employee/updateEmployee/{id} [put]
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

	context.JSON(http.StatusOK, transport.MessageResponse{Message: "employee updated"})
}

// @Summary Удаление сотрудника.
// @Security AuthorizationKey
// @Tags employee
// @Description Удаление сотрудника.
// @ID deleteEmployee
// @Produce plain
// @Param id path integer true "Идентификатор сотрудника."
// @Success 200 {object} transport.MessageResponse "Сообщение об успешном удалении."
// @Failure 400,401,500 {object} transport.ErrorResponse
// @Router /api/employee/deleteEmployee/{id} [delete]
func (handler EmployeeHandler) DeleteEmployee(context *gin.Context) {
	errorMessage := "failed to delete employee"

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

	context.JSON(http.StatusOK, transport.MessageResponse{Message: "employee deleted"})
}
