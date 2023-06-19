# OrderCollectedCallbackResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | The time this callback was sent in ISO8601 format | 
**OrderId** | **string** | The ID of the order | 

## Methods

### NewOrderCollectedCallbackResp

`func NewOrderCollectedCallbackResp(timestamp time.Time, orderId string, ) *OrderCollectedCallbackResp`

NewOrderCollectedCallbackResp instantiates a new OrderCollectedCallbackResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderCollectedCallbackRespWithDefaults

`func NewOrderCollectedCallbackRespWithDefaults() *OrderCollectedCallbackResp`

NewOrderCollectedCallbackRespWithDefaults instantiates a new OrderCollectedCallbackResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *OrderCollectedCallbackResp) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OrderCollectedCallbackResp) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OrderCollectedCallbackResp) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetOrderId

`func (o *OrderCollectedCallbackResp) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderCollectedCallbackResp) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderCollectedCallbackResp) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


