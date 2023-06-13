package main

import (
	"net/http"

	goshopify "github.com/bold-commerce/go-shopify/v3"
)

var shopifyApp goshopify.App
var shopifyClient *goshopify.Client

func initShopify() {
	shopifyApp = goshopify.App{
		ApiKey:    cfg.Shopify.ApiKey,
		Password:  cfg.Shopify.Password,
		ApiSecret: cfg.Shopify.WebhookSignature,
	}

	shopifyClient = goshopify.NewClient(shopifyApp, cfg.Shopify.ShopName, cfg.Shopify.ApiToken)
}

func ValidateWebhook(httpRequest *http.Request) bool {
	return shopifyApp.VerifyWebhookRequest(httpRequest)
}
