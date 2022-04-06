package handlers

import (
	"go-ecommerce/config"
	"go-ecommerce/delivery/helpers"
	"go-ecommerce/entities/web"
	authService "go-ecommerce/services/auth"
	"reflect"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authService *authService.AuthService
}

func NewAuthHandler(service *authService.AuthService) AuthHandler {
	return AuthHandler{
		authService: service,
	}
}

func (handler AuthHandler) Login(c echo.Context) error {
	// Populate request input
	authReq := web.AuthRequest {
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}

	links := map[string]string { "self": config.Get().App.BaseUrl + "/api/auth" }

	// Call auth service login
	authRes, err := handler.authService.Login(authReq)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
		return c.JSON(500, helpers.MakeErrorResponse("ERROR", 500, err.Error(), links))
	}

	// send response
	return c.JSON(200, web.SuccessResponse {
		Status: "OK",
		Code: 200,
		Error: nil,
		Links: links,
		Data: authRes,
	})
}

func (handler AuthHandler) Me(c echo.Context) error {
	userJWT := c.Get("user")
	links := map[string]string { "self": config.Get().App.BaseUrl + "/api/auth/me" }

	authRes, err := handler.authService.Me(userJWT)
	if err != nil {
		if reflect.TypeOf(err).String() == "web.WebError" {
			webErr := err.(web.WebError)
			return c.JSON(webErr.Code, helpers.MakeErrorResponse("ERROR", webErr.Code, webErr.Error(), links))
		}
		return c.JSON(500, helpers.MakeErrorResponse("ERROR", 500, err.Error(), links))
	}

	return c.JSON(200, web.SuccessResponse {
		Status: "OK",
		Code: 200,
		Error: nil,
		Links: links,
		Data: authRes,
	})
}