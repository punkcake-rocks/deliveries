# OrderTrackingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CourierLocation** | Pointer to [**GeoLocation**](GeoLocation.md) |  | [optional] 
**DropoffEta** | Pointer to **time.Time** | The dropoff ETA in ISO8601 format | [optional] 
**PickupEta** | Pointer to **time.Time** | The pickup ETA in ISO8601 format | [optional] 
**CourierName** | Pointer to **string** | The courier full name | [optional] 
**CourierPhone** | Pointer to **string** | The courier phone number | [optional] 
**DropoffWindowStart** | Pointer to **time.Time** | The updated dropoff widnow start in ISO8601 format | [optional] 
**DropoffWindowEnd** | Pointer to **time.Time** | The updated dropoff window end in ISO8601 format | [optional] 
**PickupWindowStart** | Pointer to **time.Time** | The updated pickup widnow start in ISO8601 format | [optional] 
**PickupWindowEnd** | Pointer to **time.Time** | The updated pickup window end in ISO8601 format | [optional] 
**IsCourierAssigned** | **bool** | Indicate whether a courier is assigned for this order | 
**VehicleType** | Pointer to [**VehicleType**](VehicleType.md) |  | [optional] 

## Methods

### NewOrderTrackingInfo

`func NewOrderTrackingInfo(isCourierAssigned bool, ) *OrderTrackingInfo`

NewOrderTrackingInfo instantiates a new OrderTrackingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderTrackingInfoWithDefaults

`func NewOrderTrackingInfoWithDefaults() *OrderTrackingInfo`

NewOrderTrackingInfoWithDefaults instantiates a new OrderTrackingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCourierLocation

`func (o *OrderTrackingInfo) GetCourierLocation() GeoLocation`

GetCourierLocation returns the CourierLocation field if non-nil, zero value otherwise.

### GetCourierLocationOk

`func (o *OrderTrackingInfo) GetCourierLocationOk() (*GeoLocation, bool)`

GetCourierLocationOk returns a tuple with the CourierLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourierLocation

`func (o *OrderTrackingInfo) SetCourierLocation(v GeoLocation)`

SetCourierLocation sets CourierLocation field to given value.

### HasCourierLocation

`func (o *OrderTrackingInfo) HasCourierLocation() bool`

HasCourierLocation returns a boolean if a field has been set.

### GetDropoffEta

`func (o *OrderTrackingInfo) GetDropoffEta() time.Time`

GetDropoffEta returns the DropoffEta field if non-nil, zero value otherwise.

### GetDropoffEtaOk

`func (o *OrderTrackingInfo) GetDropoffEtaOk() (*time.Time, bool)`

GetDropoffEtaOk returns a tuple with the DropoffEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropoffEta

`func (o *OrderTrackingInfo) SetDropoffEta(v time.Time)`

SetDropoffEta sets DropoffEta field to given value.

### HasDropoffEta

`func (o *OrderTrackingInfo) HasDropoffEta() bool`

HasDropoffEta returns a boolean if a field has been set.

### GetPickupEta

`func (o *OrderTrackingInfo) GetPickupEta() time.Time`

GetPickupEta returns the PickupEta field if non-nil, zero value otherwise.

### GetPickupEtaOk

`func (o *OrderTrackingInfo) GetPickupEtaOk() (*time.Time, bool)`

GetPickupEtaOk returns a tuple with the PickupEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupEta

`func (o *OrderTrackingInfo) SetPickupEta(v time.Time)`

SetPickupEta sets PickupEta field to given value.

### HasPickupEta

`func (o *OrderTrackingInfo) HasPickupEta() bool`

HasPickupEta returns a boolean if a field has been set.

### GetCourierName

`func (o *OrderTrackingInfo) GetCourierName() string`

GetCourierName returns the CourierName field if non-nil, zero value otherwise.

### GetCourierNameOk

`func (o *OrderTrackingInfo) GetCourierNameOk() (*string, bool)`

GetCourierNameOk returns a tuple with the CourierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourierName

`func (o *OrderTrackingInfo) SetCourierName(v string)`

SetCourierName sets CourierName field to given value.

### HasCourierName

`func (o *OrderTrackingInfo) HasCourierName() bool`

HasCourierName returns a boolean if a field has been set.

### GetCourierPhone

`func (o *OrderTrackingInfo) GetCourierPhone() string`

GetCourierPhone returns the CourierPhone field if non-nil, zero value otherwise.

### GetCourierPhoneOk

`func (o *OrderTrackingInfo) GetCourierPhoneOk() (*string, bool)`

GetCourierPhoneOk returns a tuple with the CourierPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourierPhone

`func (o *OrderTrackingInfo) SetCourierPhone(v string)`

SetCourierPhone sets CourierPhone field to given value.

### HasCourierPhone

`func (o *OrderTrackingInfo) HasCourierPhone() bool`

HasCourierPhone returns a boolean if a field has been set.

### GetDropoffWindowStart

`func (o *OrderTrackingInfo) GetDropoffWindowStart() time.Time`

GetDropoffWindowStart returns the DropoffWindowStart field if non-nil, zero value otherwise.

### GetDropoffWindowStartOk

`func (o *OrderTrackingInfo) GetDropoffWindowStartOk() (*time.Time, bool)`

GetDropoffWindowStartOk returns a tuple with the DropoffWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropoffWindowStart

`func (o *OrderTrackingInfo) SetDropoffWindowStart(v time.Time)`

SetDropoffWindowStart sets DropoffWindowStart field to given value.

### HasDropoffWindowStart

`func (o *OrderTrackingInfo) HasDropoffWindowStart() bool`

HasDropoffWindowStart returns a boolean if a field has been set.

### GetDropoffWindowEnd

`func (o *OrderTrackingInfo) GetDropoffWindowEnd() time.Time`

GetDropoffWindowEnd returns the DropoffWindowEnd field if non-nil, zero value otherwise.

### GetDropoffWindowEndOk

`func (o *OrderTrackingInfo) GetDropoffWindowEndOk() (*time.Time, bool)`

GetDropoffWindowEndOk returns a tuple with the DropoffWindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropoffWindowEnd

`func (o *OrderTrackingInfo) SetDropoffWindowEnd(v time.Time)`

SetDropoffWindowEnd sets DropoffWindowEnd field to given value.

### HasDropoffWindowEnd

`func (o *OrderTrackingInfo) HasDropoffWindowEnd() bool`

HasDropoffWindowEnd returns a boolean if a field has been set.

### GetPickupWindowStart

`func (o *OrderTrackingInfo) GetPickupWindowStart() time.Time`

GetPickupWindowStart returns the PickupWindowStart field if non-nil, zero value otherwise.

### GetPickupWindowStartOk

`func (o *OrderTrackingInfo) GetPickupWindowStartOk() (*time.Time, bool)`

GetPickupWindowStartOk returns a tuple with the PickupWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupWindowStart

`func (o *OrderTrackingInfo) SetPickupWindowStart(v time.Time)`

SetPickupWindowStart sets PickupWindowStart field to given value.

### HasPickupWindowStart

`func (o *OrderTrackingInfo) HasPickupWindowStart() bool`

HasPickupWindowStart returns a boolean if a field has been set.

### GetPickupWindowEnd

`func (o *OrderTrackingInfo) GetPickupWindowEnd() time.Time`

GetPickupWindowEnd returns the PickupWindowEnd field if non-nil, zero value otherwise.

### GetPickupWindowEndOk

`func (o *OrderTrackingInfo) GetPickupWindowEndOk() (*time.Time, bool)`

GetPickupWindowEndOk returns a tuple with the PickupWindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupWindowEnd

`func (o *OrderTrackingInfo) SetPickupWindowEnd(v time.Time)`

SetPickupWindowEnd sets PickupWindowEnd field to given value.

### HasPickupWindowEnd

`func (o *OrderTrackingInfo) HasPickupWindowEnd() bool`

HasPickupWindowEnd returns a boolean if a field has been set.

### GetIsCourierAssigned

`func (o *OrderTrackingInfo) GetIsCourierAssigned() bool`

GetIsCourierAssigned returns the IsCourierAssigned field if non-nil, zero value otherwise.

### GetIsCourierAssignedOk

`func (o *OrderTrackingInfo) GetIsCourierAssignedOk() (*bool, bool)`

GetIsCourierAssignedOk returns a tuple with the IsCourierAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCourierAssigned

`func (o *OrderTrackingInfo) SetIsCourierAssigned(v bool)`

SetIsCourierAssigned sets IsCourierAssigned field to given value.


### GetVehicleType

`func (o *OrderTrackingInfo) GetVehicleType() VehicleType`

GetVehicleType returns the VehicleType field if non-nil, zero value otherwise.

### GetVehicleTypeOk

`func (o *OrderTrackingInfo) GetVehicleTypeOk() (*VehicleType, bool)`

GetVehicleTypeOk returns a tuple with the VehicleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleType

`func (o *OrderTrackingInfo) SetVehicleType(v VehicleType)`

SetVehicleType sets VehicleType field to given value.

### HasVehicleType

`func (o *OrderTrackingInfo) HasVehicleType() bool`

HasVehicleType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


