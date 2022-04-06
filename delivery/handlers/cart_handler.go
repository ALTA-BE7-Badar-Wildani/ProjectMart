package handlers

import (
	"go-ecommerce/config"
	"go-ecommerce/delivery/helpers"
	"go-ecommerce/entities/web"
	cartService "go-ecommerce/services/cart"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartService *cartService.CartService
}

func NewCartHandler(cartService *cartService.CartService) CartHandler {
	return CartHandler{
		cartService: cartService,
	}
}


/*
 * -------------------------------------------
 * Show All cart items based on authenticated user
 * -------------------------------------------
 */
func (handler CartHandler) Index(c echo.Context) error {
	
	// hateoas
	links := map[string]string {"self": config.Get().App.BaseUrl + "/api/cart"}

	// retrieve token
	token := c.Get("user")

	// service call
	cartRes, err := handler.cartService.FindAll(token)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
		return c.JSON(500, helpers.MakeErrorResponse("ERROR", 500, err.Error(), links))
	}

	// response
	return c.JSON(200, web.SuccessListResponse{
		Status: "OK",
		Code: 200,
		Error: nil,
		Links: links,
		Data: cartRes,
		Pagination: web.Pagination{},
	})
}


/*
 * -------------------------------------------
 * Add cart items
 * -------------------------------------------
 */
func (handler CartHandler) Create(c echo.Context) error {
	
	// hateoas
	links := map[string]string {"self": config.Get().App.BaseUrl + "/api/cart"}

	// retrieve token
	token := c.Get("user")

	// Populate request
	cartReq := web.CartRequest{}
	c.Bind(&cartReq)

	// service call
	cartRes, err := handler.cartService.Create(cartReq, token)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
		return c.JSON(500, helpers.MakeErrorResponse("ERROR", 500, err.Error(), links))
	}

	// response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code: 200,
		Error: nil,
		Links: links,
		Data: cartRes,
	})
}

/*
 * -------------------------------------------
 *  Update cart items
 * -------------------------------------------
 */
func (handler CartHandler) Update(c echo.Context) error {

	// hateoas
	links := map[string]string {"self": config.Get().App.BaseUrl + "/api/cart/"}

	// Retrieve id param
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, "The provided parameter is invalid", links))
	}

	// retrieve token
	token := c.Get("user")

	// Populate request
	cartReq := web.CartRequest{}
	c.Bind(&cartReq)

	// service call
	cartRes, err := handler.cartService.Update(cartReq, id, token)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
		return c.JSON(500, helpers.MakeErrorResponse("ERROR", 500, err.Error(), links))
	}

	// response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code: 200,
		Error: nil,
		Links: links,
		Data: cartRes,
	})
}

