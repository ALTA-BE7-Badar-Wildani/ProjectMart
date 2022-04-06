package domain

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID uint
	User User `gorm:"foreignKey:UserID;references:ID"`
	AddressStreet string
	AddressCity string
	AddressProvince string
	AddressZipCode string
	PaymentCard string
	PaymentCardName string
	PaymentCardNumber string
	PaymentCardExp string
	TotalQty int
	TotalPrice int64
	Status string
}