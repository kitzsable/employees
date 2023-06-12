package server

import "employees/internal/repository"

// Конфигурация сервера.
type Config struct {
	// Адрес, на котором запускается сервер.
	BindAddress string

	// Конфигурация базу данных.
	DBConfig repository.DBConfig
}

// Создание новой конфигурации сервера.
func NewConfig(bindAddress string, dbConfig repository.DBConfig) *Config {
	return &Config{
		BindAddress: bindAddress,
		DBConfig:    dbConfig,
	}
}
