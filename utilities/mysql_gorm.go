package utilities

import (
	"fmt"
	appConfig "go-ecommerce/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormConnection(config *appConfig.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to dataase")
	}
	return db
}