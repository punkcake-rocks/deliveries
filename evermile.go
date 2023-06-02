package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	evermileApi "github.com/GIT_USER_ID/GIT_REPO_ID/EvermileClient"
	"golang.org/x/oauth2/clientcredentials"
)

func getClientContext(ctx context.Context) context.Context {
	apiSubdomain := cfg.Evermile.ApiSubdomain

	// Set up authentication
	oauthConfig := &clientcredentials.Config{
		ClientID:     cfg.Evermile.ClientId,
		ClientSecret: cfg.Evermile.ClientSecret,
		TokenURL:     fmt.Sprintf("https://auth.%s.evermile.io/oauth2/token", apiSubdomain),
	}

	// Put the token source in the context
	ctx = context.WithValue(ctx, evermileApi.ContextOAuth2, oauthConfig.TokenSource(ctx))
	// Set the server index to 0 (needed for the sandbox and prod env)
	ctx = context.WithValue(ctx, evermileApi.ContextServerIndex, 0)
	// Set the server variables api subdomain to prod/sandbox
	ctx = context.WithValue(ctx, evermileApi.ContextServerVariables, map[string]string{
		"api-subdomain": apiSubdomain,
	})
	return ctx
}

func setupApi(orderId int) (*evermileApi.APIClient, context.Context) {
	ctx := getClientContext(context.Background())

	evermileApiConfig := evermileApi.NewConfiguration()
	// Replace merchant ID with your prod one for prod
	evermileApiConfig.AddDefaultHeader("X-EVERMILE-MERCHANT-ID", cfg.Evermile.MechantId)

	// A trace ID for tracing the request through the Evermile platform
	rand.Seed(time.Now().UnixNano())
	xEVERMILETRACEID := strconv.Itoa(orderId) + "_" + strconv.Itoa(rand.Int())
	evermileApiConfig.AddDefaultHeader("X-EVERMILE-TRACE-ID", xEVERMILETRACEID)

	// A store ID for an order's store platform context
	xEVERMILESTOREID := cfg.Evermile.DefaultLocation
	evermileApiConfig.AddDefaultHeader("X-EVERMILE-STORE-ID", xEVERMILESTOREID)

	evermileApiConfig.Debug = true

	return evermileApi.NewAPIClient(evermileApiConfig), ctx
}

func deliverySlot(deadline time.Time) *evermileApi.DeliverySlot {
	year, month, date := deadline.Date()
	// loc, _ := time.LoadLocation("Europe/London")
	loc, _ := time.LoadLocation("UTC")

	deliverySlot := evermileApi.DeliverySlot{
		Start: time.Date(year, month, date-2, 0, 0, 0, 0, loc),
		End:   time.Date(year, month, date-1, 0, 0, 0, 0, loc),
	}
	// deliverySlot := evermileApi.DeliverySlot{
	// 	Start: time.Date(year, month, date, 0, 0, 0, 0, loc),
	// 	End:   time.Date(year, month, date+1, 0, 0, 0, 0, loc),
	// }
	return &deliverySlot
}

func evermile(formData OrdersCreateBody) {

	api, ctx := setupApi(int(formData.OrderNumber))
	destinationLocations := []evermileApi.DestinationLocation{
		{
			Address: &evermileApi.Address{
				AddressLine1: formData.ShippingAddress.Address1,
				City:         formData.ShippingAddress.City,
				PostalCode:   formData.ShippingAddress.Zip,
				Type:         evermileApi.OFFICE,
			},
			DeliverySlot: deliverySlot(formData.DeliverBy),
		},
	}

	quoteReq := *evermileApi.NewQuoteReq(
		[]string{cfg.Evermile.DefaultLocation},
		destinationLocations,
		[]evermileApi.Parcel{
			*evermileApi.NewParcel(
				*evermileApi.NewDimensions(float32(30), float32(30), float32(20)),
				float32(1.5),
				evermileApi.ParcelType(evermileApi.PACKAGE),
			),
		},
	)

	resp, r, err := api.QuotesAPI.QuotePost(ctx).QuoteReq(quoteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.QuotePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuotePost`: QuoteResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesApi.QuotePost`: %v\n", resp)
	log.Println(destinationLocations[0].DeliverySlot.GetEnd().Format(time.RFC3339))
	log.Println(resp.DateProposals[1].Proposals)
}
