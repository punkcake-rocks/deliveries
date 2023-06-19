# OrderDeliveredCallbackResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | The time this callback was sent in ISO8601 format | 
**OrderId** | **string** | The ID of the order | 

## Methods

### NewOrderDeliveredCallbackResp

`func NewOrderDeliveredCallbackResp(timestamp time.Time, orderId string, ) *OrderDeliveredCallbackResp`

NewOrderDeliveredCallbackResp instantiates a new OrderDeliveredCallbackResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderDeliveredCallbackRespWithDefaults

`func NewOrderDeliveredCallbackRespWithDefaults() *OrderDeliveredCallbackResp`

NewOrderDeliveredCallbackRespWithDefaults instantiates a new OrderDeliveredCallbackResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *OrderDeliveredCallbackResp) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OrderDeliveredCallbackResp) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OrderDeliveredCallbackResp) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetOrderId

`func (o *OrderDeliveredCallbackResp) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderDeliveredCallbackResp) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderDeliveredCallbackResp) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


