package handler

import (
	"employees/internal/domain"
	"employees/internal/models"
	"employees/internal/transport"
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

// @Summary Регистрация.
// @Tags auth
// @Description Регистрация пользователя в системе.
// @ID signUp
// @Accept  json
// @Produce  json
// @Param input body models.User true "Данные нового пользователя."
// @Success 200 {object} transport.EntityCreatingResponse
// @Failure 400,401,500 {object} transport.ErrorResponse
// @Router /auth/signUp [post]
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

	context.JSON(http.StatusOK, transport.EntityCreatingResponse{Id: id})
}

// @Summary Аутентификация.
// @Tags auth
// @Description Аутентификация пользователя (получение токена).
// @ID signIn
// @Accept json
// @Produce json
// @Param input body models.User true "Данные пользователя."
// @Success 200 {object} transport.SignInResponse
// @Failure 400,401,500 {object} transport.ErrorResponse
// @Router /auth/signIn [post]
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

	context.JSON(http.StatusOK, transport.SignInResponse{Token: signingToken})
}
