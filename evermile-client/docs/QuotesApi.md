# \QuotesApi

All URIs are relative to *https://api.prod.evermile.io/v1/commercial*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProposalProposalIdGet**](QuotesApi.md#ProposalProposalIdGet) | **Get** /proposal/{proposalId} | Retrieve a proposal
[**QuotePost**](QuotesApi.md#QuotePost) | **Post** /quote | Get a quote



## ProposalProposalIdGet

> ProposalResponse ProposalProposalIdGet(ctx, proposalId).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()

Retrieve a proposal



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
    proposalId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the proposal
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotesApi.ProposalProposalIdGet(context.Background(), proposalId).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.ProposalProposalIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProposalProposalIdGet`: ProposalResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.ProposalProposalIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**proposalId** | **string** | The ID of the proposal | 

### Other Parameters

Other parameters are passed through a pointer to a apiProposalProposalIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xEVERMILEMERCHANTID** | **string** | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. | 
 **xEVERMILETRACEID** | **string** | A trace ID for tracing the request through the Evermile platform | 

### Return type

[**ProposalResponse**](ProposalResponse.md)

### Authorization

[Prod](../README.md#Prod), [Sandbox](../README.md#Sandbox)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QuotePost

> QuoteResponse QuotePost(ctx).QuoteReq(quoteReq).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).XEVERMILESTOREID(xEVERMILESTOREID).Execute()

Get a quote



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
    quoteReq := *openapiclient.NewQuoteReq([]string{"f9bc2fa1-fa33-4b79-8fe5-f4fbe24c6ecc"}, []openapiclient.DestinationLocation{*openapiclient.NewDestinationLocation()}, []openapiclient.Parcel{*openapiclient.NewParcel(*openapiclient.NewDimensions(float32(10.25), float32(8.25), float32(5.1)), float32(2.3), openapiclient.parcelType("Package"))}) // QuoteReq | 
    xEVERMILEMERCHANTID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)
    xEVERMILESTOREID := "xEVERMILESTOREID_example" // string | A store ID for an order's store platform context (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QuotesApi.QuotePost(context.Background()).QuoteReq(quoteReq).XEVERMILEMERCHANTID(xEVERMILEMERCHANTID).XEVERMILETRACEID(xEVERMILETRACEID).XEVERMILESTOREID(xEVERMILESTOREID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotesApi.QuotePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QuotePost`: QuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `QuotesApi.QuotePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQuotePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **quoteReq** | [**QuoteReq**](QuoteReq.md) |  | 
 **xEVERMILEMERCHANTID** | **string** | The merchant ID, if using a client credentials token. Will be ignored with a regular user token. | 
 **xEVERMILETRACEID** | **string** | A trace ID for tracing the request through the Evermile platform | 
 **xEVERMILESTOREID** | **string** | A store ID for an order&#39;s store platform context | 

### Return type

[**QuoteResponse**](QuoteResponse.md)

### Authorization

[Prod](../README.md#Prod), [Sandbox](../README.md#Sandbox)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

