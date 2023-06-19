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

// checks if the Item type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Item{}

// Item A single item inside a package
type Item struct {
	// What the item is
	Description string `json:"description"`
	Value Price `json:"value"`
	// The number of these items in a parcel
	Quantity int32 `json:"quantity"`
	// An optional SKU for this item
	Sku *string `json:"sku,omitempty"`
	// Weight of the item in Kilograms
	WeightKg *float64 `json:"weightKg,omitempty"`
}

// NewItem instantiates a new Item object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItem(description string, value Price, quantity int32) *Item {
	this := Item{}
	this.Description = description
	this.Value = value
	this.Quantity = quantity
	return &this
}

// NewItemWithDefaults instantiates a new Item object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemWithDefaults() *Item {
	this := Item{}
	var quantity int32 = 1
	this.Quantity = quantity
	return &this
}

// GetDescription returns the Description field value
func (o *Item) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Item) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Item) SetDescription(v string) {
	o.Description = v
}

// GetValue returns the Value field value
func (o *Item) GetValue() Price {
	if o == nil {
		var ret Price
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Item) GetValueOk() (*Price, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Item) SetValue(v Price) {
	o.Value = v
}

// GetQuantity returns the Quantity field value
func (o *Item) GetQuantity() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *Item) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *Item) SetQuantity(v int32) {
	o.Quantity = v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *Item) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *Item) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *Item) SetSku(v string) {
	o.Sku = &v
}

// GetWeightKg returns the WeightKg field value if set, zero value otherwise.
func (o *Item) GetWeightKg() float64 {
	if o == nil || IsNil(o.WeightKg) {
		var ret float64
		return ret
	}
	return *o.WeightKg
}

// GetWeightKgOk returns a tuple with the WeightKg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Item) GetWeightKgOk() (*float64, bool) {
	if o == nil || IsNil(o.WeightKg) {
		return nil, false
	}
	return o.WeightKg, true
}

// HasWeightKg returns a boolean if a field has been set.
func (o *Item) HasWeightKg() bool {
	if o != nil && !IsNil(o.WeightKg) {
		return true
	}

	return false
}

// SetWeightKg gets a reference to the given float64 and assigns it to the WeightKg field.
func (o *Item) SetWeightKg(v float64) {
	o.WeightKg = &v
}

func (o Item) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Item) ToMap() (map[string]interface{}, error) {
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

type NullableItem struct {
	value *Item
	isSet bool
}

func (v NullableItem) Get() *Item {
	return v.value
}

func (v *NullableItem) Set(val *Item) {
	v.value = val
	v.isSet = true
}

func (v NullableItem) IsSet() bool {
	return v.isSet
}

func (v *NullableItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItem(val *Item) *NullableItem {
	return &NullableItem{value: val, isSet: true}
}

func (v NullableItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


