package main

import (
	"go-ecommerce/config"
	"go-ecommerce/delivery/handlers"
	"go-ecommerce/delivery/routes"
	userRepository "go-ecommerce/repositories/user"
	authService "go-ecommerce/services/auth"
	userService "go-ecommerce/services/user"
	"go-ecommerce/utilities"

	"github.com/labstack/echo/v4"
)

func main() {

	config := config.Get()
	db := utilities.NewGormConnection(config)
	utilities.Migrate(db)


	e := echo.New()

	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	routes.RegisterUserRoute(e, userHandler)

	authService := authService.NewAuthService(userRepository)
	authHandler := handlers.NewAuthHandler(authService)
	routes.RegisterAuthRoute(e, authHandler)

	e.Logger.Fatal(e.Start(":" + config.App.Port))
}