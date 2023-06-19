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

// ProposalLabelEnum <u>Indication of a pricing label</u>: <br>             <table>   <tr><td>best</td>      <td>The \"best\" option (balanced between quickest and cheapest)</td></tr>   <tr><td>cheapest</td>  <td>The cheapest option</td></tr>   <tr><td>quickest</td>  <td>The quickest option</td></tr> </table> 
type ProposalLabelEnum string

// List of proposalLabelEnum
const (
	PROPOSALLABELENUM_BEST ProposalLabelEnum = "best"
	PROPOSALLABELENUM_CHEAPEST ProposalLabelEnum = "cheapest"
	PROPOSALLABELENUM_QUICKEST ProposalLabelEnum = "quickest"
	PROPOSALLABELENUM_CURRENT ProposalLabelEnum = "current"
)

// All allowed values of ProposalLabelEnum enum
var AllowedProposalLabelEnumEnumValues = []ProposalLabelEnum{
	"best",
	"cheapest",
	"quickest",
	"current",
}

func (v *ProposalLabelEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProposalLabelEnum(value)
	for _, existing := range AllowedProposalLabelEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProposalLabelEnum", value)
}

// NewProposalLabelEnumFromValue returns a pointer to a valid ProposalLabelEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProposalLabelEnumFromValue(v string) (*ProposalLabelEnum, error) {
	ev := ProposalLabelEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProposalLabelEnum: valid values are %v", v, AllowedProposalLabelEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProposalLabelEnum) IsValid() bool {
	for _, existing := range AllowedProposalLabelEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to proposalLabelEnum value
func (v ProposalLabelEnum) Ptr() *ProposalLabelEnum {
	return &v
}

type NullableProposalLabelEnum struct {
	value *ProposalLabelEnum
	isSet bool
}

func (v NullableProposalLabelEnum) Get() *ProposalLabelEnum {
	return v.value
}

func (v *NullableProposalLabelEnum) Set(val *ProposalLabelEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableProposalLabelEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableProposalLabelEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProposalLabelEnum(val *ProposalLabelEnum) *NullableProposalLabelEnum {
	return &NullableProposalLabelEnum{value: val, isSet: true}
}

func (v NullableProposalLabelEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProposalLabelEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
