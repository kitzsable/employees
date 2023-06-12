package domain

import (
	"employees/internal/models"
	"employees/internal/repository"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Константы для генерации токена.
const (
	// Время жизни токена.
	tokenExpirationTime = 12 * time.Hour

	// Ключ подписи токена.
	signingKey = "elwn3odneucje92+os9%hwn22&o"
)

// Возможные ошибки в работе сервиса.
var (
	// Ошибка некорректных данных аутентификации.
	ErrSignInDataIncorrect = errors.New("incorrect sign in data")

	// Ошибка некорректного токена.
	ErrAccessTokenIncorrect = errors.New("incorrect access token")
)

// Структура данных токена.
type tokenClaims struct {
	jwt.StandardClaims
	UserId int64 `json:"userId"`
}

// Сервис для работы с авторизацией / аутентификацией.
type AuthService struct {
	userRepository UserRepository
}

// Создание нового сервиса для работы с авторизацией / аутентификацией.
func NewAuthService(userRepository UserRepository) AuthService {
	return AuthService{
		userRepository: userRepository,
	}
}

// Создание пользователя (регистрация).
func (service AuthService) CreateUser(user *models.User) (int64, error) {
	var id int64 = 0

	passwordHash, err := service.generatePasswordHash(user.Password)
	if err != nil {
		return id, fmt.Errorf("cannot generate password hash: %v", err)
	}
	user.PasswordHash = passwordHash

	id, err = service.userRepository.Create(user)
	if err != nil {
		return id, err
	}

	return id, nil
}

// Аутентификация пользователя.
func (service AuthService) SignInUser(user *models.User) (string, error) {
	var signingToken = ""

	userDB, err := service.userRepository.Get(user.Username)
	if err != nil {
		if err == repository.ErrRecordNotFound {
			return signingToken, ErrSignInDataIncorrect
		}
		return signingToken, err
	}

	isPasswordCorrect := service.checkPasswordHash(user.Password, userDB.PasswordHash)

	if !isPasswordCorrect {
		return signingToken, ErrSignInDataIncorrect
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpirationTime).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userDB.Id,
	})

	signingToken, err = token.SignedString([]byte(signingKey))
	if err != nil {
		return signingToken, fmt.Errorf("cannot generate token: %v", err)
	}

	return signingToken, nil
}

// Проверка токена пользователя (авторизация).
func (service AuthService) ParseAccessToken(accessToken string) (int64, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrAccessTokenIncorrect
		}

		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, ErrAccessTokenIncorrect
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, ErrAccessTokenIncorrect
	}

	return claims.UserId, nil
}

// Генерация хэша пароля.
func (service AuthService) generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}

// Проверка хэша пароля.
func (service AuthService) checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
