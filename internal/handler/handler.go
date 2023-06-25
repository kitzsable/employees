package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	_ "employees/docs"
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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/signUp", handler.AuthHandler.SignUp)
		auth.POST("/signIn", handler.AuthHandler.SignIn)
	}

	api := router.Group("/api", handler.AuthMiddleware.Autorize)
	{
		employee := api.Group("/employee")
		{
			employee.POST("/createEmployee", handler.EmployeeHandler.CreateEmployee)
			employee.PUT("/setDepartment", handler.EmployeeHandler.SetEmployeeDepartment)
			employee.GET("/getEmployee/:id", handler.EmployeeHandler.GetEmployee)
			employee.GET("/getAllEmployees", handler.EmployeeHandler.GetAllEmployees)
			employee.PUT("/updateEmployee/:id", handler.EmployeeHandler.UpdateEmployee)
			employee.DELETE("/deleteEmployee/:id", handler.EmployeeHandler.DeleteEmployee)
		}

		department := api.Group("/department")
		{
			department.POST("/createDepartment", handler.DepartmentHandler.CreateDepartment)
			department.GET("/getDepartment/:id", handler.DepartmentHandler.GetDepartment)
			department.GET("/getAllDepartments", handler.DepartmentHandler.GetAllDepartments)
			department.PUT("/updateDepartment/:id", handler.DepartmentHandler.UpdateDepartment)
			department.DELETE("/deleteDepartment/:id", handler.DepartmentHandler.DeleteDepartment)
		}
	}

	return router
}
