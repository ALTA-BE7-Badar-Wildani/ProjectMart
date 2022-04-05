package main

import (
	"fmt"
	"go-ecommerce/config"
)

func main() {

	config := config.Get()
	fmt.Println(config)

}