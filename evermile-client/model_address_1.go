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

// checks if the Address1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Address1{}

// Address1 An address for picking up or dropping off
type Address1 struct {
	// Street and house number for the address
	AddressLine1 string `json:"addressLine1"`
	// Additional address details such as flat number
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// The city for the address
	City string `json:"city"`
	// The postal code for the address
	PostalCode string `json:"postalCode"`
	GeoLocation *GeoLocation `json:"geoLocation,omitempty"`
	Type AddressType `json:"type"`
}

// NewAddress1 instantiates a new Address1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress1(addressLine1 string, city string, postalCode string, type_ AddressType) *Address1 {
	this := Address1{}
	this.AddressLine1 = addressLine1
	this.City = city
	this.PostalCode = postalCode
	this.Type = type_
	return &this
}

// NewAddress1WithDefaults instantiates a new Address1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddress1WithDefaults() *Address1 {
	this := Address1{}
	var type_ AddressType = ADDRESSTYPE_STORE
	this.Type = type_
	return &this
}

// GetAddressLine1 returns the AddressLine1 field value
func (o *Address1) GetAddressLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AddressLine1
}

// GetAddressLine1Ok returns a tuple with the AddressLine1 field value
// and a boolean to check if the value has been set.
func (o *Address1) GetAddressLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddressLine1, true
}

// SetAddressLine1 sets field value
func (o *Address1) SetAddressLine1(v string) {
	o.AddressLine1 = v
}

// GetAddressLine2 returns the AddressLine2 field value if set, zero value otherwise.
func (o *Address1) GetAddressLine2() string {
	if o == nil || IsNil(o.AddressLine2) {
		var ret string
		return ret
	}
	return *o.AddressLine2
}

// GetAddressLine2Ok returns a tuple with the AddressLine2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address1) GetAddressLine2Ok() (*string, bool) {
	if o == nil || IsNil(o.AddressLine2) {
		return nil, false
	}
	return o.AddressLine2, true
}

// HasAddressLine2 returns a boolean if a field has been set.
func (o *Address1) HasAddressLine2() bool {
	if o != nil && !IsNil(o.AddressLine2) {
		return true
	}

	return false
}

// SetAddressLine2 gets a reference to the given string and assigns it to the AddressLine2 field.
func (o *Address1) SetAddressLine2(v string) {
	o.AddressLine2 = &v
}

// GetCity returns the City field value
func (o *Address1) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *Address1) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *Address1) SetCity(v string) {
	o.City = v
}

// GetPostalCode returns the PostalCode field value
func (o *Address1) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *Address1) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *Address1) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetGeoLocation returns the GeoLocation field value if set, zero value otherwise.
func (o *Address1) GetGeoLocation() GeoLocation {
	if o == nil || IsNil(o.GeoLocation) {
		var ret GeoLocation
		return ret
	}
	return *o.GeoLocation
}

// GetGeoLocationOk returns a tuple with the GeoLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address1) GetGeoLocationOk() (*GeoLocation, bool) {
	if o == nil || IsNil(o.GeoLocation) {
		return nil, false
	}
	return o.GeoLocation, true
}

// HasGeoLocation returns a boolean if a field has been set.
func (o *Address1) HasGeoLocation() bool {
	if o != nil && !IsNil(o.GeoLocation) {
		return true
	}

	return false
}

// SetGeoLocation gets a reference to the given GeoLocation and assigns it to the GeoLocation field.
func (o *Address1) SetGeoLocation(v GeoLocation) {
	o.GeoLocation = &v
}

// GetType returns the Type field value
func (o *Address1) GetType() AddressType {
	if o == nil {
		var ret AddressType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Address1) GetTypeOk() (*AddressType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Address1) SetType(v AddressType) {
	o.Type = v
}

func (o Address1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Address1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addressLine1"] = o.AddressLine1
	if !IsNil(o.AddressLine2) {
		toSerialize["addressLine2"] = o.AddressLine2
	}
	toSerialize["city"] = o.City
	toSerialize["postalCode"] = o.PostalCode
	if !IsNil(o.GeoLocation) {
		toSerialize["geoLocation"] = o.GeoLocation
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAddress1 struct {
	value *Address1
	isSet bool
}

func (v NullableAddress1) Get() *Address1 {
	return v.value
}

func (v *NullableAddress1) Set(val *Address1) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress1) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress1(val *Address1) *NullableAddress1 {
	return &NullableAddress1{value: val, isSet: true}
}

func (v NullableAddress1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


