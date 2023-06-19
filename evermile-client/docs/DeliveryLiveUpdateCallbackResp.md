# DeliveryLiveUpdateCallbackResp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** | The time this callback was sent in ISO8601 format | 
**OrderId** | **string** | The ID of the order | 
**Eta** | Pointer to **time.Time** | The expected delivery time | [optional] 
**CourierLocation** | Pointer to [**GeoLocation**](GeoLocation.md) |  | [optional] 

## Methods

### NewDeliveryLiveUpdateCallbackResp

`func NewDeliveryLiveUpdateCallbackResp(timestamp time.Time, orderId string, ) *DeliveryLiveUpdateCallbackResp`

NewDeliveryLiveUpdateCallbackResp instantiates a new DeliveryLiveUpdateCallbackResp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryLiveUpdateCallbackRespWithDefaults

`func NewDeliveryLiveUpdateCallbackRespWithDefaults() *DeliveryLiveUpdateCallbackResp`

NewDeliveryLiveUpdateCallbackRespWithDefaults instantiates a new DeliveryLiveUpdateCallbackResp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *DeliveryLiveUpdateCallbackResp) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DeliveryLiveUpdateCallbackResp) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DeliveryLiveUpdateCallbackResp) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetOrderId

`func (o *DeliveryLiveUpdateCallbackResp) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *DeliveryLiveUpdateCallbackResp) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *DeliveryLiveUpdateCallbackResp) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetEta

`func (o *DeliveryLiveUpdateCallbackResp) GetEta() time.Time`

GetEta returns the Eta field if non-nil, zero value otherwise.

### GetEtaOk

`func (o *DeliveryLiveUpdateCallbackResp) GetEtaOk() (*time.Time, bool)`

GetEtaOk returns a tuple with the Eta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEta

`func (o *DeliveryLiveUpdateCallbackResp) SetEta(v time.Time)`

SetEta sets Eta field to given value.

### HasEta

`func (o *DeliveryLiveUpdateCallbackResp) HasEta() bool`

HasEta returns a boolean if a field has been set.

### GetCourierLocation

`func (o *DeliveryLiveUpdateCallbackResp) GetCourierLocation() GeoLocation`

GetCourierLocation returns the CourierLocation field if non-nil, zero value otherwise.

### GetCourierLocationOk

`func (o *DeliveryLiveUpdateCallbackResp) GetCourierLocationOk() (*GeoLocation, bool)`

GetCourierLocationOk returns a tuple with the CourierLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourierLocation

`func (o *DeliveryLiveUpdateCallbackResp) SetCourierLocation(v GeoLocation)`

SetCourierLocation sets CourierLocation field to given value.

### HasCourierLocation

`func (o *DeliveryLiveUpdateCallbackResp) HasCourierLocation() bool`

HasCourierLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


