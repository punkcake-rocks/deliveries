# OrderCancelledCallbackResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | The time this callback was sent in ISO8601 format | 
**OrderId** | **string** | The ID of the order | 
**CancellationFee** | Pointer to [**Price1**](Price1.md) |  | [optional] 

## Methods

### NewOrderCancelledCallbackResp

`func NewOrderCancelledCallbackResp(timestamp time.Time, orderId string, ) *OrderCancelledCallbackResp`

NewOrderCancelledCallbackResp instantiates a new OrderCancelledCallbackResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCancelledCallbackRespWithDefaults

`func NewOrderCancelledCallbackRespWithDefaults() *OrderCancelledCallbackResp`

NewOrderCancelledCallbackRespWithDefaults instantiates a new OrderCancelledCallbackResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *OrderCancelledCallbackResp) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OrderCancelledCallbackResp) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OrderCancelledCallbackResp) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetOrderId

`func (o *OrderCancelledCallbackResp) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderCancelledCallbackResp) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderCancelledCallbackResp) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetCancellationFee

`func (o *OrderCancelledCallbackResp) GetCancellationFee() Price1`

GetCancellationFee returns the CancellationFee field if non-nil, zero value otherwise.

### GetCancellationFeeOk

`func (o *OrderCancelledCallbackResp) GetCancellationFeeOk() (*Price1, bool)`

GetCancellationFeeOk returns a tuple with the CancellationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationFee

`func (o *OrderCancelledCallbackResp) SetCancellationFee(v Price1)`

SetCancellationFee sets CancellationFee field to given value.

### HasCancellationFee

`func (o *OrderCancelledCallbackResp) HasCancellationFee() bool`

HasCancellationFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


