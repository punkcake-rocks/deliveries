# \WebhooksApi

All URIs are relative to *https://api.prod.evermile.io/v1/commercial*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WebhookGet**](WebhooksApi.md#WebhookGet) | **Get** /webhook | List all webhooks
[**WebhookPost**](WebhooksApi.md#WebhookPost) | **Post** /webhook | Register webhook
[**WebhookWebhookIdDelete**](WebhooksApi.md#WebhookWebhookIdDelete) | **Delete** /webhook/{webhookId} | Delete a webhook
[**WebhookWebhookIdGet**](WebhooksApi.md#WebhookWebhookIdGet) | **Get** /webhook/{webhookId} | Get webhook details



## WebhookGet

> WebhookGetResponse WebhookGet(ctx).XEVERMILEAUTHHEADER(xEVERMILEAUTHHEADER).XEVERMILEWEBHOOKTOPIC(xEVERMILEWEBHOOKTOPIC).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()

List all webhooks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    xEVERMILEAUTHHEADER := "xEVERMILEAUTHHEADER_example" // string | The auth header specified in the URL registration request (the actual header name is up to the user)
    xEVERMILEWEBHOOKTOPIC := openapiclient.webhookTopic("order.created") // WebhookTopic | The webhook topic this callback is referring to
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.WebhookGet(context.Background()).XEVERMILEAUTHHEADER(xEVERMILEAUTHHEADER).XEVERMILEWEBHOOKTOPIC(xEVERMILEWEBHOOKTOPIC).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhookGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhookGet`: WebhookGetResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WebhookGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEVERMILEAUTHHEADER** | **string** | The auth header specified in the URL registration request (the actual header name is up to the user) | 
 **xEVERMILEWEBHOOKTOPIC** | [**WebhookTopic**](WebhookTopic.md) | The webhook topic this callback is referring to | 
 **xEVERMILEMERCHANTID** | **string** | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. | 
 **xEVERMILETRACEID** | **string** | A trace ID for tracing the request through the Evermile platform | 

### Return type

[**WebhookGetResponse**](WebhookGetResponse.md)

### Authorization

[Prod](../README.md#Prod), [Sandbox](../README.md#Sandbox)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookPost

> WebhookPostResponse WebhookPost(ctx).XEVERMILEAUTHHEADER(xEVERMILEAUTHHEADER).XEVERMILEWEBHOOKTOPIC(xEVERMILEWEBHOOKTOPIC).WebhookRequest(webhookRequest).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()

Register webhook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    xEVERMILEAUTHHEADER := "xEVERMILEAUTHHEADER_example" // string | The auth header specified in the URL registration request (the actual header name is up to the user)
    xEVERMILEWEBHOOKTOPIC := openapiclient.webhookTopic("order.created") // WebhookTopic | The webhook topic this callback is referring to
    webhookRequest := *openapiclient.NewWebhookRequest("https://my.shop.com/webhook", "X-Auth-From-Evermile", "This is Me!") // WebhookRequest | 
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.WebhookPost(context.Background()).XEVERMILEAUTHHEADER(xEVERMILEAUTHHEADER).XEVERMILEWEBHOOKTOPIC(xEVERMILEWEBHOOKTOPIC).WebhookRequest(webhookRequest).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhookPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhookPost`: WebhookPostResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WebhookPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xEVERMILEAUTHHEADER** | **string** | The auth header specified in the URL registration request (the actual header name is up to the user) | 
 **xEVERMILEWEBHOOKTOPIC** | [**WebhookTopic**](WebhookTopic.md) | The webhook topic this callback is referring to | 
 **webhookRequest** | [**WebhookRequest**](WebhookRequest.md) |  | 
 **xEVERMILEMERCHANTID** | **string** | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. | 
 **xEVERMILETRACEID** | **string** | A trace ID for tracing the request through the Evermile platform | 

### Return type

[**WebhookPostResponse**](WebhookPostResponse.md)

### Authorization

[Prod](../README.md#Prod), [Sandbox](../README.md#Sandbox)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookWebhookIdDelete

> WebhookWebhookIdDelete(ctx, webhookId).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()

Delete a webhook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the webhook
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WebhooksApi.WebhookWebhookIdDelete(context.Background(), webhookId).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhookWebhookIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | The ID of the webhook | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookWebhookIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEVERMILEMERCHANTID** | **string** | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. | 
 **xEVERMILETRACEID** | **string** | A trace ID for tracing the request through the Evermile platform | 

### Return type

 (empty response body)

### Authorization

[Prod](../README.md#Prod), [Sandbox](../README.md#Sandbox)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookWebhookIdGet

> Webhook WebhookWebhookIdGet(ctx, webhookId).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()

Get webhook details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    webhookId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the webhook
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.WebhookWebhookIdGet(context.Background(), webhookId).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.WebhookWebhookIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhookWebhookIdGet`: Webhook
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.WebhookWebhookIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** | The ID of the webhook | 

### Other Parameters

Other parameters are passed through a pointer to a apiWebhookWebhookIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEVERMILEMERCHANTID** | **string** | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. | 
 **xEVERMILETRACEID** | **string** | A trace ID for tracing the request through the Evermile platform | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[Prod](../README.md#Prod), [Sandbox](../README.md#Sandbox)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

