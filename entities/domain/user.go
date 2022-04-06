package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id string 
	Name string 
	Email string 
	Username string 
	Password string 
	Gender string 
	Address string 
	Avatar string 
	CreatedAt time.Time
	UpdatedAt time.Time
}