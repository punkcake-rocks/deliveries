# \LocationsApi

All URIs are relative to *https://api.prod.evermile.io/v1/commercial*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SenderInfoIdDelete**](LocationsApi.md#SenderInfoIdDelete) | **Delete** /sender_info/{id} | Delete a delivery location
[**SenderInfoPut**](LocationsApi.md#SenderInfoPut) | **Put** /sender_info | Create or update a delivery location
[**SenderInfosGet**](LocationsApi.md#SenderInfosGet) | **Get** /sender_infos | Get sender and location infos



## SenderInfoIdDelete

> SenderInfoIdDelete(ctx, id).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()

Delete a delivery location

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
    id := "f9bc2fa1-fa33-4b79-8fe5-f4fbe24c6eaa" // string | The sender info ID
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.LocationsApi.SenderInfoIdDelete(context.Background(), id).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.SenderInfoIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The sender info ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSenderInfoIdDeleteRequest struct via the builder pattern


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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SenderInfoPut

> LocationResp SenderInfoPut(ctx).Type_(type_).Body(body).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).XEVERMILESTOREID(xEVERMILESTOREID).Execute()

Create or update a delivery location

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
    type_ := openapiclient.locationTypeEnum("pickup") // LocationTypeEnum | Whether this action refers to pickup or sender info (pickup/info) (default to "pickup")
    body := Location1(987) // Location1 | 
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)
    xEVERMILESTOREID := "xEVERMILESTOREID_example" // string | A store ID for an order's store platform context (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.SenderInfoPut(context.Background()).Type_(type_).Body(body).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).XEVERMILESTOREID(xEVERMILESTOREID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.SenderInfoPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SenderInfoPut`: LocationResp
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.SenderInfoPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSenderInfoPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**LocationTypeEnum**](LocationTypeEnum.md) | Whether this action refers to pickup or sender info (pickup/info) | [default to &quot;pickup&quot;]
 **body** | **Location1** |  | 
 **xEVERMILEMERCHANTID** | **string** | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. | 
 **xEVERMILETRACEID** | **string** | A trace ID for tracing the request through the Evermile platform | 
 **xEVERMILESTOREID** | **string** | A store ID for an order&#39;s store platform context | 

### Return type

[**LocationResp**](LocationResp.md)

### Authorization

[Prod](../README.md#Prod), [Sandbox](../README.md#Sandbox)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SenderInfosGet

> LocationsResponse SenderInfosGet(ctx).Type_(type_).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).XEVERMILESTOREID(xEVERMILESTOREID).Execute()

Get sender and location infos



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
    type_ := openapiclient.locationTypeEnum("pickup") // LocationTypeEnum | Whether this action refers to pickup or sender info (pickup/info) (default to "pickup")
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)
    xEVERMILESTOREID := "xEVERMILESTOREID_example" // string | A store ID for an order's store platform context (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.SenderInfosGet(context.Background()).Type_(type_).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).XEVERMILESTOREID(xEVERMILESTOREID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.SenderInfosGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SenderInfosGet`: LocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.SenderInfosGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSenderInfosGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | [**LocationTypeEnum**](LocationTypeEnum.md) | Whether this action refers to pickup or sender info (pickup/info) | [default to &quot;pickup&quot;]
 **xEVERMILEMERCHANTID** | **string** | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. | 
 **xEVERMILETRACEID** | **string** | A trace ID for tracing the request through the Evermile platform | 
 **xEVERMILESTOREID** | **string** | A store ID for an order&#39;s store platform context | 

### Return type

[**LocationsResponse**](LocationsResponse.md)

### Authorization

[Prod](../README.md#Prod), [Sandbox](../README.md#Sandbox)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

