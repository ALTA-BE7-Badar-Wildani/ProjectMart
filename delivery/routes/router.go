package routes

import (
	"go-ecommerce/delivery/handlers"
	handler "go-ecommerce/delivery/handlers"
	webMiddleware "go-ecommerce/delivery/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoute(e *echo.Echo, userHandler handler.UserHandler) {
	e.GET("/api/users", userHandler.Index)
	e.GET("/api/users/:id", userHandler.Show)
	e.POST("/api/users", userHandler.Create)
	e.PUT("/api/users/:id", userHandler.Update)
	e.DELETE("/api/users/:id", userHandler.Delete)
}

func RegisterAuthRoute(e *echo.Echo, authHandler handler.AuthHandler) {
	e.POST("/api/auth", authHandler.Login)
	e.GET("/api/auth/me", authHandler.Me, webMiddleware.JWTMiddleware())
}

func RegisterProductRoute(e *echo.Echo, productHandler handler.ProductHandler) {
	e.GET("/api/products", productHandler.Index)
	e.GET("/api/products/:id", productHandler.Show)
	e.GET("/api/users/:id/products", productHandler.GetUserProduct)
	e.POST("/api/users/:id/products", productHandler.Create, webMiddleware.JWTMiddleware())
	e.PUT("/api/users/:id/products/:productID", productHandler.Update, webMiddleware.JWTMiddleware())
	e.DELETE("/api/users/:id/products/:productID", productHandler.Delete, webMiddleware.JWTMiddleware())
}

func RegisterCategoryRoute(e *echo.Echo, categoryHandler handler.CategoryHandler) {
	e.GET("/api/categories", categoryHandler.Index)
	e.GET("/api/categories/:id", categoryHandler.Show)
	e.POST("/api/categories", categoryHandler.Create)
	e.PUT("/api/categories/:id", categoryHandler.Update)
	e.DELETE("/api/categories/:id", categoryHandler.Delete)
}

func RegisterCartRoute(e *echo.Echo, cartHandler handler.CartHandler) {
	e.GET("/api/carts", cartHandler.Index, webMiddleware.JWTMiddleware())
	e.POST("/api/carts", cartHandler.Create, webMiddleware.JWTMiddleware())
	e.PUT("/api/carts/:id", cartHandler.Update, webMiddleware.JWTMiddleware())
	e.DELETE("/api/carts/:id", cartHandler.Delete, webMiddleware.JWTMiddleware())
	e.PUT("/api/carts/checkout", cartHandler.Checkout, webMiddleware.JWTMiddleware())
}

func RegisterTransactionRoute(e *echo.Echo, transactionHandler handlers.TransactionHandler) {
	e.GET("/api/users/:id/transactions", transactionHandler.Index, webMiddleware.JWTMiddleware())
}