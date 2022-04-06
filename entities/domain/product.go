package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title string
	Price int
	Description string
	Image string
	CategoryID uint
	UserID uint
	Category Category `gorm:"foreignKey:CategoryID;references:ID"`
	User User `gorm:"foreignKey:UserID;references:ID"`
}