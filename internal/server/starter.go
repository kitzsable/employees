package server

import (
	"database/sql"
	"employees/internal/domain"
	"employees/internal/handler"
	"employees/internal/repository"
	"fmt"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sirupsen/logrus"
)

// Запуск сервера с конфигурацией.
func Start(config *Config) error {
	logrus.Info(fmt.Sprintf("Server starting with %+v", config))

	database, err := repository.NewPostgresDB(config.DBConfig)
	if err != nil {
		logrus.Fatalf("error creating database: %s", err.Error())
	}
	defer database.Close()

	handler, err := createHandler(database)
	if err != nil {
		logrus.Fatalf("error creating handler: %s", err.Error())
	}

	server := newServer(handler)
	return http.ListenAndServe(config.BindAddress, server)
}

func createHandler(database *sql.DB) (handler.Handler, error) {
	departmentRepository := repository.NewDepartmentRepositoryImpl(database)
	departmentHandler := handler.NewDepartmentHandler(departmentRepository)

	employeeRepository := repository.NewEmployeeRepositoryImpl(database)
	employeeHandler := handler.NewEmployeeHandler(employeeRepository)

	userRepository := repository.NewUserRepositoryImpl(database)
	authService := domain.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)
	authMiddleware := handler.NewAuthMiddleware(authService)

	handler := handler.NewHandler(authMiddleware, authHandler, departmentHandler, employeeHandler)

	return handler, nil
}
