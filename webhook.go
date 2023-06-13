package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type OrdersCreateBody struct {
	Id              int64  `json:"id"`
	Email           string `json:"email"`
	FinancialStatus string `json:"financial_status"`
	Note            string `json:"note"`
	NoteAttributes  []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"note_attributes"`
	OrderNumber     int64  `json:"order_number"`
	Name            string `json:"name"`
	Currency        string `json:"currency"`
	SubTotalPrice   string `json:"subtotal_price"`
	ShippingAddress struct {
		Name     string `json:"name"`
		Address1 string `json:"address1"`
		Address2 string `json:"address2"`
		City     string `json:"city"`
		Company  string `json:"company"`
		Country  string `json:"country"`
		Phone    string `json:"phone"`
		Zip      string `json:"zip"`
	} `json:"shipping_address"`
	DeliverBy time.Time `json:"deliver_by"`
}

func verifyWebhook(data, hmacHeder []byte) bool {
	mac := hmac.New(sha256.New, []byte(cfg.Shopify.WebhookSignature))
	mac.Write(data)

	expectedMAC := []byte(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

	return hmac.Equal(hmacHeder, expectedMAC)
}

func ordersCreate(w http.ResponseWriter, body []byte) {
	var formData OrdersCreateBody

	if err := json.Unmarshal(body, &formData); err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	var deliverByDate time.Time
	var deliverByTime string

	for _, v := range formData.NoteAttributes {

		if v.Name == "Delivery-Date" {
			deliverByDate, _ = time.Parse("2006/01/02", v.Value)
		}
		if v.Name == "Delivery-Time" {
			if !validateDeliveryTime(v.Value) {
				http.Error(w, fmt.Sprintf("Wrong format for delivery time: w%s ", v.Value), http.StatusUnprocessableEntity)
				return
			} else {
				deliverByTime = strings.Split(v.Value, " - ")[1]
			}
		}
	}

	log.Println(deliverByTime)
	if deliverByDate.IsZero() || deliverByTime == "" {
		log.Println(fmt.Sprintf("#%d is a collection, skipping", formData.OrderNumber) + "")
		return
	}

	location, _ := time.LoadLocation("Europe/London")
	formData.DeliverBy, _ = time.ParseInLocation("2006/01/02 15:04", deliverByDate.Format("2006/01/02")+" "+deliverByTime, location)

	// fmt.Printf("FORM DATA: %+v \n", formData)
	evermile(formData)
}

func validateDeliveryTime(deliveryTime string) bool {
	matched, err := regexp.Match(`\d{2}:\d{2} - \d{2}:\d{2}`, []byte(deliveryTime))

	if err != nil {
		log.Println(err)
		return false
	}

	return matched
}

func webhook(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/webhook" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}

	if !verifyWebhook(body, []byte(r.Header.Get("X-Shopify-Hmac-Sha256"))) {
		http.Error(w, "Error verifying webhook", http.StatusUnauthorized)
		return
	}

	switch r.Header.Get("X-Shopify-Topic") {
	case "orders/create":
		ordersCreate(w, body)
		return
	default:
		http.Error(w, "Topic not found", http.StatusBadRequest)
		return
	}

	// reqDump, _ := httputil.DumpRequest(r, true)

	// fmt.Printf("REQUEST:\n%s", string(reqDump))
	// w.Write([]byte("Hello World"))
}
