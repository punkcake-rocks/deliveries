# Item1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | What the item is | 
**Value** | [**Price1**](Price1.md) |  | 
**Quantity** | **int32** | The number of these items in a parcel | [default to 1]
**Sku** | Pointer to **string** | An optional SKU for this item | [optional] 
**WeightKg** | Pointer to **float64** | Weight of the item in Kilograms | [optional] 

## Methods

### NewItem1

`func NewItem1(description string, value Price1, quantity int32, ) *Item1`

NewItem1 instantiates a new Item1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItem1WithDefaults

`func NewItem1WithDefaults() *Item1`

NewItem1WithDefaults instantiates a new Item1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Item1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Item1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Item1) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetValue

`func (o *Item1) GetValue() Price1`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Item1) GetValueOk() (*Price1, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Item1) SetValue(v Price1)`

SetValue sets Value field to given value.


### GetQuantity

`func (o *Item1) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Item1) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Item1) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetSku

`func (o *Item1) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Item1) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Item1) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *Item1) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetWeightKg

`func (o *Item1) GetWeightKg() float64`

GetWeightKg returns the WeightKg field if non-nil, zero value otherwise.

### GetWeightKgOk

`func (o *Item1) GetWeightKgOk() (*float64, bool)`

GetWeightKgOk returns a tuple with the WeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightKg

`func (o *Item1) SetWeightKg(v float64)`

SetWeightKg sets WeightKg field to given value.

### HasWeightKg

`func (o *Item1) HasWeightKg() bool`

HasWeightKg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


