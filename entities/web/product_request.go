package web

type ProductRequest struct {
	Title string `form:"title"`
	Price int `form:"price"`
	Description string `form:"description"`
	Image string `form:"image"`
	CategoryID int `form:"category_id"`
}