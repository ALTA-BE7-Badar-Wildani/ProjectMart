package handlers

import (
	"go-ecommerce/config"
	"go-ecommerce/delivery/helpers"
	"go-ecommerce/entities/web"
	userService "go-ecommerce/services/user"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *userService.UserService
}

func NewUserHandler(service *userService.UserService) UserHandler {
	return UserHandler{
		userService: service,
	}
}

func (handler UserHandler) Index(c echo.Context) error {

	// Translate query param to map of filters
	filters := []map[string]string{}
	q := c.QueryParam("q") 
	if q != "" {
		filters = append(filters, map[string]string{
			"field": "name",
			"operator": "LIKE",
			"value": "%" + q + "%",
		})
	} 
	username := c.QueryParam("username") 
	if username != "" {
		filters = append(filters, map[string]string{
			"field": "username",
			"operator": "=",
			"value": username,
		})
	} 
	gender := c.QueryParam("gender") 
	if gender != "" {
		filters = append(filters, map[string]string{
			"field": "gender",
			"operator": "=",
			"value": gender,
		})
	} 
	address := c.QueryParam("address") 
	if address != "" {
		filters = append(filters, map[string]string{
			"field": "address",
			"operator": "=",
			"value": address,
		})
	}
	links := map[string]string {"self": config.Get().App.BaseUrl + "/api/users"}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, "Limit Parameter format is invalid", links))
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		links := map[string]string {"self": config.Get().App.BaseUrl}
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, "page Parameter format is invalid", links))
	}

	usersRes, err := handler.userService.FindAll(limit, page, filters, []map[string]interface{}{})
	if err != nil {
		if reflect.TypeOf(err).String() == "web.RepositoryError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
		panic("not returning custom error")
	}
	
	pagination, err := handler.userService.GetPagination(limit, page)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.RepositoryError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
		panic("not returning custom error")
	}
	
	return c.JSON(200, web.SuccessResponse{
		Status: "OK",
		Code: 200,
		Error: nil,
		Links: links,
		Data: usersRes,
		Pagination: pagination,
	})
}