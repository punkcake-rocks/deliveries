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

// checks if the WebhookGetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookGetResponse{}

// WebhookGetResponse The response for a get all webhooks request
type WebhookGetResponse struct {
	// The list of webhooks
	Webhooks []Webhook `json:"webhooks,omitempty"`
}

// NewWebhookGetResponse instantiates a new WebhookGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookGetResponse() *WebhookGetResponse {
	this := WebhookGetResponse{}
	return &this
}

// NewWebhookGetResponseWithDefaults instantiates a new WebhookGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookGetResponseWithDefaults() *WebhookGetResponse {
	this := WebhookGetResponse{}
	return &this
}

// GetWebhooks returns the Webhooks field value if set, zero value otherwise.
func (o *WebhookGetResponse) GetWebhooks() []Webhook {
	if o == nil || IsNil(o.Webhooks) {
		var ret []Webhook
		return ret
	}
	return o.Webhooks
}

// GetWebhooksOk returns a tuple with the Webhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookGetResponse) GetWebhooksOk() ([]Webhook, bool) {
	if o == nil || IsNil(o.Webhooks) {
		return nil, false
	}
	return o.Webhooks, true
}

// HasWebhooks returns a boolean if a field has been set.
func (o *WebhookGetResponse) HasWebhooks() bool {
	if o != nil && !IsNil(o.Webhooks) {
		return true
	}

	return false
}

// SetWebhooks gets a reference to the given []Webhook and assigns it to the Webhooks field.
func (o *WebhookGetResponse) SetWebhooks(v []Webhook) {
	o.Webhooks = v
}

func (o WebhookGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookGetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Webhooks) {
		toSerialize["webhooks"] = o.Webhooks
	}
	return toSerialize, nil
}

type NullableWebhookGetResponse struct {
	value *WebhookGetResponse
	isSet bool
}

func (v NullableWebhookGetResponse) Get() *WebhookGetResponse {
	return v.value
}

func (v *NullableWebhookGetResponse) Set(val *WebhookGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookGetResponse(val *WebhookGetResponse) *NullableWebhookGetResponse {
	return &NullableWebhookGetResponse{value: val, isSet: true}
}

func (v NullableWebhookGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


