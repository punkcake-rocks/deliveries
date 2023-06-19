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

// checks if the Item1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Item1{}

// Item1 A single item inside a package
type Item1 struct {
	// What the item is
	Description string `json:"description"`
	Value Price1 `json:"value"`
	// The number of these items in a parcel
	Quantity int32 `json:"quantity"`
	// An optional SKU for this item
	Sku *string `json:"sku,omitempty"`
	// Weight of the item in Kilograms
	WeightKg *float64 `json:"weightKg,omitempty"`
}

// NewItem1 instantiates a new Item1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItem1(description string, value Price1, quantity int32) *Item1 {
	this := Item1{}
	this.Description = description
	this.Value = value
	this.Quantity = quantity
	return &this
}

// NewItem1WithDefaults instantiates a new Item1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItem1WithDefaults() *Item1 {
	this := Item1{}
	var quantity int32 = 1
	this.Quantity = quantity
	return &this
}

// GetDescription returns the Description field value
func (o *Item1) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Item1) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Item1) SetDescription(v string) {
	o.Description = v
}

// GetValue returns the Value field value
func (o *Item1) GetValue() Price1 {
	if o == nil {
		var ret Price1
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Item1) GetValueOk() (*Price1, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Item1) SetValue(v Price1) {
	o.Value = v
}

// GetQuantity returns the Quantity field value
func (o *Item1) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *Item1) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *Item1) SetQuantity(v int32) {
	o.Quantity = v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *Item1) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item1) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *Item1) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *Item1) SetSku(v string) {
	o.Sku = &v
}

// GetWeightKg returns the WeightKg field value if set, zero value otherwise.
func (o *Item1) GetWeightKg() float64 {
	if o == nil || IsNil(o.WeightKg) {
		var ret float64
		return ret
	}
	return *o.WeightKg
}

// GetWeightKgOk returns a tuple with the WeightKg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item1) GetWeightKgOk() (*float64, bool) {
	if o == nil || IsNil(o.WeightKg) {
		return nil, false
	}
	return o.WeightKg, true
}

// HasWeightKg returns a boolean if a field has been set.
func (o *Item1) HasWeightKg() bool {
	if o != nil && !IsNil(o.WeightKg) {
		return true
	}

	return false
}

// SetWeightKg gets a reference to the given float64 and assigns it to the WeightKg field.
func (o *Item1) SetWeightKg(v float64) {
	o.WeightKg = &v
}

func (o Item1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Item1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["value"] = o.Value
	toSerialize["quantity"] = o.Quantity
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.WeightKg) {
		toSerialize["weightKg"] = o.WeightKg
	}
	return toSerialize, nil
}

type NullableItem1 struct {
	value *Item1
	isSet bool
}

func (v NullableItem1) Get() *Item1 {
	return v.value
}

func (v *NullableItem1) Set(val *Item1) {
	v.value = val
	v.isSet = true
}

func (v NullableItem1) IsSet() bool {
	return v.isSet
}

func (v *NullableItem1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItem1(val *Item1) *NullableItem1 {
	return &NullableItem1{value: val, isSet: true}
}

func (v NullableItem1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItem1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

