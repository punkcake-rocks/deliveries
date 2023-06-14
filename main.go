package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	config()
	initShopify()
	initEvermile()

	http.HandleFunc(("/webhook"), webhook)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
