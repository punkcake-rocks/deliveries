# OrderCreatedCallbackResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | The time this callback was sent in ISO8601 format | 
**Order** | [**OrderDetails1**](OrderDetails1.md) |  | 

## Methods

### NewOrderCreatedCallbackResp

`func NewOrderCreatedCallbackResp(timestamp time.Time, order OrderDetails1, ) *OrderCreatedCallbackResp`

NewOrderCreatedCallbackResp instantiates a new OrderCreatedCallbackResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCreatedCallbackRespWithDefaults

`func NewOrderCreatedCallbackRespWithDefaults() *OrderCreatedCallbackResp`

NewOrderCreatedCallbackRespWithDefaults instantiates a new OrderCreatedCallbackResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *OrderCreatedCallbackResp) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OrderCreatedCallbackResp) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OrderCreatedCallbackResp) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetOrder

`func (o *OrderCreatedCallbackResp) GetOrder() OrderDetails1`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OrderCreatedCallbackResp) GetOrderOk() (*OrderDetails1, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OrderCreatedCallbackResp) SetOrder(v OrderDetails1)`

SetOrder sets Order field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


