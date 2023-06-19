/*
Evermile Commercial Quotes and Booking API

This is Evermile's commercial API for handling delivery quotes and orders

API version: 1.0
Contact: support@evermile.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package EvermileClient

import (
	"encoding/json"
	"time"
)

// checks if the OrderTrackingInfo1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderTrackingInfo1{}

// OrderTrackingInfo1 Live tracking info for an order
type OrderTrackingInfo1 struct {
	CourierLocation *GeoLocation `json:"courierLocation,omitempty"`
	// The dropoff ETA in ISO8601 format
	DropoffEta *time.Time `json:"dropoffEta,omitempty"`
	// The pickup ETA in ISO8601 format
	PickupEta *time.Time `json:"pickupEta,omitempty"`
	// The courier full name
	CourierName *string `json:"courierName,omitempty"`
	// The courier phone number
	CourierPhone *string `json:"courierPhone,omitempty"`
	// The updated dropoff widnow start in ISO8601 format
	DropoffWindowStart *time.Time `json:"dropoffWindowStart,omitempty"`
	// The updated dropoff window end in ISO8601 format
	DropoffWindowEnd *time.Time `json:"dropoffWindowEnd,omitempty"`
	// The updated pickup widnow start in ISO8601 format
	PickupWindowStart *time.Time `json:"pickupWindowStart,omitempty"`
	// The updated pickup window end in ISO8601 format
	PickupWindowEnd *time.Time `json:"pickupWindowEnd,omitempty"`
	// Indicate whether a courier is assigned for this order
	IsCourierAssigned bool `json:"isCourierAssigned"`
	VehicleType *VehicleType `json:"vehicleType,omitempty"`
}

// NewOrderTrackingInfo1 instantiates a new OrderTrackingInfo1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderTrackingInfo1(isCourierAssigned bool) *OrderTrackingInfo1 {
	this := OrderTrackingInfo1{}
	this.IsCourierAssigned = isCourierAssigned
	return &this
}

// NewOrderTrackingInfo1WithDefaults instantiates a new OrderTrackingInfo1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderTrackingInfo1WithDefaults() *OrderTrackingInfo1 {
	this := OrderTrackingInfo1{}
	return &this
}

// GetCourierLocation returns the CourierLocation field value if set, zero value otherwise.
func (o *OrderTrackingInfo1) GetCourierLocation() GeoLocation {
	if o == nil || IsNil(o.CourierLocation) {
		var ret GeoLocation
		return ret
	}
	return *o.CourierLocation
}

// GetCourierLocationOk returns a tuple with the CourierLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTrackingInfo1) GetCourierLocationOk() (*GeoLocation, bool) {
	if o == nil || IsNil(o.CourierLocation) {
		return nil, false
	}
	return o.CourierLocation, true
}

// HasCourierLocation returns a boolean if a field has been set.
func (o *OrderTrackingInfo1) HasCourierLocation() bool {
	if o != nil && !IsNil(o.CourierLocation) {
		return true
	}

	return false
}

// SetCourierLocation gets a reference to the given GeoLocation and assigns it to the CourierLocation field.
func (o *OrderTrackingInfo1) SetCourierLocation(v GeoLocation) {
	o.CourierLocation = &v
}

// GetDropoffEta returns the DropoffEta field value if set, zero value otherwise.
func (o *OrderTrackingInfo1) GetDropoffEta() time.Time {
	if o == nil || IsNil(o.DropoffEta) {
		var ret time.Time
		return ret
	}
	return *o.DropoffEta
}

// GetDropoffEtaOk returns a tuple with the DropoffEta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTrackingInfo1) GetDropoffEtaOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DropoffEta) {
		return nil, false
	}
	return o.DropoffEta, true
}

// HasDropoffEta returns a boolean if a field has been set.
func (o *OrderTrackingInfo1) HasDropoffEta() bool {
	if o != nil && !IsNil(o.DropoffEta) {
		return true
	}

	return false
}

// SetDropoffEta gets a reference to the given time.Time and assigns it to the DropoffEta field.
func (o *OrderTrackingInfo1) SetDropoffEta(v time.Time) {
	o.DropoffEta = &v
}

// GetPickupEta returns the PickupEta field value if set, zero value otherwise.
func (o *OrderTrackingInfo1) GetPickupEta() time.Time {
	if o == nil || IsNil(o.PickupEta) {
		var ret time.Time
		return ret
	}
	return *o.PickupEta
}

// GetPickupEtaOk returns a tuple with the PickupEta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTrackingInfo1) GetPickupEtaOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PickupEta) {
		return nil, false
	}
	return o.PickupEta, true
}

// HasPickupEta returns a boolean if a field has been set.
func (o *OrderTrackingInfo1) HasPickupEta() bool {
	if o != nil && !IsNil(o.PickupEta) {
		return true
	}

	return false
}

// SetPickupEta gets a reference to the given time.Time and assigns it to the PickupEta field.
func (o *OrderTrackingInfo1) SetPickupEta(v time.Time) {
	o.PickupEta = &v
}

// GetCourierName returns the CourierName field value if set, zero value otherwise.
func (o *OrderTrackingInfo1) GetCourierName() string {
	if o == nil || IsNil(o.CourierName) {
		var ret string
		return ret
	}
	return *o.CourierName
}

// GetCourierNameOk returns a tuple with the CourierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTrackingInfo1) GetCourierNameOk() (*string, bool) {
	if o == nil || IsNil(o.CourierName) {
		return nil, false
	}
	return o.CourierName, true
}

// HasCourierName returns a boolean if a field has been set.
func (o *OrderTrackingInfo1) HasCourierName() bool {
	if o != nil && !IsNil(o.CourierName) {
		return true
	}

	return false
}

// SetCourierName gets a reference to the given string and assigns it to the CourierName field.
func (o *OrderTrackingInfo1) SetCourierName(v string) {
	o.CourierName = &v
}

// GetCourierPhone returns the CourierPhone field value if set, zero value otherwise.
func (o *OrderTrackingInfo1) GetCourierPhone() string {
	if o == nil || IsNil(o.CourierPhone) {
		var ret string
		return ret
	}
	return *o.CourierPhone
}

// GetCourierPhoneOk returns a tuple with the CourierPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTrackingInfo1) GetCourierPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.CourierPhone) {
		return nil, false
	}
	return o.CourierPhone, true
}

// HasCourierPhone returns a boolean if a field has been set.
func (o *OrderTrackingInfo1) HasCourierPhone() bool {
	if o != nil && !IsNil(o.CourierPhone) {
		return true
	}

	return false
}

// SetCourierPhone gets a reference to the given string and assigns it to the CourierPhone field.
func (o *OrderTrackingInfo1) SetCourierPhone(v string) {
	o.CourierPhone = &v
}

// GetDropoffWindowStart returns the DropoffWindowStart field value if set, zero value otherwise.
func (o *OrderTrackingInfo1) GetDropoffWindowStart() time.Time {
	if o == nil || IsNil(o.DropoffWindowStart) {
		var ret time.Time
		return ret
	}
	return *o.DropoffWindowStart
}

// GetDropoffWindowStartOk returns a tuple with the DropoffWindowStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTrackingInfo1) GetDropoffWindowStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DropoffWindowStart) {
		return nil, false
	}
	return o.DropoffWindowStart, true
}

// HasDropoffWindowStart returns a boolean if a field has been set.
func (o *OrderTrackingInfo1) HasDropoffWindowStart() bool {
	if o != nil && !IsNil(o.DropoffWindowStart) {
		return true
	}

	return false
}

// SetDropoffWindowStart gets a reference to the given time.Time and assigns it to the DropoffWindowStart field.
func (o *OrderTrackingInfo1) SetDropoffWindowStart(v time.Time) {
	o.DropoffWindowStart = &v
}

// GetDropoffWindowEnd returns the DropoffWindowEnd field value if set, zero value otherwise.
func (o *OrderTrackingInfo1) GetDropoffWindowEnd() time.Time {
	if o == nil || IsNil(o.DropoffWindowEnd) {
		var ret time.Time
		return ret
	}
	return *o.DropoffWindowEnd
}

// GetDropoffWindowEndOk returns a tuple with the DropoffWindowEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTrackingInfo1) GetDropoffWindowEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DropoffWindowEnd) {
		return nil, false
	}
	return o.DropoffWindowEnd, true
}

// HasDropoffWindowEnd returns a boolean if a field has been set.
func (o *OrderTrackingInfo1) HasDropoffWindowEnd() bool {
	if o != nil && !IsNil(o.DropoffWindowEnd) {
		return true
	}

	return false
}

// SetDropoffWindowEnd gets a reference to the given time.Time and assigns it to the DropoffWindowEnd field.
func (o *OrderTrackingInfo1) SetDropoffWindowEnd(v time.Time) {
	o.DropoffWindowEnd = &v
}

// GetPickupWindowStart returns the PickupWindowStart field value if set, zero value otherwise.
func (o *OrderTrackingInfo1) GetPickupWindowStart() time.Time {
	if o == nil || IsNil(o.PickupWindowStart) {
		var ret time.Time
		return ret
	}
	return *o.PickupWindowStart
}

// GetPickupWindowStartOk returns a tuple with the PickupWindowStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTrackingInfo1) GetPickupWindowStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PickupWindowStart) {
		return nil, false
	}
	return o.PickupWindowStart, true
}

// HasPickupWindowStart returns a boolean if a field has been set.
func (o *OrderTrackingInfo1) HasPickupWindowStart() bool {
	if o != nil && !IsNil(o.PickupWindowStart) {
		return true
	}

	return false
}

// SetPickupWindowStart gets a reference to the given time.Time and assigns it to the PickupWindowStart field.
func (o *OrderTrackingInfo1) SetPickupWindowStart(v time.Time) {
	o.PickupWindowStart = &v
}

// GetPickupWindowEnd returns the PickupWindowEnd field value if set, zero value otherwise.
func (o *OrderTrackingInfo1) GetPickupWindowEnd() time.Time {
	if o == nil || IsNil(o.PickupWindowEnd) {
		var ret time.Time
		return ret
	}
	return *o.PickupWindowEnd
}

// GetPickupWindowEndOk returns a tuple with the PickupWindowEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTrackingInfo1) GetPickupWindowEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.PickupWindowEnd) {
		return nil, false
	}
	return o.PickupWindowEnd, true
}

// HasPickupWindowEnd returns a boolean if a field has been set.
func (o *OrderTrackingInfo1) HasPickupWindowEnd() bool {
	if o != nil && !IsNil(o.PickupWindowEnd) {
		return true
	}

	return false
}

// SetPickupWindowEnd gets a reference to the given time.Time and assigns it to the PickupWindowEnd field.
func (o *OrderTrackingInfo1) SetPickupWindowEnd(v time.Time) {
	o.PickupWindowEnd = &v
}

// GetIsCourierAssigned returns the IsCourierAssigned field value
func (o *OrderTrackingInfo1) GetIsCourierAssigned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCourierAssigned
}

// GetIsCourierAssignedOk returns a tuple with the IsCourierAssigned field value
// and a boolean to check if the value has been set.
func (o *OrderTrackingInfo1) GetIsCourierAssignedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCourierAssigned, true
}

// SetIsCourierAssigned sets field value
func (o *OrderTrackingInfo1) SetIsCourierAssigned(v bool) {
	o.IsCourierAssigned = v
}

// GetVehicleType returns the VehicleType field value if set, zero value otherwise.
func (o *OrderTrackingInfo1) GetVehicleType() VehicleType {
	if o == nil || IsNil(o.VehicleType) {
		var ret VehicleType
		return ret
	}
	return *o.VehicleType
}

// GetVehicleTypeOk returns a tuple with the VehicleType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderTrackingInfo1) GetVehicleTypeOk() (*VehicleType, bool) {
	if o == nil || IsNil(o.VehicleType) {
		return nil, false
	}
	return o.VehicleType, true
}

// HasVehicleType returns a boolean if a field has been set.
func (o *OrderTrackingInfo1) HasVehicleType() bool {
	if o != nil && !IsNil(o.VehicleType) {
		return true
	}

	return false
}

// SetVehicleType gets a reference to the given VehicleType and assigns it to the VehicleType field.
func (o *OrderTrackingInfo1) SetVehicleType(v VehicleType) {
	o.VehicleType = &v
}

func (o OrderTrackingInfo1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderTrackingInfo1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CourierLocation) {
		toSerialize["courierLocation"] = o.CourierLocation
	}
	if !IsNil(o.DropoffEta) {
		toSerialize["dropoffEta"] = o.DropoffEta
	}
	if !IsNil(o.PickupEta) {
		toSerialize["pickupEta"] = o.PickupEta
	}
	if !IsNil(o.CourierName) {
		toSerialize["courierName"] = o.CourierName
	}
	if !IsNil(o.CourierPhone) {
		toSerialize["courierPhone"] = o.CourierPhone
	}
	if !IsNil(o.DropoffWindowStart) {
		toSerialize["dropoffWindowStart"] = o.DropoffWindowStart
	}
	if !IsNil(o.DropoffWindowEnd) {
		toSerialize["dropoffWindowEnd"] = o.DropoffWindowEnd
	}
	if !IsNil(o.PickupWindowStart) {
		toSerialize["pickupWindowStart"] = o.PickupWindowStart
	}
	if !IsNil(o.PickupWindowEnd) {
		toSerialize["pickupWindowEnd"] = o.PickupWindowEnd
	}
	toSerialize["isCourierAssigned"] = o.IsCourierAssigned
	if !IsNil(o.VehicleType) {
		toSerialize["vehicleType"] = o.VehicleType
	}
	return toSerialize, nil
}

type NullableOrderTrackingInfo1 struct {
	value *OrderTrackingInfo1
	isSet bool
}

func (v NullableOrderTrackingInfo1) Get() *OrderTrackingInfo1 {
	return v.value
}

func (v *NullableOrderTrackingInfo1) Set(val *OrderTrackingInfo1) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderTrackingInfo1) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderTrackingInfo1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderTrackingInfo1(val *OrderTrackingInfo1) *NullableOrderTrackingInfo1 {
	return &NullableOrderTrackingInfo1{value: val, isSet: true}
}

func (v NullableOrderTrackingInfo1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderTrackingInfo1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

