package routes

import (
	handler "go-ecommerce/delivery/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoute(e *echo.Echo, userHandler handler.UserHandler) {
	e.GET("/api/users", userHandler.Index)
}