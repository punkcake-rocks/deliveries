package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Evermile struct {
		MechantId       string `env:"EVERMILE_MERCHANT_ID"`
		ClientId        string `env:"EVERMILE_CLIENT_ID"`
		ClientSecret    string `env:"EVERMILE_CLIENT_SECRET"`
		ApiSubdomain    string `env:"EVERMILE_API_SUBDOMAIN"`
		DefaultLocation string `env:"EVERMILE_DEFAULT_LOCATION"`
		Debug           bool   `env:"EVERMILE_DEBUG"`
	}
	Shopify struct {
		WebhookSignature string `env:"SHOPIFY_WEBHOOK_SIGNATURE"`
		ApiKey           string `env:"SHOPIFY_API_KEY"`
		Password         string `env:"SHOPIFY_PASSWORD"`
		ShopName         string `env:"SHOPIFY_SHOP_NAME"`
		ApiToken         string `env:"SHOPIFY_API_TOKEN"`
	}
}

var cfg Config

func main() {

	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cfg)
	http.HandleFunc(("/webhook"), webhook)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
