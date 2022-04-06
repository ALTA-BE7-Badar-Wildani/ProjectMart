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

	// Get all users
	usersRes, err := handler.userService.FindAll(limit, page, filters, []map[string]interface{}{})
	if err != nil {
		if reflect.TypeOf(err).String() == "web.RepositoryError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
		panic("not returning custom error")
	}
	
	// Get pagination data
	pagination, err := handler.userService.GetPagination(limit, page)
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
		Data: usersRes,
		Pagination: pagination,
	})
}

func (handler UserHandler) Show(c echo.Context) error {
	// Get param
	id, err := strconv.Atoi(c.Param("id"))
	links := map[string]string{ "self": config.Get().App.BaseUrl + "/users/" + c.Param("id") }
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, err.Error(), links))
	}

	// Get userdata 
	user, err := handler.userService.Find(id)
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
		Data: user,
	})
}


func (handler UserHandler) Create(c echo.Context) error {
	// Populate form
	userReq := web.UserRequest{}
	c.Bind(&userReq)
	
	links := map[string]string{ "self": config.Get().App.BaseUrl + "/users"}

	// Insert user
	userRes, err := handler.userService.Create(userReq)
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
		Data: userRes,
	})
}


func (handler UserHandler) Update(c echo.Context) error {
	// Populate form
	userReq := web.UserRequest{}
	c.Bind(&userReq)

	id, err := strconv.Atoi(c.Param("id")) 
	links := map[string]string{ "self": config.Get().App.BaseUrl + "/users/" + c.Param("id")}
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, err.Error(), links))
	}

	userRes, err := handler.userService.Update(userReq, id)
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
		Data: userRes,
	})
}



func (handler UserHandler) Delete(c echo.Context) error {

	// Get params ID
	id, err := strconv.Atoi(c.Param("id")) 
	links := map[string]string{ "self": config.Get().App.BaseUrl + "/users/" + c.Param("id")}
	if err != nil {
		return c.JSON(400, helpers.MakeErrorResponse("ERROR", 400, err.Error(), links))
	}

	// call delete service
	err = handler.userService.Delete(id)
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