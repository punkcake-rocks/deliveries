package main

import (
	"log"

	goshopify "github.com/bold-commerce/go-shopify/v3"
)

func shopify() {
	// Create an app somewhere.
	app := goshopify.App{
		ApiKey:   cfg.Shopify.ApiKey,
		Password: cfg.Shopify.Password,
	}

	// Create a new API client (notice the token parameter is the empty string)
	client := goshopify.NewClient(app, cfg.Shopify.ShopName, cfg.Shopify.ApiToken)

	// Fetch the number of products.
	numProducts, err := client.Order.Count(nil)
	log.Println(numProducts)
	log.Println(err)
}
