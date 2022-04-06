package handlers

import (
	"go-ecommerce/config"
	"go-ecommerce/delivery/helpers"
	"go-ecommerce/entities/web"
	categoryService "go-ecommerce/services/category"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryService *categoryService.CategoryService
}

func NewCategoryHandler(service *categoryService.CategoryService) CategoryHandler {
	return CategoryHandler{
		categoryService: service,
	}
}


/*
 * -------------------------------------------
 * Show All categories based on available queries
 * -------------------------------------------
 */
func (handler CategoryHandler) Index(c echo.Context) error {

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
	links := map[string]string {"self": config.Get().App.BaseUrl + "/api/categories"}

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

	// Get all categories
	categoriesRes, err := handler.categoryService.FindAll(limit, page, filters, sorts)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
		panic("not returning custom error")
	}
	
	// Get pagination data
	pagination, err := handler.categoryService.GetPagination(limit, page)
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
		Data: categoriesRes,
		Pagination: pagination,
	})
}

/*
 * -------------------------------------------
 * Show single category detail by ID
 * -------------------------------------------
 */
func (handler CategoryHandler) Show(c echo.Context) error {
	// Get param
	id, err := strconv.Atoi(c.Param("id"))
	links := map[string]string{ "self": config.Get().App.BaseUrl + "/categories/" + c.Param("id") }
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, err.Error(), links))
	}
	// Get categorydata 
	category, err := handler.categoryService.Find(id)
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
		Data: category,
	})
}

/*
 * -------------------------------------------
 * Create category resource
 * -------------------------------------------
 */
func (handler CategoryHandler) Create(c echo.Context) error {
	// Populate form
	categoryReq := web.CategoryRequest{}
	c.Bind(&categoryReq)
	
	// Define hateoas links
	links := map[string]string{ "self": config.Get().App.BaseUrl + "/categories"}

	// Insert category
	categoryRes, err := handler.categoryService.Create(categoryReq)
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
		Data: categoryRes,
	})
}


/*
 * -------------------------------------------
 * Update category resource
 * -------------------------------------------
 */
func (handler CategoryHandler) Update(c echo.Context) error {
	// Populate form
	categoryReq := web.CategoryRequest{}
	c.Bind(&categoryReq)

	// Get param ID
	id, err := strconv.Atoi(c.Param("id")) 
	links := map[string]string{ "self": config.Get().App.BaseUrl + "/categories/" + c.Param("id")}
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, err.Error(), links))
	}

	// Service call
	categoryRes, err := handler.categoryService.Update(categoryReq, id)
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
		Data: categoryRes,
	})
}


/*
 * -------------------------------------------
 * Delete category resource
 * -------------------------------------------
 */
func (handler CategoryHandler) Delete(c echo.Context) error {

	// Get params ID
	id, err := strconv.Atoi(c.Param("id")) 
	links := map[string]string{ "self": config.Get().App.BaseUrl + "/categories/" + c.Param("id")}
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, err.Error(), links))
	}

	// call delete service
	err = handler.categoryService.Delete(id)
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