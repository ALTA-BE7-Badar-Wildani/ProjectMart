package web

import "time"

type ProductCartResponse struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Price int `json:"price"`
	Description string `json:"description"`
	Image string `json:"image"`
	CategoryID int `json:"category_id"`
	Category CategoryResponse `json:"category"`
	UserID int `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}