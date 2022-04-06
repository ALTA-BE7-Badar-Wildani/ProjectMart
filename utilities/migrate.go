package utilities

import (
	entityDomain "go-ecommerce/entities/domain"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&entityDomain.User{},
		&entityDomain.Category{},
		&entityDomain.Product{},
	)
}