package handlers

import (
	"go-ecommerce/config"
	"go-ecommerce/delivery/helpers"
	"go-ecommerce/entities/web"
	productService "go-ecommerce/services/product"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService *productService.ProductService
}

func NewProductHandler(service *productService.ProductService) ProductHandler {
	return ProductHandler{
		productService: service,
	}
}


/*
 * -------------------------------------------
 * Show All products based on available queries
 * -------------------------------------------
 */
func (handler ProductHandler) Index(c echo.Context) error {

	// Translate query param to map of filters
	filters := []map[string]string{}
	q := c.QueryParam("q") 
	if q != "" {
		filters = append(filters, map[string]string{
			"field": "title",
			"operator": "LIKE",
			"value": "%" + q + "%",
		})
	} 
	// Sort parameter
	sorts := []map[string]interface{} {}
	sortPrice := c.QueryParam("sortPrice") 
	if sortPrice != "" {
		switch sortPrice {
		case "1":
			sorts = append(sorts, map[string]interface{} {
				"field": "title",
				"desc": true,
			})
		case "0":
			sorts = append(sorts, map[string]interface{} {
				"field": "title",
				"desc": false,
			})
		}
	}
	links := map[string]string {"self": config.Get().App.BaseUrl + "/api/products"}

	// pagination param
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, "Limit Parameter format is invalid", links))
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		links := map[string]string {"self": config.Get().App.BaseUrl}
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, "page Parameter format is invalid", links))
	}

	// Get all products
	productsRes, err := handler.productService.FindAll(limit, page, filters, sorts)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
		panic("not returning custom error")
	}
	
	// Get pagination data
	pagination, err := handler.productService.GetPagination(limit, page)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
		panic("not returning custom error")
	}
	
	// success response
	return c.JSON(200, web.SuccessListResponse{
		Status: "OK",
		Code: 200,
		Error: nil,
		Links: links,
		Data: productsRes,
		Pagination: pagination,
	})
}

/*
 * -------------------------------------------
 * Show single product detail by ID
 * -------------------------------------------
 */
func (handler ProductHandler) Show(c echo.Context) error {
	// Get param
	id, err := strconv.Atoi(c.Param("id"))
	links := map[string]string{ "self": config.Get().App.BaseUrl + "/products/" + c.Param("id") }
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, err.Error(), links))
	}
	// Get productdata 
	product, err := handler.productService.Find(id)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
	}
	// response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code: 200,
		Error: nil,
		Links: links,
		Data: product,
	})
}


/*
 * -------------------------------------------
 * Get All user's products based on available queries
 * -------------------------------------------
 */
func (handler ProductHandler) GetUserProduct(c echo.Context) error {

	// Translate query param to map of filters
	filters := []map[string]string{}
	q := c.QueryParam("q") 
	if q != "" {
		filters = append(filters, map[string]string{
			"field": "title",
			"operator": "LIKE",
			"value": "%" + q + "%",
		})
	} 
	// Sort parameter
	sorts := []map[string]interface{} {}
	sortPrice := c.QueryParam("sortPrice") 
	if sortPrice != "" {
		switch sortPrice {
		case "1":
			sorts = append(sorts, map[string]interface{} {
				"field": "title",
				"desc": true,
			})
		case "0":
			sorts = append(sorts, map[string]interface{} {
				"field": "title",
				"desc": false,
			})
		}
	}
	links := map[string]string {"self": config.Get().App.BaseUrl + "/api/products"}

	// pagination param
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, "Limit Parameter format is invalid", links))
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		links := map[string]string {"self": config.Get().App.BaseUrl}
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, "page Parameter format is invalid", links))
	}

	// get user param ID
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, "requested id is invalid", links))
	}
	filters = append(filters, map[string]string{
		"field": "user_id",
		"operator": "=",
		"value": strconv.Itoa(userID),
	})

	// Get all products
	productsRes, err := handler.productService.FindAll(limit, page, filters, sorts)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
		panic("not returning custom error")
	}
	
	// Get pagination data
	pagination, err := handler.productService.GetPagination(limit, page)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
		panic("not returning custom error")
	}
	
	// success response
	return c.JSON(200, web.SuccessListResponse{
		Status: "OK",
		Code: 200,
		Error: nil,
		Links: links,
		Data: productsRes,
		Pagination: pagination,
	})
}

/*
 * -------------------------------------------
 * Create product resource
 * -------------------------------------------
 */
func (handler ProductHandler) Create(c echo.Context) error {
	// Populate form
	productReq := web.ProductRequest{}
	c.Bind(&productReq)
	
	// Define hateoas links
	links := map[string]string{ "self": config.Get().App.BaseUrl + "/products"}

	token := c.Get("user")

	// Insert product
	productRes, err := handler.productService.Create(productReq, token)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
	}

	// response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code: 200,
		Error: nil,
		Links: links,
		Data: productRes,
	})
}


/*
 * -------------------------------------------
 * Update product resource
 * -------------------------------------------
 */
func (handler ProductHandler) Update(c echo.Context) error {
	// Populate form
	productReq := web.ProductRequest{}
	c.Bind(&productReq)

	id, err := strconv.Atoi(c.Param("productID")) 
	links := map[string]string{ "self": config.Get().App.BaseUrl + "/products/" + c.Param("id")}
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, err.Error(), links))
	}

	token := c.Get("user")

	// Product service call
	productRes, err := handler.productService.Update(productReq, id, token)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
	}

	// response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code: 200,
		Error: nil,
		Links: links,
		Data: productRes,
	})
}


/*
 * -------------------------------------------
 * Delete product resource
 * -------------------------------------------
 */
func (handler ProductHandler) Delete(c echo.Context) error {

	// Get params ID
	id, err := strconv.Atoi(c.Param("productID")) 
	links := map[string]string{ "self": config.Get().App.BaseUrl + "/products/" + c.Param("id")}
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, err.Error(), links))
	}

	token := c.Get("user")

	// call delete on product service
	err = handler.productService.Delete(id, token)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
	}

	// response
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code: 200,
		Error: nil,
		Links: links,
		Data: map[string]interface{} {
			"id": id,
		},
	})
}