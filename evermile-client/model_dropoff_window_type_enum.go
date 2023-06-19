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

// DropoffWindowTypeEnum <u>The reported issue</u>: <br>             <table>   <tr><td>scheduledWindow</td>     <td>A scheduled delivery window</td></tr>   <tr><td>allDay</td>              <td>An all day window</td></tr>   <tr><td>express</td>             <td>An express window</td></tr> </table> 
type DropoffWindowTypeEnum string

// List of dropoffWindowTypeEnum
const (
	DROPOFFWINDOWTYPEENUM_SCHEDULED_WINDOW DropoffWindowTypeEnum = "scheduledWindow"
	DROPOFFWINDOWTYPEENUM_ALL_DAY DropoffWindowTypeEnum = "allDay"
	DROPOFFWINDOWTYPEENUM_EXPRESS DropoffWindowTypeEnum = "express"
)

// All allowed values of DropoffWindowTypeEnum enum
var AllowedDropoffWindowTypeEnumEnumValues = []DropoffWindowTypeEnum{
	"scheduledWindow",
	"allDay",
	"express",
}

func (v *DropoffWindowTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DropoffWindowTypeEnum(value)
	for _, existing := range AllowedDropoffWindowTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DropoffWindowTypeEnum", value)
}

// NewDropoffWindowTypeEnumFromValue returns a pointer to a valid DropoffWindowTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDropoffWindowTypeEnumFromValue(v string) (*DropoffWindowTypeEnum, error) {
	ev := DropoffWindowTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DropoffWindowTypeEnum: valid values are %v", v, AllowedDropoffWindowTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DropoffWindowTypeEnum) IsValid() bool {
	for _, existing := range AllowedDropoffWindowTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to dropoffWindowTypeEnum value
func (v DropoffWindowTypeEnum) Ptr() *DropoffWindowTypeEnum {
	return &v
}

type NullableDropoffWindowTypeEnum struct {
	value *DropoffWindowTypeEnum
	isSet bool
}

func (v NullableDropoffWindowTypeEnum) Get() *DropoffWindowTypeEnum {
	return v.value
}

func (v *NullableDropoffWindowTypeEnum) Set(val *DropoffWindowTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableDropoffWindowTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableDropoffWindowTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDropoffWindowTypeEnum(val *DropoffWindowTypeEnum) *NullableDropoffWindowTypeEnum {
	return &NullableDropoffWindowTypeEnum{value: val, isSet: true}
}

func (v NullableDropoffWindowTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDropoffWindowTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
