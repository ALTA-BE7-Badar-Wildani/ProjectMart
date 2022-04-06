package main

import (
	"go-ecommerce/config"
	"go-ecommerce/delivery/handlers"
	"go-ecommerce/delivery/routes"
	productRepository "go-ecommerce/repositories/product"
	userRepository "go-ecommerce/repositories/user"
	authService "go-ecommerce/services/auth"
	productService "go-ecommerce/services/product"
	userService "go-ecommerce/services/user"
	"go-ecommerce/utilities"

	"github.com/labstack/echo/v4"
)

func main() {

	config := config.Get()
	db := utilities.NewGormConnection(config)
	utilities.Migrate(db)


	e := echo.New()

	// User App Provider
	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	routes.RegisterUserRoute(e, userHandler)

	// Auth App Provider
	authService := authService.NewAuthService(userRepository)
	authHandler := handlers.NewAuthHandler(authService)
	routes.RegisterAuthRoute(e, authHandler)

	// Product App Provider
	productRepository := productRepository.NewProductRepository(db)
	productService := productService.NewProductService(productRepository, userRepository)
	productHandler := handlers.NewProductHandler(productService)
	routes.RegisterProductRoute(e, productHandler)

	e.Logger.Fatal(e.Start(":" + config.App.Port))
}