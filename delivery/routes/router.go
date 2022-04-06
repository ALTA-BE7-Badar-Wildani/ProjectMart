package routes

import (
	handler "go-ecommerce/delivery/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoute(e *echo.Echo, userHandler handler.UserHandler) {
	e.GET("/api/users", userHandler.Index)
	e.GET("/api/users/:id", userHandler.Show)
	e.POST("/api/users", userHandler.Create)
	e.PUT("/api/users/:id", userHandler.Update)
	e.DELETE("/api/users/:id", userHandler.Delete)
}