# OrderFailedCallbackResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | The time this callback was sent in ISO8601 format | 
**OrderId** | **string** | The ID of the order | 
**FailureReason** | Pointer to **string** | An optional string detailing the failure reason | [optional] 

## Methods

### NewOrderFailedCallbackResp

`func NewOrderFailedCallbackResp(timestamp time.Time, orderId string, ) *OrderFailedCallbackResp`

NewOrderFailedCallbackResp instantiates a new OrderFailedCallbackResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderFailedCallbackRespWithDefaults

`func NewOrderFailedCallbackRespWithDefaults() *OrderFailedCallbackResp`

NewOrderFailedCallbackRespWithDefaults instantiates a new OrderFailedCallbackResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *OrderFailedCallbackResp) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OrderFailedCallbackResp) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OrderFailedCallbackResp) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetOrderId

`func (o *OrderFailedCallbackResp) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderFailedCallbackResp) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderFailedCallbackResp) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetFailureReason

`func (o *OrderFailedCallbackResp) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *OrderFailedCallbackResp) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *OrderFailedCallbackResp) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *OrderFailedCallbackResp) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


