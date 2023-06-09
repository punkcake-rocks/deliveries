# Go API client for EvermileClient

This is Evermile's commercial API for handling delivery quotes and orders

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://www.evermile.io/support](https://www.evermile.io/support)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import EvermileClient "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), EvermileClient.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), EvermileClient.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), EvermileClient.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), EvermileClient.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.prod.evermile.io/v1/commercial*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*LocationsApi* | [**SenderInfoIdDelete**](docs/LocationsApi.md#senderinfoiddelete) | **Delete** /sender_info/{id} | Delete a delivery location
*LocationsApi* | [**SenderInfoPut**](docs/LocationsApi.md#senderinfoput) | **Put** /sender_info | Create or update a delivery location
*LocationsApi* | [**SenderInfosGet**](docs/LocationsApi.md#senderinfosget) | **Get** /sender_infos | Get sender and location infos
*OrdersApi* | [**OrderOrderIdDelete**](docs/OrdersApi.md#orderorderiddelete) | **Delete** /order/{orderId} | Cancel an order
*OrdersApi* | [**OrderOrderIdGet**](docs/OrdersApi.md#orderorderidget) | **Get** /order/{orderId} | Fetch order details
*OrdersApi* | [**OrderOrderIdLabelGet**](docs/OrdersApi.md#orderorderidlabelget) | **Get** /order/{orderId}/label | Create a label for an order
*OrdersApi* | [**OrderOrderIdLiveTrackingGet**](docs/OrdersApi.md#orderorderidlivetrackingget) | **Get** /order/{orderId}/liveTracking | Fetch live order tracking
*OrdersApi* | [**OrderOrderIdPatch**](docs/OrdersApi.md#orderorderidpatch) | **Patch** /order/{orderId} | Update order details
*OrdersApi* | [**OrderPost**](docs/OrdersApi.md#orderpost) | **Post** /order | Order a delivery
*OrdersApi* | [**OrdersGet**](docs/OrdersApi.md#ordersget) | **Get** /orders | Get orders
*QuotesApi* | [**ProposalProposalIdGet**](docs/QuotesApi.md#proposalproposalidget) | **Get** /proposal/{proposalId} | Retrieve a proposal
*QuotesApi* | [**QuotePost**](docs/QuotesApi.md#quotepost) | **Post** /quote | Get a quote
*ValidationsApi* | [**AddressValidateGet**](docs/ValidationsApi.md#addressvalidateget) | **Get** /address/validate | Validate address
*ValidationsApi* | [**EmailValidateGet**](docs/ValidationsApi.md#emailvalidateget) | **Get** /email/validate | Validate email
*ValidationsApi* | [**PhoneValidateGet**](docs/ValidationsApi.md#phonevalidateget) | **Get** /phone/validate | Validate phone number
*WebhooksApi* | [**WebhookGet**](docs/WebhooksApi.md#webhookget) | **Get** /webhook | List all webhooks
*WebhooksApi* | [**WebhookPost**](docs/WebhooksApi.md#webhookpost) | **Post** /webhook | Register webhook
*WebhooksApi* | [**WebhookWebhookIdDelete**](docs/WebhooksApi.md#webhookwebhookiddelete) | **Delete** /webhook/{webhookId} | Delete a webhook
*WebhooksApi* | [**WebhookWebhookIdGet**](docs/WebhooksApi.md#webhookwebhookidget) | **Get** /webhook/{webhookId} | Get webhook details


## Documentation For Models

 - [Address](docs/Address.md)
 - [Address1](docs/Address1.md)
 - [Address2](docs/Address2.md)
 - [AddressComponents](docs/AddressComponents.md)
 - [AddressType](docs/AddressType.md)
 - [CarriersType](docs/CarriersType.md)
 - [ContactDetails](docs/ContactDetails.md)
 - [Customer](docs/Customer.md)
 - [DateProposals](docs/DateProposals.md)
 - [DeliveryLiveUpdateCallbackResp](docs/DeliveryLiveUpdateCallbackResp.md)
 - [DeliverySlot](docs/DeliverySlot.md)
 - [DestinationLocation](docs/DestinationLocation.md)
 - [Dimensions](docs/Dimensions.md)
 - [Discount](docs/Discount.md)
 - [DropoffWindowTypeEnum](docs/DropoffWindowTypeEnum.md)
 - [GeoLocation](docs/GeoLocation.md)
 - [HandlingInstruction](docs/HandlingInstruction.md)
 - [HandoffContactInfo](docs/HandoffContactInfo.md)
 - [HandoffContactInfo1](docs/HandoffContactInfo1.md)
 - [HandoffInfo](docs/HandoffInfo.md)
 - [HandoffInfo1](docs/HandoffInfo1.md)
 - [HandoffType](docs/HandoffType.md)
 - [InvalidAddressBody](docs/InvalidAddressBody.md)
 - [Item](docs/Item.md)
 - [Item1](docs/Item1.md)
 - [Location](docs/Location.md)
 - [Location1](docs/Location1.md)
 - [LocationResp](docs/LocationResp.md)
 - [LocationTimeWindow](docs/LocationTimeWindow.md)
 - [LocationTypeEnum](docs/LocationTypeEnum.md)
 - [LocationWeeklySchedule](docs/LocationWeeklySchedule.md)
 - [LocationWeeklySchedule1](docs/LocationWeeklySchedule1.md)
 - [LocationWithId](docs/LocationWithId.md)
 - [LocationWithIdAllOf](docs/LocationWithIdAllOf.md)
 - [LocationsResponse](docs/LocationsResponse.md)
 - [OptionalTimeWindow](docs/OptionalTimeWindow.md)
 - [OrderCancelledCallbackResp](docs/OrderCancelledCallbackResp.md)
 - [OrderCollectedCallbackResp](docs/OrderCollectedCallbackResp.md)
 - [OrderCreatedCallbackResp](docs/OrderCreatedCallbackResp.md)
 - [OrderDeliveredCallbackResp](docs/OrderDeliveredCallbackResp.md)
 - [OrderDetails](docs/OrderDetails.md)
 - [OrderDetails1](docs/OrderDetails1.md)
 - [OrderExtendedStatus](docs/OrderExtendedStatus.md)
 - [OrderExtendedStatus1](docs/OrderExtendedStatus1.md)
 - [OrderFailedCallbackResp](docs/OrderFailedCallbackResp.md)
 - [OrderListItem](docs/OrderListItem.md)
 - [OrderOrderIdDelete402Response](docs/OrderOrderIdDelete402Response.md)
 - [OrderOrderIdPatchRequest](docs/OrderOrderIdPatchRequest.md)
 - [OrderRequest](docs/OrderRequest.md)
 - [OrderStatus](docs/OrderStatus.md)
 - [OrderTrackingInfo](docs/OrderTrackingInfo.md)
 - [OrderTrackingInfo1](docs/OrderTrackingInfo1.md)
 - [Parcel](docs/Parcel.md)
 - [Parcel1](docs/Parcel1.md)
 - [ParcelType](docs/ParcelType.md)
 - [Price](docs/Price.md)
 - [Price1](docs/Price1.md)
 - [ProofOfDelivery](docs/ProofOfDelivery.md)
 - [ProofOfDeliveryRequirement](docs/ProofOfDeliveryRequirement.md)
 - [Proposal](docs/Proposal.md)
 - [ProposalLabel](docs/ProposalLabel.md)
 - [ProposalLabelEnum](docs/ProposalLabelEnum.md)
 - [ProposalResponse](docs/ProposalResponse.md)
 - [ProposalType](docs/ProposalType.md)
 - [QuoteReq](docs/QuoteReq.md)
 - [QuoteResponse](docs/QuoteResponse.md)
 - [TimeWindow](docs/TimeWindow.md)
 - [TimeslotProposal](docs/TimeslotProposal.md)
 - [VehicleType](docs/VehicleType.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookGetResponse](docs/WebhookGetResponse.md)
 - [WebhookPostResponse](docs/WebhookPostResponse.md)
 - [WebhookRequest](docs/WebhookRequest.md)
 - [WebhookTopic](docs/WebhookTopic.md)


## Documentation For Authorization



### Sandbox


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **order**: Order deliveries
 - **read**: Get quotes and read information about deliveries

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


### Prod


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **order**: Order deliveries
 - **read**: Get quotes and read information about deliveries

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@evermile.io

