package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	App struct {
		Port string
		BaseUrl string
	}
	Database struct {
		Username string
		Password string
		Host string
		Port string
		Name string
		Driver string
	}
}

var appConfig *AppConfig

func Get() *AppConfig {
	if appConfig == nil {
		appConfig = initConfig()
	}
	return appConfig
}


func initConfig() *AppConfig {

	appConfig := AppConfig{}

	appConfig.App.Port = "8000"
	appConfig.App.BaseUrl = "localhost"
	appConfig.Database.Username = "root"
	appConfig.Database.Password = "root"
	appConfig.Database.Host = "localhost"
	appConfig.Database.Port = "3306"
	appConfig.Database.Name = "ecommerce"
	appConfig.Database.Driver = "mysql"

	err := godotenv.Load()
	if err != nil {
		return &appConfig
	}

	appConfig.App.Port = os.Getenv("APP_PORT")
	appConfig.App.BaseUrl = os.Getenv("APP_BASE_URL")
	appConfig.Database.Username = os.Getenv("DB_USERNAME")
	appConfig.Database.Password = os.Getenv("DB_PASSWORD")
	appConfig.Database.Host = os.Getenv("DB_HOST")
	appConfig.Database.Port = os.Getenv("DB_PORT")
	appConfig.Database.Name = os.Getenv("DB_NAME")
	appConfig.Database.Driver = os.Getenv("DB_DRIVER")

	return &appConfig
}