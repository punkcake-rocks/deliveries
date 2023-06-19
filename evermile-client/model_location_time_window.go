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

// checks if the LocationTimeWindow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationTimeWindow{}

// LocationTimeWindow A time window during which a location is open for pickups
type LocationTimeWindow struct {
	// Whether the location is open or not
	IsOpen bool `json:"isOpen"`
	// The time when this location opens in HH:MM format
	OpenWindowStart string `json:"openWindowStart"`
	// The time when this location closes in HH:MM format
	OpenWindowEnd string `json:"openWindowEnd"`
}

// NewLocationTimeWindow instantiates a new LocationTimeWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationTimeWindow(isOpen bool, openWindowStart string, openWindowEnd string) *LocationTimeWindow {
	this := LocationTimeWindow{}
	this.IsOpen = isOpen
	this.OpenWindowStart = openWindowStart
	this.OpenWindowEnd = openWindowEnd
	return &this
}

// NewLocationTimeWindowWithDefaults instantiates a new LocationTimeWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationTimeWindowWithDefaults() *LocationTimeWindow {
	this := LocationTimeWindow{}
	var isOpen bool = true
	this.IsOpen = isOpen
	return &this
}

// GetIsOpen returns the IsOpen field value
func (o *LocationTimeWindow) GetIsOpen() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsOpen
}

// GetIsOpenOk returns a tuple with the IsOpen field value
// and a boolean to check if the value has been set.
func (o *LocationTimeWindow) GetIsOpenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsOpen, true
}

// SetIsOpen sets field value
func (o *LocationTimeWindow) SetIsOpen(v bool) {
	o.IsOpen = v
}

// GetOpenWindowStart returns the OpenWindowStart field value
func (o *LocationTimeWindow) GetOpenWindowStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OpenWindowStart
}

// GetOpenWindowStartOk returns a tuple with the OpenWindowStart field value
// and a boolean to check if the value has been set.
func (o *LocationTimeWindow) GetOpenWindowStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenWindowStart, true
}

// SetOpenWindowStart sets field value
func (o *LocationTimeWindow) SetOpenWindowStart(v string) {
	o.OpenWindowStart = v
}

// GetOpenWindowEnd returns the OpenWindowEnd field value
func (o *LocationTimeWindow) GetOpenWindowEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OpenWindowEnd
}

// GetOpenWindowEndOk returns a tuple with the OpenWindowEnd field value
// and a boolean to check if the value has been set.
func (o *LocationTimeWindow) GetOpenWindowEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenWindowEnd, true
}

// SetOpenWindowEnd sets field value
func (o *LocationTimeWindow) SetOpenWindowEnd(v string) {
	o.OpenWindowEnd = v
}

func (o LocationTimeWindow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationTimeWindow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isOpen"] = o.IsOpen
	toSerialize["openWindowStart"] = o.OpenWindowStart
	toSerialize["openWindowEnd"] = o.OpenWindowEnd
	return toSerialize, nil
}

type NullableLocationTimeWindow struct {
	value *LocationTimeWindow
	isSet bool
}

func (v NullableLocationTimeWindow) Get() *LocationTimeWindow {
	return v.value
}

func (v *NullableLocationTimeWindow) Set(val *LocationTimeWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationTimeWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationTimeWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationTimeWindow(val *LocationTimeWindow) *NullableLocationTimeWindow {
	return &NullableLocationTimeWindow{value: val, isSet: true}
}

func (v NullableLocationTimeWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationTimeWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


