package main

import (
	"go-ecommerce/config"
	"go-ecommerce/utilities"
)

func main() {

	config := config.Get()
	utilities.NewGormConnection(config)

}