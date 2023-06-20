package handler

import (
	"employees/internal/domain"
	"employees/internal/repository"
	"employees/internal/transport"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Обработчик ошибок.
func HandleError(context *gin.Context, message string, err error) {
	logrus.Errorf("%s: %s", message, err.Error())
	switch err {
	case domain.ErrAccessTokenIncorrect, domain.ErrSignInDataIncorrect:
		context.AbortWithStatusJSON(http.StatusUnauthorized, transport.ErrorResponse{Message: message})
	case repository.ErrRecordNotFound, ErrQueryParamsIncorrectFormat, ErrRequestBodyIncorrectFormat:
		context.AbortWithStatusJSON(http.StatusBadRequest, transport.ErrorResponse{Message: message})
	default:
		context.AbortWithStatusJSON(http.StatusInternalServerError, transport.ErrorResponse{Message: message})
	}
}
