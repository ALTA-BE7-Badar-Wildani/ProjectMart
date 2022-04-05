package main

import (
	"go-ecommerce/config"
	"go-ecommerce/utilities"
)

func main() {

	config := config.Get()
	db := utilities.NewGormConnection(config)
	utilities.Migrate(db)

}