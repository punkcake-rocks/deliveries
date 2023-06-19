# OrderOrderIdDelete402Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payment** | Pointer to [**Price**](Price.md) |  | [optional] 
**Token** | Pointer to **string** | A token that should be returned in order to confirm the cancellation. To confirm the cancellation, pass this in an X-EVERMILE-TOKEN header. | [optional] 

## Methods

### NewOrderOrderIdDelete402Response

`func NewOrderOrderIdDelete402Response() *OrderOrderIdDelete402Response`

NewOrderOrderIdDelete402Response instantiates a new OrderOrderIdDelete402Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderOrderIdDelete402ResponseWithDefaults

`func NewOrderOrderIdDelete402ResponseWithDefaults() *OrderOrderIdDelete402Response`

NewOrderOrderIdDelete402ResponseWithDefaults instantiates a new OrderOrderIdDelete402Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayment

`func (o *OrderOrderIdDelete402Response) GetPayment() Price`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *OrderOrderIdDelete402Response) GetPaymentOk() (*Price, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *OrderOrderIdDelete402Response) SetPayment(v Price)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *OrderOrderIdDelete402Response) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetToken

`func (o *OrderOrderIdDelete402Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OrderOrderIdDelete402Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OrderOrderIdDelete402Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *OrderOrderIdDelete402Response) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


