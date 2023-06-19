# \ValidationsApi

All URIs are relative to *https://api.prod.evermile.io/v1/commercial*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressValidateGet**](ValidationsApi.md#AddressValidateGet) | **Get** /address/validate | Validate address
[**EmailValidateGet**](ValidationsApi.md#EmailValidateGet) | **Get** /email/validate | Validate email
[**PhoneValidateGet**](ValidationsApi.md#PhoneValidateGet) | **Get** /phone/validate | Validate phone number



## AddressValidateGet

> AddressValidateGet(ctx).Street(street).City(city).Postcode(postcode).Country(country).Coordinates(coordinates).XEVERMILETRACEID(xEVERMILETRACEID).Execute()

Validate address



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
    street := "Oxford Rd." // string | The street for delivery
    city := "London" // string | The city for delivery
    postcode := "W1S 3PT" // string | The postal code for delivery
    country := "UK" // string | 2-letter code for the country for delivery (optional)
    coordinates := "51.528596197212174,-0.1199161745400151" // string | The (lat, lng) coordinates. Not required, but will be used if provided. (optional)
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ValidationsApi.AddressValidateGet(context.Background()).Street(street).City(city).Postcode(postcode).Country(country).Coordinates(coordinates).XEVERMILETRACEID(xEVERMILETRACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ValidationsApi.AddressValidateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddressValidateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **street** | **string** | The street for delivery | 
 **city** | **string** | The city for delivery | 
 **postcode** | **string** | The postal code for delivery | 
 **country** | **string** | 2-letter code for the country for delivery | 
 **coordinates** | **string** | The (lat, lng) coordinates. Not required, but will be used if provided. | 
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


## EmailValidateGet

> EmailValidateGet(ctx).Email(email).XEVERMILETRACEID(xEVERMILETRACEID).Execute()

Validate email



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
    email := "customer@evermile.io" // string | The email to validate
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ValidationsApi.EmailValidateGet(context.Background()).Email(email).XEVERMILETRACEID(xEVERMILETRACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ValidationsApi.EmailValidateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmailValidateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string** | The email to validate | 
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


## PhoneValidateGet

> PhoneValidateGet(ctx).Phone(phone).XEVERMILETRACEID(xEVERMILETRACEID).Execute()

Validate phone number



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
    phone := "441242240310" // string | The phone number to validate
    xEVERMILETRACEID := "xEVERMILETRACEID_example" // string | A trace ID for tracing the request through the Evermile platform (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ValidationsApi.PhoneValidateGet(context.Background()).Phone(phone).XEVERMILETRACEID(xEVERMILETRACEID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ValidationsApi.PhoneValidateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPhoneValidateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phone** | **string** | The phone number to validate | 
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

