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
)

// checks if the HandoffContactInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandoffContactInfo{}

// HandoffContactInfo Information about the handoff contact info
type HandoffContactInfo struct {
	Name string `json:"name"`
	// A phone number of the handoff receiver (leader)
	Phone string `json:"phone"`
	Address AddressComponents `json:"address"`
}

// NewHandoffContactInfo instantiates a new HandoffContactInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandoffContactInfo(name string, phone string, address AddressComponents) *HandoffContactInfo {
	this := HandoffContactInfo{}
	this.Name = name
	this.Phone = phone
	this.Address = address
	return &this
}

// NewHandoffContactInfoWithDefaults instantiates a new HandoffContactInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandoffContactInfoWithDefaults() *HandoffContactInfo {
	this := HandoffContactInfo{}
	return &this
}

// GetName returns the Name field value
func (o *HandoffContactInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HandoffContactInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *HandoffContactInfo) SetName(v string) {
	o.Name = v
}

// GetPhone returns the Phone field value
func (o *HandoffContactInfo) GetPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value
// and a boolean to check if the value has been set.
func (o *HandoffContactInfo) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phone, true
}

// SetPhone sets field value
func (o *HandoffContactInfo) SetPhone(v string) {
	o.Phone = v
}

// GetAddress returns the Address field value
func (o *HandoffContactInfo) GetAddress() AddressComponents {
	if o == nil {
		var ret AddressComponents
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *HandoffContactInfo) GetAddressOk() (*AddressComponents, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *HandoffContactInfo) SetAddress(v AddressComponents) {
	o.Address = v
}

func (o HandoffContactInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandoffContactInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["phone"] = o.Phone
	toSerialize["address"] = o.Address
	return toSerialize, nil
}

type NullableHandoffContactInfo struct {
	value *HandoffContactInfo
	isSet bool
}

func (v NullableHandoffContactInfo) Get() *HandoffContactInfo {
	return v.value
}

func (v *NullableHandoffContactInfo) Set(val *HandoffContactInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHandoffContactInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHandoffContactInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandoffContactInfo(val *HandoffContactInfo) *NullableHandoffContactInfo {
	return &NullableHandoffContactInfo{value: val, isSet: true}
}

func (v NullableHandoffContactInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandoffContactInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

