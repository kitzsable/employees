package handler

import (
	"employees/internal/domain"
	"employees/internal/models"
	"employees/internal/transport"
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

// @Summary Создание нового отдела.
// @Security AuthorizationKey
// @Tags department
// @Description Создание нового отдела.
// @ID createDepartment
// @Accept json
// @Produce json
// @Param input body models.Department true "Данные нового отдела."
// @Success 200 {object} transport.EntityCreatingResponse
// @Failure 400,401,500 {object} transport.ErrorResponse
// @Router /api/department/createDepartment [post]
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

	context.JSON(http.StatusOK, transport.EntityCreatingResponse{Id: id})
}

// @Summary Получение отдела.
// @Security AuthorizationKey
// @Tags department
// @Description Получение информации об отделе.
// @ID getDepartment
// @Produce json
// @Param id path integer true "Идентификатор отдела."
// @Success 200 {object} models.Department
// @Failure 400,401,500 {object} transport.ErrorResponse
// @Router /api/department/getDepartment/{id} [get]
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

// @Summary Получение списка всех отделов.
// @Security AuthorizationKey
// @Tags department
// @Description Получение списка всех отделов.
// @ID getAllDepartments
// @Produce  json
// @Success 200 {object} []models.Department
// @Failure 400,401,500 {object} transport.ErrorResponse
// @Router /api/department/getAllDepartments [get]
func (handler DepartmentHandler) GetAllDepartments(context *gin.Context) {
	errorMessage := "failed to get departments"

	departments, err := handler.departmentRepository.GetAll()
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}

	context.JSON(http.StatusOK, departments)
}

// @Summary Обновление данных отдела.
// @Security AuthorizationKey
// @Tags department
// @Description Обновление данных отдела.
// @ID updateDepartment
// @Accept json
// @Produce json
// @Param id path integer true "Идентификатор отдела."
// @Param input body models.Department true "Новые данные отдела."
// @Success 200 {object} transport.MessageResponse "Сообщение об успешном обновлении."
// @Failure 400,401,500 {object} transport.ErrorResponse
// @Router /api/department/updateDepartment/{id} [put]
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

	context.JSON(http.StatusOK, transport.MessageResponse{Message: "department updated"})
}

// @Summary Удаление отдела.
// @Security AuthorizationKey
// @Tags department
// @Description Удаление отдела.
// @ID deleteDepartment
// @Produce  plain
// @Param id path integer true "Идентификатор отдела."
// @Success 200 {object} transport.MessageResponse "Сообщение об успешном удалении."
// @Failure 400,401,500 {object} transport.ErrorResponse
// @Router /api/department/deleteDepartment/{id} [delete]
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

	context.JSON(http.StatusOK, transport.MessageResponse{Message: "department deleted"})
}
