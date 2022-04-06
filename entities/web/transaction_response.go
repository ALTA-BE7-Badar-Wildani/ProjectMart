package web

import "time"

type TransactionResponse struct {
	ID uint `json:"id"`
	AddressStreet string `json:"address_street"`
	AddressCity string `json:"address_city"`
	AddressProvince string `json:"address_province"`
	AddressZipCode string `json:"address_zip_code"`
	PaymentCart string `json:"payment_card"`
	PaymentCardName string `json:"payment_card_name"`
	PaymentCartNumber string `json:"payment_card_number"`
	PaymentCardExp string `json:"payment_card_exp"`
	TotalQty int `json:"total_qty"`
	TotalPrice int64 `json:"total_price"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}