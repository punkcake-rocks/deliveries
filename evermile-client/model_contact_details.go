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

// checks if the ContactDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactDetails{}

// ContactDetails The contact details at the address
type ContactDetails struct {
	// The name of the contact person at address
	ContactName string `json:"contactName"`
	// The phone number of the contact person at address
	ContactPhone *string `json:"contactPhone,omitempty"`
	// The email address of the contact person at address
	ContactEmail string `json:"contactEmail"`
	// The name of the company at address
	LocationCompanyName *string `json:"locationCompanyName,omitempty"`
	// Instructions for delivery
	Instructions *string `json:"instructions,omitempty"`
	// Additional contact address details such as flat number
	AddressLine2 *string `json:"addressLine2,omitempty"`
}

// NewContactDetails instantiates a new ContactDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactDetails(contactName string, contactEmail string) *ContactDetails {
	this := ContactDetails{}
	this.ContactName = contactName
	this.ContactEmail = contactEmail
	return &this
}

// NewContactDetailsWithDefaults instantiates a new ContactDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactDetailsWithDefaults() *ContactDetails {
	this := ContactDetails{}
	return &this
}

// GetContactName returns the ContactName field value
func (o *ContactDetails) GetContactName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContactName
}

// GetContactNameOk returns a tuple with the ContactName field value
// and a boolean to check if the value has been set.
func (o *ContactDetails) GetContactNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactName, true
}

// SetContactName sets field value
func (o *ContactDetails) SetContactName(v string) {
	o.ContactName = v
}

// GetContactPhone returns the ContactPhone field value if set, zero value otherwise.
func (o *ContactDetails) GetContactPhone() string {
	if o == nil || IsNil(o.ContactPhone) {
		var ret string
		return ret
	}
	return *o.ContactPhone
}

// GetContactPhoneOk returns a tuple with the ContactPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactDetails) GetContactPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.ContactPhone) {
		return nil, false
	}
	return o.ContactPhone, true
}

// HasContactPhone returns a boolean if a field has been set.
func (o *ContactDetails) HasContactPhone() bool {
	if o != nil && !IsNil(o.ContactPhone) {
		return true
	}

	return false
}

// SetContactPhone gets a reference to the given string and assigns it to the ContactPhone field.
func (o *ContactDetails) SetContactPhone(v string) {
	o.ContactPhone = &v
}

// GetContactEmail returns the ContactEmail field value
func (o *ContactDetails) GetContactEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value
// and a boolean to check if the value has been set.
func (o *ContactDetails) GetContactEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactEmail, true
}

// SetContactEmail sets field value
func (o *ContactDetails) SetContactEmail(v string) {
	o.ContactEmail = v
}

// GetLocationCompanyName returns the LocationCompanyName field value if set, zero value otherwise.
func (o *ContactDetails) GetLocationCompanyName() string {
	if o == nil || IsNil(o.LocationCompanyName) {
		var ret string
		return ret
	}
	return *o.LocationCompanyName
}

// GetLocationCompanyNameOk returns a tuple with the LocationCompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactDetails) GetLocationCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.LocationCompanyName) {
		return nil, false
	}
	return o.LocationCompanyName, true
}

// HasLocationCompanyName returns a boolean if a field has been set.
func (o *ContactDetails) HasLocationCompanyName() bool {
	if o != nil && !IsNil(o.LocationCompanyName) {
		return true
	}

	return false
}

// SetLocationCompanyName gets a reference to the given string and assigns it to the LocationCompanyName field.
func (o *ContactDetails) SetLocationCompanyName(v string) {
	o.LocationCompanyName = &v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *ContactDetails) GetInstructions() string {
	if o == nil || IsNil(o.Instructions) {
		var ret string
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactDetails) GetInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *ContactDetails) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given string and assigns it to the Instructions field.
func (o *ContactDetails) SetInstructions(v string) {
	o.Instructions = &v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *ContactDetails) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactDetails) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *ContactDetails) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *ContactDetails) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

func (o ContactDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contactName"] = o.ContactName
	if !IsNil(o.ContactPhone) {
		toSerialize["contactPhone"] = o.ContactPhone
	}
	toSerialize["contactEmail"] = o.ContactEmail
	if !IsNil(o.LocationCompanyName) {
		toSerialize["locationCompanyName"] = o.LocationCompanyName
	}
	if !IsNil(o.Instructions) {
		toSerialize["instructions"] = o.Instructions
	}
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	return toSerialize, nil
}

type NullableContactDetails struct {
	value *ContactDetails
	isSet bool
}

func (v NullableContactDetails) Get() *ContactDetails {
	return v.value
}

func (v *NullableContactDetails) Set(val *ContactDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableContactDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableContactDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactDetails(val *ContactDetails) *NullableContactDetails {
	return &NullableContactDetails{value: val, isSet: true}
}

func (v NullableContactDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

