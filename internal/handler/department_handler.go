package handler

import (
	"employees/internal/domain"
	"employees/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Обработчик запросов, связанных с отделами.
type DepartmentHandler struct {
	departmentRepository domain.DepartmentRepository
}

// Создание обработчика запросов, связанных с отделами.
func NewDepartmentHandler(departmentRepository domain.DepartmentRepository) DepartmentHandler {
	return DepartmentHandler{
		departmentRepository: departmentRepository,
	}
}

// Создание нового отдела.
func (handler DepartmentHandler) CreateDepartment(context *gin.Context) {
	errorMessage := "failed to create department"

	var department models.Department

	if err := context.BindJSON(&department); err != nil {
		HandleError(context, errorMessage, ErrRequestBodyIncorrectFormat)
		return
	}

	id, err := handler.departmentRepository.Create(&department)
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// Получение отдела.
func (handler DepartmentHandler) GetDepartment(context *gin.Context) {
	errorMessage := "failed to get department"

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		HandleError(context, errorMessage, ErrQueryParamsIncorrectFormat)
		return
	}

	department, err := handler.departmentRepository.Get(int64(id))
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}

	context.JSON(http.StatusOK, department)
}

// Получение списка всех отделов.
func (handler DepartmentHandler) GetAllDepartments(context *gin.Context) {
	errorMessage := "failed to get departments"

	departments, err := handler.departmentRepository.GetAll()
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}

	context.JSON(http.StatusOK, departments)
}

// Обновление данных отдела.
func (handler DepartmentHandler) UpdateDepartment(context *gin.Context) {
	errorMessage := "failed to update department"

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		HandleError(context, errorMessage, ErrQueryParamsIncorrectFormat)
		return
	}

	var department models.Department

	if err := context.BindJSON(&department); err != nil {
		HandleError(context, errorMessage, ErrRequestBodyIncorrectFormat)
		return
	}

	err = handler.departmentRepository.Update(int64(id), department)
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// Удаление отдела.
func (handler DepartmentHandler) DeleteDepartment(context *gin.Context) {
	errorMessage := "failed to delete department"

	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		HandleError(context, errorMessage, ErrQueryParamsIncorrectFormat)
		return
	}

	err = handler.departmentRepository.Delete(int64(id))
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}
	context.String(http.StatusOK, "department deleted")
}
