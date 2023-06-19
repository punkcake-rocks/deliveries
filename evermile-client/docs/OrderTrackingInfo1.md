# OrderTrackingInfo1

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

### NewOrderTrackingInfo1

`func NewOrderTrackingInfo1(isCourierAssigned bool, ) *OrderTrackingInfo1`

NewOrderTrackingInfo1 instantiates a new OrderTrackingInfo1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderTrackingInfo1WithDefaults

`func NewOrderTrackingInfo1WithDefaults() *OrderTrackingInfo1`

NewOrderTrackingInfo1WithDefaults instantiates a new OrderTrackingInfo1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCourierLocation

`func (o *OrderTrackingInfo1) GetCourierLocation() GeoLocation`

GetCourierLocation returns the CourierLocation field if non-nil, zero value otherwise.

### GetCourierLocationOk

`func (o *OrderTrackingInfo1) GetCourierLocationOk() (*GeoLocation, bool)`

GetCourierLocationOk returns a tuple with the CourierLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourierLocation

`func (o *OrderTrackingInfo1) SetCourierLocation(v GeoLocation)`

SetCourierLocation sets CourierLocation field to given value.

### HasCourierLocation

`func (o *OrderTrackingInfo1) HasCourierLocation() bool`

HasCourierLocation returns a boolean if a field has been set.

### GetDropoffEta

`func (o *OrderTrackingInfo1) GetDropoffEta() time.Time`

GetDropoffEta returns the DropoffEta field if non-nil, zero value otherwise.

### GetDropoffEtaOk

`func (o *OrderTrackingInfo1) GetDropoffEtaOk() (*time.Time, bool)`

GetDropoffEtaOk returns a tuple with the DropoffEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropoffEta

`func (o *OrderTrackingInfo1) SetDropoffEta(v time.Time)`

SetDropoffEta sets DropoffEta field to given value.

### HasDropoffEta

`func (o *OrderTrackingInfo1) HasDropoffEta() bool`

HasDropoffEta returns a boolean if a field has been set.

### GetPickupEta

`func (o *OrderTrackingInfo1) GetPickupEta() time.Time`

GetPickupEta returns the PickupEta field if non-nil, zero value otherwise.

### GetPickupEtaOk

`func (o *OrderTrackingInfo1) GetPickupEtaOk() (*time.Time, bool)`

GetPickupEtaOk returns a tuple with the PickupEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupEta

`func (o *OrderTrackingInfo1) SetPickupEta(v time.Time)`

SetPickupEta sets PickupEta field to given value.

### HasPickupEta

`func (o *OrderTrackingInfo1) HasPickupEta() bool`

HasPickupEta returns a boolean if a field has been set.

### GetCourierName

`func (o *OrderTrackingInfo1) GetCourierName() string`

GetCourierName returns the CourierName field if non-nil, zero value otherwise.

### GetCourierNameOk

`func (o *OrderTrackingInfo1) GetCourierNameOk() (*string, bool)`

GetCourierNameOk returns a tuple with the CourierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourierName

`func (o *OrderTrackingInfo1) SetCourierName(v string)`

SetCourierName sets CourierName field to given value.

### HasCourierName

`func (o *OrderTrackingInfo1) HasCourierName() bool`

HasCourierName returns a boolean if a field has been set.

### GetCourierPhone

`func (o *OrderTrackingInfo1) GetCourierPhone() string`

GetCourierPhone returns the CourierPhone field if non-nil, zero value otherwise.

### GetCourierPhoneOk

`func (o *OrderTrackingInfo1) GetCourierPhoneOk() (*string, bool)`

GetCourierPhoneOk returns a tuple with the CourierPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCourierPhone

`func (o *OrderTrackingInfo1) SetCourierPhone(v string)`

SetCourierPhone sets CourierPhone field to given value.

### HasCourierPhone

`func (o *OrderTrackingInfo1) HasCourierPhone() bool`

HasCourierPhone returns a boolean if a field has been set.

### GetDropoffWindowStart

`func (o *OrderTrackingInfo1) GetDropoffWindowStart() time.Time`

GetDropoffWindowStart returns the DropoffWindowStart field if non-nil, zero value otherwise.

### GetDropoffWindowStartOk

`func (o *OrderTrackingInfo1) GetDropoffWindowStartOk() (*time.Time, bool)`

GetDropoffWindowStartOk returns a tuple with the DropoffWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropoffWindowStart

`func (o *OrderTrackingInfo1) SetDropoffWindowStart(v time.Time)`

SetDropoffWindowStart sets DropoffWindowStart field to given value.

### HasDropoffWindowStart

`func (o *OrderTrackingInfo1) HasDropoffWindowStart() bool`

HasDropoffWindowStart returns a boolean if a field has been set.

### GetDropoffWindowEnd

`func (o *OrderTrackingInfo1) GetDropoffWindowEnd() time.Time`

GetDropoffWindowEnd returns the DropoffWindowEnd field if non-nil, zero value otherwise.

### GetDropoffWindowEndOk

`func (o *OrderTrackingInfo1) GetDropoffWindowEndOk() (*time.Time, bool)`

GetDropoffWindowEndOk returns a tuple with the DropoffWindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropoffWindowEnd

`func (o *OrderTrackingInfo1) SetDropoffWindowEnd(v time.Time)`

SetDropoffWindowEnd sets DropoffWindowEnd field to given value.

### HasDropoffWindowEnd

`func (o *OrderTrackingInfo1) HasDropoffWindowEnd() bool`

HasDropoffWindowEnd returns a boolean if a field has been set.

### GetPickupWindowStart

`func (o *OrderTrackingInfo1) GetPickupWindowStart() time.Time`

GetPickupWindowStart returns the PickupWindowStart field if non-nil, zero value otherwise.

### GetPickupWindowStartOk

`func (o *OrderTrackingInfo1) GetPickupWindowStartOk() (*time.Time, bool)`

GetPickupWindowStartOk returns a tuple with the PickupWindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupWindowStart

`func (o *OrderTrackingInfo1) SetPickupWindowStart(v time.Time)`

SetPickupWindowStart sets PickupWindowStart field to given value.

### HasPickupWindowStart

`func (o *OrderTrackingInfo1) HasPickupWindowStart() bool`

HasPickupWindowStart returns a boolean if a field has been set.

### GetPickupWindowEnd

`func (o *OrderTrackingInfo1) GetPickupWindowEnd() time.Time`

GetPickupWindowEnd returns the PickupWindowEnd field if non-nil, zero value otherwise.

### GetPickupWindowEndOk

`func (o *OrderTrackingInfo1) GetPickupWindowEndOk() (*time.Time, bool)`

GetPickupWindowEndOk returns a tuple with the PickupWindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupWindowEnd

`func (o *OrderTrackingInfo1) SetPickupWindowEnd(v time.Time)`

SetPickupWindowEnd sets PickupWindowEnd field to given value.

### HasPickupWindowEnd

`func (o *OrderTrackingInfo1) HasPickupWindowEnd() bool`

HasPickupWindowEnd returns a boolean if a field has been set.

### GetIsCourierAssigned

`func (o *OrderTrackingInfo1) GetIsCourierAssigned() bool`

GetIsCourierAssigned returns the IsCourierAssigned field if non-nil, zero value otherwise.

### GetIsCourierAssignedOk

`func (o *OrderTrackingInfo1) GetIsCourierAssignedOk() (*bool, bool)`

GetIsCourierAssignedOk returns a tuple with the IsCourierAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCourierAssigned

`func (o *OrderTrackingInfo1) SetIsCourierAssigned(v bool)`

SetIsCourierAssigned sets IsCourierAssigned field to given value.


### GetVehicleType

`func (o *OrderTrackingInfo1) GetVehicleType() VehicleType`

GetVehicleType returns the VehicleType field if non-nil, zero value otherwise.

### GetVehicleTypeOk

`func (o *OrderTrackingInfo1) GetVehicleTypeOk() (*VehicleType, bool)`

GetVehicleTypeOk returns a tuple with the VehicleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVehicleType

`func (o *OrderTrackingInfo1) SetVehicleType(v VehicleType)`

SetVehicleType sets VehicleType field to given value.

### HasVehicleType

`func (o *OrderTrackingInfo1) HasVehicleType() bool`

HasVehicleType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


