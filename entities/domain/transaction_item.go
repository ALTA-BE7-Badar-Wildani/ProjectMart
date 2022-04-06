package domain

import "gorm.io/gorm"

type TransactionItem struct {
	gorm.Model
	ProductID uint
	Product Product `gorm:"foreignKey:ProductID;references:ID"`
	TransactionID uint
	Transaction Transaction `gorm:"foreignKey:TransactionID;references:ID"`
	ProductPrice int64
	Qty int
	SubTotal int64
}