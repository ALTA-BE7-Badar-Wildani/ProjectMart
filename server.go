package main

import (
	"go-ecommerce/config"
	"go-ecommerce/utilities"

	"github.com/labstack/echo/v4"
)

func main() {

	config := config.Get()
	db := utilities.NewGormConnection(config)
	utilities.Migrate(db)


	e := echo.New()
	e.Logger.Fatal(e.Start(":" + config.App.Port))
}