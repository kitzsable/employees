package main

import (
	"employees/internal/repository"
	"employees/internal/server"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.Info("Welcome to The Employees App!")
	if err := initConfig(); err != nil {
		logrus.Fatalf("error loading configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
}

func main() {
	dbConfig := repository.DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		DBName:   viper.GetString("database.dbname"),
		SSLMode:  viper.GetString("database.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	server.Start(server.NewConfig(viper.GetString("port"), dbConfig))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
