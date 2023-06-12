package handler

import "errors"

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
