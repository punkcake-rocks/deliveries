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

// HandoffType the model 'HandoffType'
type HandoffType string

// List of handoffType
const (
	HANDOFFTYPE_NONE HandoffType = "none"
	HANDOFFTYPE_RECEIVE HandoffType = "receive"
	HANDOFFTYPE_HANDOFF HandoffType = "handoff"
)

// All allowed values of HandoffType enum
var AllowedHandoffTypeEnumValues = []HandoffType{
	"none",
	"receive",
	"handoff",
}

func (v *HandoffType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HandoffType(value)
	for _, existing := range AllowedHandoffTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HandoffType", value)
}

// NewHandoffTypeFromValue returns a pointer to a valid HandoffType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHandoffTypeFromValue(v string) (*HandoffType, error) {
	ev := HandoffType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HandoffType: valid values are %v", v, AllowedHandoffTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HandoffType) IsValid() bool {
	for _, existing := range AllowedHandoffTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to handoffType value
func (v HandoffType) Ptr() *HandoffType {
	return &v
}

type NullableHandoffType struct {
	value *HandoffType
	isSet bool
}

func (v NullableHandoffType) Get() *HandoffType {
	return v.value
}

func (v *NullableHandoffType) Set(val *HandoffType) {
	v.value = val
	v.isSet = true
}

func (v NullableHandoffType) IsSet() bool {
	return v.isSet
}

func (v *NullableHandoffType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHandoffType(val *HandoffType) *NullableHandoffType {
	return &NullableHandoffType{value: val, isSet: true}
}

func (v NullableHandoffType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHandoffType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

