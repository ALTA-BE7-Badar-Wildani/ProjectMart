package web

type TransactionRequest struct {
	AddressStreet string `form:"address_street"`
	AddressCity string `form:"address_city"`
	AddressProvince string `form:"address_province"`
	AddressZipCode string `form:"address_zip_code"`
	PaymentCard string `form:"payment_card"`
	PaymentCardName string `form:"payment_card_name"`
	PaymentCardNumber string `form:"payment_card_number"`
	PaymentCardExp string `form:"payment_card_exp"`
}