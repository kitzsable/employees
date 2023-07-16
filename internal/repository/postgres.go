package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// Названия таблиц базы данных.
const (
	// Название таблицы пользователей системы.
	usersTable = "users"

	// Название таблицы отделов.
	depatmentsTable = "departments"

	// Название таблицы сотрудников.
	employeesTable = "employees"
)

// Возможные ошибки при работе с базой данных.
var (
	// Ошибка отсутствия искомой записи в базе данных.
	ErrRecordNotFound = errors.New("record doesn't exist in database")
)

// Конфигурация базы данных.
type DBConfig struct {
	// Хост, на котором расположена база данных.
	Host string

	// Порт, на котором расположена база данных.
	Port string

	// Имя пользователя базы данных.
	Username string

	// Пароль пользователя базы данных.
	Password string

	// Наименование базы данных.
	DBName string

	// Режим подключения
	SSLMode string
}

// Функция создания нового экземпляра базы данных по её конфигурации.
func NewPostgresDB(dbConfig DBConfig) (*sql.DB, error) {
	databaseURL := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.SSLMode)
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		return db, fmt.Errorf("cannot open database: %v", err)
	}

	for i := 0; i < 10; i++ {
		if err = db.Ping(); err == nil {
			return db, nil
		}
		time.Sleep(1 * time.Second)
	}

	return db, fmt.Errorf("cannot connect to database: %v", err)
}
