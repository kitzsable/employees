package handler

import (
	"employees/internal/domain"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	UserKey             = "userId"
)

// Обработчик авторизации пользователя.
type AuthMiddleware struct {
	authService domain.AuthService
}

// Создание обработчика авторизации пользователя.
func NewAuthMiddleware(authService domain.AuthService) AuthMiddleware {
	return AuthMiddleware{
		authService: authService,
	}
}

// Проверка токена доступа пользователя.
func (middleware AuthMiddleware) Autorize(context *gin.Context) {
	errorMessage := "failed user authorize"
	header := context.GetHeader(authorizationHeader)
	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 || headerParts[0] != "Bearer" || len(headerParts[1]) == 0 {
		HandleError(context, errorMessage, domain.ErrAccessTokenIncorrect)
		return
	}

	userId, err := middleware.authService.ParseAccessToken(headerParts[1])
	if err != nil {
		HandleError(context, errorMessage, err)
		return
	}

	context.Set(UserKey, userId)
}
