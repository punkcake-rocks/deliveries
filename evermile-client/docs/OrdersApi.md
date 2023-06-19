# \OrdersApi

All URIs are relative to *https://api.prod.evermile.io/v1/commercial*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrderOrderIdDelete**](OrdersApi.md#OrderOrderIdDelete) | **Delete** /order/{orderId} | Cancel an order
[**OrderOrderIdGet**](OrdersApi.md#OrderOrderIdGet) | **Get** /order/{orderId} | Fetch order details
[**OrderOrderIdLabelGet**](OrdersApi.md#OrderOrderIdLabelGet) | **Get** /order/{orderId}/label | Create a label for an order
[**OrderOrderIdLiveTrackingGet**](OrdersApi.md#OrderOrderIdLiveTrackingGet) | **Get** /order/{orderId}/liveTracking | Fetch live order tracking
[**OrderOrderIdPatch**](OrdersApi.md#OrderOrderIdPatch) | **Patch** /order/{orderId} | Update order details
[**OrderPost**](OrdersApi.md#OrderPost) | **Post** /order | Order a delivery
[**OrdersGet**](OrdersApi.md#OrdersGet) | **Get** /orders | Get orders



## OrderOrderIdDelete

> OrderOrderIdDelete(ctx, orderId).NoConfirmIfNoFee(noConfirmIfNoFee).XEvermileToken(xEvermileToken).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()

Cancel an order



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
    orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the order
    noConfirmIfNoFee := true // bool | Should the order be cancelled immediately if there is no cancellation fee. Default is false. (optional)
    xEvermileToken := "xEvermileToken_example" // string | A token to confirm cancellation after receiving 402 code (optional)
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrdersApi.OrderOrderIdDelete(context.Background(), orderId).NoConfirmIfNoFee(noConfirmIfNoFee).XEvermileToken(xEvermileToken).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderOrderIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | The ID of the order | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderOrderIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noConfirmIfNoFee** | **bool** | Should the order be cancelled immediately if there is no cancellation fee. Default is false. | 
 **xEvermileToken** | **string** | A token to confirm cancellation after receiving 402 code | 
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


## OrderOrderIdGet

> OrderDetails OrderOrderIdGet(ctx, orderId).IncludeTrackingInfo(includeTrackingInfo).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()

Fetch order details



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
    orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the order
    includeTrackingInfo := true // bool | Should include live tracking info with this order. Default is false. (optional)
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.OrderOrderIdGet(context.Background(), orderId).IncludeTrackingInfo(includeTrackingInfo).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderOrderIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderOrderIdGet`: OrderDetails
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderOrderIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | The ID of the order | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderOrderIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeTrackingInfo** | **bool** | Should include live tracking info with this order. Default is false. | 
 **xEVERMILEMERCHANTID** | **string** | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. | 
 **xEVERMILETRACEID** | **string** | A trace ID for tracing the request through the Evermile platform | 

### Return type

[**OrderDetails**](OrderDetails.md)

### Authorization

[Prod](../README.md#Prod), [Sandbox](../README.md#Sandbox)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderOrderIdLabelGet

> *os.File OrderOrderIdLabelGet(ctx, orderId).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()

Create a label for an order



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
    orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the order
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.OrderOrderIdLabelGet(context.Background(), orderId).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderOrderIdLabelGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderOrderIdLabelGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderOrderIdLabelGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | The ID of the order | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderOrderIdLabelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEVERMILEMERCHANTID** | **string** | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. | 
 **xEVERMILETRACEID** | **string** | A trace ID for tracing the request through the Evermile platform | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[Prod](../README.md#Prod), [Sandbox](../README.md#Sandbox)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderOrderIdLiveTrackingGet

> OrderTrackingInfo OrderOrderIdLiveTrackingGet(ctx, orderId).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()

Fetch live order tracking



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
    orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the order
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.OrderOrderIdLiveTrackingGet(context.Background(), orderId).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderOrderIdLiveTrackingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderOrderIdLiveTrackingGet`: OrderTrackingInfo
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderOrderIdLiveTrackingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | The ID of the order | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderOrderIdLiveTrackingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEVERMILEMERCHANTID** | **string** | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. | 
 **xEVERMILETRACEID** | **string** | A trace ID for tracing the request through the Evermile platform | 

### Return type

[**OrderTrackingInfo**](OrderTrackingInfo.md)

### Authorization

[Prod](../README.md#Prod), [Sandbox](../README.md#Sandbox)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderOrderIdPatch

> OrderOrderIdPatch(ctx, orderId).OrderOrderIdPatchRequest(orderOrderIdPatchRequest).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()

Update order details



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
    orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the order
    orderOrderIdPatchRequest := *openapiclient.NewOrderOrderIdPatchRequest(openapiclient.orderStatus("Pending")) // OrderOrderIdPatchRequest | 
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OrdersApi.OrderOrderIdPatch(context.Background(), orderId).OrderOrderIdPatchRequest(orderOrderIdPatchRequest).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderOrderIdPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | The ID of the order | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderOrderIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderOrderIdPatchRequest** | [**OrderOrderIdPatchRequest**](OrderOrderIdPatchRequest.md) |  | 
 **xEVERMILEMERCHANTID** | **string** | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. | 
 **xEVERMILETRACEID** | **string** | A trace ID for tracing the request through the Evermile platform | 

### Return type

 (empty response body)

### Authorization

[Prod](../README.md#Prod), [Sandbox](../README.md#Sandbox)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrderPost

> OrderDetails OrderPost(ctx).OrderRequest(orderRequest).XEvermileToken(xEvermileToken).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).XEVERMILESTOREID(xEVERMILESTOREID).Execute()

Order a delivery



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
    orderRequest := *openapiclient.NewOrderRequest("f9bc2fa1-fa33-4b79-8fe5-f4fbe24c6ecc", *openapiclient.NewContactDetails("John Smith", "johnsmith@email.com"), false) // OrderRequest | 
    xEvermileToken := "xEvermileToken_example" // string | A token to confirm cancellation after receiving 402 code (optional)
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)
    xEVERMILESTOREID := "xEVERMILESTOREID_example" // string | A store ID for an order's store platform context (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.OrderPost(context.Background()).OrderRequest(orderRequest).XEvermileToken(xEvermileToken).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).XEVERMILESTOREID(xEVERMILESTOREID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrderPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrderPost`: OrderDetails
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderRequest** | [**OrderRequest**](OrderRequest.md) |  | 
 **xEvermileToken** | **string** | A token to confirm cancellation after receiving 402 code | 
 **xEVERMILEMERCHANTID** | **string** | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. | 
 **xEVERMILETRACEID** | **string** | A trace ID for tracing the request through the Evermile platform | 
 **xEVERMILESTOREID** | **string** | A store ID for an order&#39;s store platform context | 

### Return type

[**OrderDetails**](OrderDetails.md)

### Authorization

[Prod](../README.md#Prod), [Sandbox](../README.md#Sandbox)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrdersGet

> []OrderListItem OrdersGet(ctx).From(from).To(to).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).XEVERMILESTOREID(xEVERMILESTOREID).Execute()

Get orders



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    from := time.Now() // string | The start date to query (YYYY-MM-DD). Default is today. (optional)
    to := time.Now() // string | The end date to query (YYYY-MM-DD). Default is today. (optional)
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)
    xEVERMILESTOREID := "xEVERMILESTOREID_example" // string | A store ID for an order's store platform context (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.OrdersGet(context.Background()).From(from).To(to).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).XEVERMILESTOREID(xEVERMILESTOREID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.OrdersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrdersGet`: []OrderListItem
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.OrdersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrdersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The start date to query (YYYY-MM-DD). Default is today. | 
 **to** | **string** | The end date to query (YYYY-MM-DD). Default is today. | 
 **xEVERMILEMERCHANTID** | **string** | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. | 
 **xEVERMILETRACEID** | **string** | A trace ID for tracing the request through the Evermile platform | 
 **xEVERMILESTOREID** | **string** | A store ID for an order&#39;s store platform context | 

### Return type

[**[]OrderListItem**](OrderListItem.md)

### Authorization

[Prod](../README.md#Prod), [Sandbox](../README.md#Sandbox)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

