package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// Возможные ошибки при обработке запросов.
var (
	// Ошибка некорректного формата данных в теле запроса.
	ErrRequestBodyIncorrectFormat = errors.New("request body has incorrect format")

	// Ошибка некорректного формата данных в параметрах запроса.
	ErrQueryParamsIncorrectFormat = errors.New("request params has incorrect format")
)

// Обработчик запросов.
type Handler struct {
	// Обработчик авторизации пользователя.
	AuthMiddleware AuthMiddleware

	// Обработчик запросов, связанных с регистрацией / аутентификацией.
	AuthHandler AuthHandler

	// Обработчик запросов, связанных с отделами.
	DepartmentHandler DepartmentHandler

	// Обработчик запросов, связанных с сотрудниками.
	EmployeeHandler EmployeeHandler
}

// Создание нового обработчика запросов.
func NewHandler(
	authMiddleware AuthMiddleware,
	authHandler AuthHandler,
	departmentHandler DepartmentHandler,
	employeeHandler EmployeeHandler,
) Handler {
	return Handler{
		AuthMiddleware:    authMiddleware,
		AuthHandler:       authHandler,
		DepartmentHandler: departmentHandler,
		EmployeeHandler:   employeeHandler,
	}
}

func (handler *Handler) ConfigureRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signUp", handler.AuthHandler.SignUp)
		auth.POST("/signIn", handler.AuthHandler.SignIn)
	}

	api := router.Group("/api", handler.AuthMiddleware.Autorize)
	{
		api.POST("/employee", handler.EmployeeHandler.CreateEmployee)
		api.POST("/setDepartment", handler.EmployeeHandler.SetEmployeeDepartment)
		api.GET("/employee/:id", handler.EmployeeHandler.GetEmployee)
		api.GET("/employees", handler.EmployeeHandler.GetAllEmployees)
		api.PUT("/employee/:id", handler.EmployeeHandler.UpdateEmployee)
		api.DELETE("/employee/:id", handler.EmployeeHandler.DeleteEmployee)

		api.POST("/department", handler.DepartmentHandler.CreateDepartment)
		api.GET("/department/:id", handler.DepartmentHandler.GetDepartment)
		api.GET("/departments", handler.DepartmentHandler.GetAllDepartments)
		api.PUT("/department/:id", handler.DepartmentHandler.UpdateDepartment)
		api.DELETE("/department/:id", handler.DepartmentHandler.DeleteDepartment)
	}

	return router
}
