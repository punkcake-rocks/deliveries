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

// checks if the HandoffInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HandoffInfo{}

// HandoffInfo Information about the handoff
type HandoffInfo struct {
	ContactInfo HandoffContactInfo `json:"contactInfo"`
	HandoffTime OptionalTimeWindow `json:"handoffTime"`
	// walking distance in meters
	WalkingDistanceMeters int64 `json:"walkingDistanceMeters"`
}

// NewHandoffInfo instantiates a new HandoffInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHandoffInfo(contactInfo HandoffContactInfo, handoffTime OptionalTimeWindow, walkingDistanceMeters int64) *HandoffInfo {
	this := HandoffInfo{}
	this.ContactInfo = contactInfo
	this.HandoffTime = handoffTime
	this.WalkingDistanceMeters = walkingDistanceMeters
	return &this
}

// NewHandoffInfoWithDefaults instantiates a new HandoffInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHandoffInfoWithDefaults() *HandoffInfo {
	this := HandoffInfo{}
	return &this
}

// GetContactInfo returns the ContactInfo field value
func (o *HandoffInfo) GetContactInfo() HandoffContactInfo {
	if o == nil {
		var ret HandoffContactInfo
		return ret
	}

	return o.ContactInfo
}

// GetContactInfoOk returns a tuple with the ContactInfo field value
// and a boolean to check if the value has been set.
func (o *HandoffInfo) GetContactInfoOk() (*HandoffContactInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactInfo, true
}

// SetContactInfo sets field value
func (o *HandoffInfo) SetContactInfo(v HandoffContactInfo) {
	o.ContactInfo = v
}

// GetHandoffTime returns the HandoffTime field value
func (o *HandoffInfo) GetHandoffTime() OptionalTimeWindow {
	if o == nil {
		var ret OptionalTimeWindow
		return ret
	}

	return o.HandoffTime
}

// GetHandoffTimeOk returns a tuple with the HandoffTime field value
// and a boolean to check if the value has been set.
func (o *HandoffInfo) GetHandoffTimeOk() (*OptionalTimeWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandoffTime, true
}

// SetHandoffTime sets field value
func (o *HandoffInfo) SetHandoffTime(v OptionalTimeWindow) {
	o.HandoffTime = v
}

// GetWalkingDistanceMeters returns the WalkingDistanceMeters field value
func (o *HandoffInfo) GetWalkingDistanceMeters() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.WalkingDistanceMeters
}

// GetWalkingDistanceMetersOk returns a tuple with the WalkingDistanceMeters field value
// and a boolean to check if the value has been set.
func (o *HandoffInfo) GetWalkingDistanceMetersOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WalkingDistanceMeters, true
}

// SetWalkingDistanceMeters sets field value
func (o *HandoffInfo) SetWalkingDistanceMeters(v int64) {
	o.WalkingDistanceMeters = v
}

func (o HandoffInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HandoffInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contactInfo"] = o.ContactInfo
	toSerialize["handoffTime"] = o.HandoffTime
	toSerialize["walkingDistanceMeters"] = o.WalkingDistanceMeters
	return toSerialize, nil
}

type NullableHandoffInfo struct {
	value *HandoffInfo
	isSet bool
}

func (v NullableHandoffInfo) Get() *HandoffInfo {
	return v.value
}

func (v *NullableHandoffInfo) Set(val *HandoffInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableHandoffInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableHandoffInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandoffInfo(val *HandoffInfo) *NullableHandoffInfo {
	return &NullableHandoffInfo{value: val, isSet: true}
}

func (v NullableHandoffInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandoffInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


