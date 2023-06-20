package server

import (
	"context"
	"database/sql"
	"employees/internal/domain"
	"employees/internal/handler"
	"employees/internal/repository"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/sirupsen/logrus"
)

// Запуск сервера с конфигурацией.
func Start(config *Config) {
	database, err := repository.NewPostgresDB(config.DBConfig)
	if err != nil {
		logrus.Fatalf("error creating database: %s", err.Error())
	}

	handler, err := createHandler(database)
	if err != nil {
		logrus.Fatalf("error creating handler: %s", err.Error())
	}

	router := handler.ConfigureRouter()

	server := newServer(config.BindAddress, router)
	go server.ListenAndServe()

	logrus.Info("server started")

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGTERM, syscall.SIGINT)
	<-quitChannel

	logrus.Print("server is shutting down")

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error shutting down server: %s", err.Error())
	}

	if err := database.Close(); err != nil {
		logrus.Errorf("error closing database connection: %s", err.Error())
	}
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
