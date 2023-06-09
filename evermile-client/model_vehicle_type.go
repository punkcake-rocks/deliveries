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
	"fmt"
)

// VehicleType the model 'VehicleType'
type VehicleType string

// List of vehicleType
const (
	VEHICLETYPE_PUSHBIKE VehicleType = "pushbike"
	VEHICLETYPE_MOTORBIKE VehicleType = "motorbike"
	VEHICLETYPE_CARGO_BIKE VehicleType = "cargo_bike"
	VEHICLETYPE_CAR VehicleType = "car"
	VEHICLETYPE_SMALL_VAN VehicleType = "small_van"
	VEHICLETYPE_LARGE_VAN VehicleType = "large_van"
)

// All allowed values of VehicleType enum
var AllowedVehicleTypeEnumValues = []VehicleType{
	"pushbike",
	"motorbike",
	"cargo_bike",
	"car",
	"small_van",
	"large_van",
}

func (v *VehicleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VehicleType(value)
	for _, existing := range AllowedVehicleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VehicleType", value)
}

// NewVehicleTypeFromValue returns a pointer to a valid VehicleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVehicleTypeFromValue(v string) (*VehicleType, error) {
	ev := VehicleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VehicleType: valid values are %v", v, AllowedVehicleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VehicleType) IsValid() bool {
	for _, existing := range AllowedVehicleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vehicleType value
func (v VehicleType) Ptr() *VehicleType {
	return &v
}

type NullableVehicleType struct {
	value *VehicleType
	isSet bool
}

func (v NullableVehicleType) Get() *VehicleType {
	return v.value
}

func (v *NullableVehicleType) Set(val *VehicleType) {
	v.value = val
	v.isSet = true
}

func (v NullableVehicleType) IsSet() bool {
	return v.isSet
}

func (v *NullableVehicleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVehicleType(val *VehicleType) *NullableVehicleType {
	return &NullableVehicleType{value: val, isSet: true}
}

func (v NullableVehicleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVehicleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

