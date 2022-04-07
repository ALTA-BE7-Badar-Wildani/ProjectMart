package main

import (
	"go-ecommerce/config"
	"go-ecommerce/delivery/handlers"
	"go-ecommerce/delivery/routes"
	categoryRepository "go-ecommerce/repositories/category"
	productRepository "go-ecommerce/repositories/product"
	trRepository "go-ecommerce/repositories/transaction"
	trItemRepository "go-ecommerce/repositories/transaction_item"
	userRepository "go-ecommerce/repositories/user"
	authService "go-ecommerce/services/auth"
	cartService "go-ecommerce/services/cart"
	categoryService "go-ecommerce/services/category"
	productService "go-ecommerce/services/product"
	trService "go-ecommerce/services/transaction"
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

	// Product App Provider
	categoryRepository := categoryRepository.NewCategoryRepository(db)
	categoryService := categoryService.NewCategoryService(categoryRepository)
	serviceHandler := handlers.NewCategoryHandler(categoryService)
	routes.RegisterCategoryRoute(e, serviceHandler)

	trRepository := trRepository.NewTransactionRepository(db)
	trItemRepository := trItemRepository.NewTransactionItemRepository(db)
	cartService := cartService.NewCartService(trItemRepository, trRepository, userRepository, productRepository)
	cartHandler := handlers.NewCartHandler(cartService)
	routes.RegisterCartRoute(e, cartHandler)
	
	// User's transaction
	transactionService := trService.NewTransactionService(trRepository, userRepository)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	routes.RegisterTransactionRoute(e, transactionHandler)

	e.Logger.Fatal(e.Start(":" + config.App.Port))
}