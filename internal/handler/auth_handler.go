package handler

import (
	"employees/internal/domain"
	"employees/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Обработчик запросов, связанных с регистрацией / аутентификацией.
type AuthHandler struct {
	authService domain.AuthService
}

// Создание нового обработчика запросов, связанных с регистрацией / аутентификацией.
func NewAuthHandler(authService domain.AuthService) AuthHandler {
	return AuthHandler{
		authService: authService,
	}
}

// Регистрация пользователя.
func (handler AuthHandler) SignUp(context *gin.Context) {
	errorMessage := "failed to sign up user"

	var user models.User

	if err := context.BindJSON(&user); err != nil {
		HandleError(context, errorMessage, ErrRequestBodyIncorrectFormat)
		return
	}

	id, err := handler.authService.CreateUser(&user)
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// Аутентификация пользователя (получение токена).
func (handler AuthHandler) SignIn(context *gin.Context) {
	errorMessage := "failed to sign in user"

	var user models.User

	if err := context.BindJSON(&user); err != nil {
		HandleError(context, errorMessage, ErrRequestBodyIncorrectFormat)
		return
	}

	signingToken, err := handler.authService.SignInUser(&user)
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}

	context.JSON(http.StatusOK, map[string]interface{}{"token": signingToken})
}
