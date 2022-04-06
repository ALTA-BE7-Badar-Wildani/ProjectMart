package web

type CartRequest struct {
	ProductID uint `form:"product_id"` 
	Qty int `form:"qty"`
}