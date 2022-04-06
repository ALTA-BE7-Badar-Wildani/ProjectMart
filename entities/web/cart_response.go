package web

import "time"

type CartResponse struct {
	ID uint `json:"id"`
	ProductID uint `json:"product_id"`
	Product ProductCartResponse `json:"product"`
	ProductPrice int64 `json:"product_price"`
	TransactionID int `json:"transaction_id"`
	Qty int `json:"qty"`
	SubTotal int64 `json:"sub_total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}