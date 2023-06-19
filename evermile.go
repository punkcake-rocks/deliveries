package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	evermileApi "github.com/GIT_USER_ID/GIT_REPO_ID/EvermileClient"
	goshopify "github.com/bold-commerce/go-shopify/v3"
	"golang.org/x/oauth2/clientcredentials"
)

var evermileClient *evermileApi.APIClient
var evermileContext context.Context

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

func initEvermile() {
	evermileContext = getClientContext(context.Background())

	evermileApiConfig := evermileApi.NewConfiguration()
	// Replace merchant ID with your prod one for prod
	evermileApiConfig.AddDefaultHeader("X-EVERMILE-MERCHANT-ID", cfg.Evermile.MechantId)

	evermileApiConfig.Debug = cfg.Evermile.Debug

	evermileClient = evermileApi.NewAPIClient(evermileApiConfig)

	// registerWebhooks()
}

func evermile(order goshopify.Order, deadline time.Time) {

	deliverySlot := evermileApi.NewDeliverySlot(
		deadline.Add(time.Duration(-4)*time.Hour).UTC(),
		deadline.UTC(),
	)

	destinationLocations := []evermileApi.DestinationLocation{
		{
			Address: &evermileApi.Address{
				AddressLine1: order.ShippingAddress.Address1,
				City:         order.ShippingAddress.City,
				PostalCode:   order.ShippingAddress.Zip,
				Type:         evermileApi.ADDRESSTYPE_OFFICE,
			},
			DeliverySlot: deliverySlot,
		},
	}

	parcel := *evermileApi.NewParcel(
		*evermileApi.NewDimensions(float32(30), float32(30), float32(20)),
		float32(1.5),
		evermileApi.ParcelType(evermileApi.PARCELTYPE_PACKAGE),
	)

	item := *evermileApi.NewItem(
		"Cake",
		*evermileApi.NewPrice(int64(order.SubtotalPrice.IntPart()*100), order.Currency),
		1,
	)

	parcel.SetItemsList([]evermileApi.Item{
		item,
	},
	)

	quoteReq := *evermileApi.NewQuoteReq(
		[]string{cfg.Evermile.DefaultLocation},
		destinationLocations,
		[]evermileApi.Parcel{
			parcel,
		},
	)
	quoteReq.SetHandling([]evermileApi.HandlingInstruction{
		evermileApi.HANDLINGINSTRUCTION_CAKE,
		evermileApi.HANDLINGINSTRUCTION_PERISHABLE,
	})
	quoteReq.SetProofOfDeliveryRequirement([]evermileApi.ProofOfDeliveryRequirement{
		evermileApi.PROOFOFDELIVERYREQUIREMENT_PARCEL_PHOTO,
	})
	quoteReq.SetExcludedVehicleTypes([]evermileApi.VehicleType{
		evermileApi.VEHICLETYPE_PUSHBIKE,
		evermileApi.VEHICLETYPE_MOTORBIKE,
		evermileApi.VEHICLETYPE_CARGO_BIKE,
	})
	quoteReq.SetInstructions("Fragile! Handle with care!")

	resp, r, err := evermileClient.QuotesApi.QuotePost(evermileContext).QuoteReq(quoteReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.QuotePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QuotePost`: QuoteResponse
	fmt.Fprintf(os.Stdout, "Response from `QuotesApi.QuotePost`: %v\n", resp)

	dateProps, datePropsOk := resp.GetDateProposalsOk()
	if !datePropsOk {
		log.Println("No proposals at all")
		return
	}

	props, ok := dateProps[0].GetProposalsOk()
	if !ok {
		log.Println("No proposals for the day")
		return
	}

	var proposal evermileApi.Proposal

	for _, prop := range props {
		if prop.HasProposal() && prop.GetLabel() != "allDay" {
			proposal = prop.GetProposal()
			break
		}
	}

	log.Println(proposal)

	evermileOrder, eoError := executeOrderRequest(proposal, order)
	if eoError != nil {
		newOrder := goshopify.Order{
			ID:   order.ID,
			Tags: order.Tags + ", evermile",
			NoteAttributes: append(order.NoteAttributes, goshopify.NoteAttribute{
				Name:  "evermile-order-id",
				Value: evermileOrder.GetId(),
			}),
		}
		_, err = shopifyClient.Order.Update(newOrder)
		if err != nil {
			log.Println(err)
			return
		}

	}
}

func executeOrderRequest(proposal evermileApi.Proposal, order goshopify.Order) (*evermileApi.OrderDetails, error) {
	contactDetails := *evermileApi.NewContactDetails(
		order.ShippingAddress.Name,
		order.Email,
	)
	contactDetails.ContactEmail = order.Email
	contactDetails.ContactPhone = &order.ShippingAddress.Phone
	contactDetails.Instructions = &order.Note

	orderRequest := *evermileApi.NewOrderRequest(
		proposal.GetId(),
		contactDetails,
		true,
	)
	orderRequest.SetPickupLocationId(cfg.Evermile.DefaultLocation)
	orderRequest.SetExternalOrderId(strconv.FormatInt(int64(order.OrderNumber), 10))

	resp, r, err := evermileClient.OrdersApi.OrderPost(evermileContext).OrderRequest(orderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return nil, err
	}
	// response from `OrderPost`: OrderResponse
	// fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderPost`: %v\n", resp)
	return resp, nil
}

// func registerWebhooks() {
// 	// get all webhooks

// 	getWebhooksForTopic(evermileApi.WEBHOOKTOPIC_ORDER_CREATED)
// }

// func getWebhooksForTopic(topic evermileApi.WebhookTopic) {
// 	resp, r, err := evermileClient.WebhooksApi.WebhookGet(evermileContext).XEVERMILEAUTHHEADER("PUNKCAKE_AUTH_HEADER").XEVERMILEWEBHOOKTOPIC(topic).Execute()
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhookGet``: %v\n", err)
// 		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
// 		return
// 	}
// 	// response from `WebhookGet`: []Webhook
// 	fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WebhookGet`: %v\n", resp)
// }
