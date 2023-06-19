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

// checks if the DateProposals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DateProposals{}

// DateProposals struct for DateProposals
type DateProposals struct {
	// The date for which these proposals are proposed in format YYYY-MM-DD
	Date string `json:"date"`
	// A list of timeslot proposals in the given date
	Proposals []TimeslotProposal `json:"proposals"`
}

// NewDateProposals instantiates a new DateProposals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDateProposals(date string, proposals []TimeslotProposal) *DateProposals {
	this := DateProposals{}
	this.Date = date
	this.Proposals = proposals
	return &this
}

// NewDateProposalsWithDefaults instantiates a new DateProposals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDateProposalsWithDefaults() *DateProposals {
	this := DateProposals{}
	return &this
}

// GetDate returns the Date field value
func (o *DateProposals) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *DateProposals) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *DateProposals) SetDate(v string) {
	o.Date = v
}

// GetProposals returns the Proposals field value
func (o *DateProposals) GetProposals() []TimeslotProposal {
	if o == nil {
		var ret []TimeslotProposal
		return ret
	}

	return o.Proposals
}

// GetProposalsOk returns a tuple with the Proposals field value
// and a boolean to check if the value has been set.
func (o *DateProposals) GetProposalsOk() ([]TimeslotProposal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Proposals, true
}

// SetProposals sets field value
func (o *DateProposals) SetProposals(v []TimeslotProposal) {
	o.Proposals = v
}

func (o DateProposals) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DateProposals) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["proposals"] = o.Proposals
	return toSerialize, nil
}

type NullableDateProposals struct {
	value *DateProposals
	isSet bool
}

func (v NullableDateProposals) Get() *DateProposals {
	return v.value
}

func (v *NullableDateProposals) Set(val *DateProposals) {
	v.value = val
	v.isSet = true
}

func (v NullableDateProposals) IsSet() bool {
	return v.isSet
}

func (v *NullableDateProposals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDateProposals(val *DateProposals) *NullableDateProposals {
	return &NullableDateProposals{value: val, isSet: true}
}

func (v NullableDateProposals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDateProposals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


