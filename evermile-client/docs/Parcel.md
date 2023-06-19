# Parcel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the parcel object in evermile db | [optional] 
**Dimensions** | [**Dimensions**](Dimensions.md) |  | 
**WeightKg** | **float32** | Weight of the package in Kilograms | 
**ParcelType** | [**ParcelType**](ParcelType.md) |  | [default to PARCELTYPE_PACKAGE]
**ItemsList** | Pointer to [**[]Item**](Item.md) |  | [optional] 

## Methods

### NewParcel

`func NewParcel(dimensions Dimensions, weightKg float32, parcelType ParcelType, ) *Parcel`

NewParcel instantiates a new Parcel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcelWithDefaults

`func NewParcelWithDefaults() *Parcel`

NewParcelWithDefaults instantiates a new Parcel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Parcel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Parcel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Parcel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Parcel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDimensions

`func (o *Parcel) GetDimensions() Dimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *Parcel) GetDimensionsOk() (*Dimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *Parcel) SetDimensions(v Dimensions)`

SetDimensions sets Dimensions field to given value.


### GetWeightKg

`func (o *Parcel) GetWeightKg() float32`

GetWeightKg returns the WeightKg field if non-nil, zero value otherwise.

### GetWeightKgOk

`func (o *Parcel) GetWeightKgOk() (*float32, bool)`

GetWeightKgOk returns a tuple with the WeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightKg

`func (o *Parcel) SetWeightKg(v float32)`

SetWeightKg sets WeightKg field to given value.


### GetParcelType

`func (o *Parcel) GetParcelType() ParcelType`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *Parcel) GetParcelTypeOk() (*ParcelType, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *Parcel) SetParcelType(v ParcelType)`

SetParcelType sets ParcelType field to given value.


### GetItemsList

`func (o *Parcel) GetItemsList() []Item`

GetItemsList returns the ItemsList field if non-nil, zero value otherwise.

### GetItemsListOk

`func (o *Parcel) GetItemsListOk() (*[]Item, bool)`

GetItemsListOk returns a tuple with the ItemsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsList

`func (o *Parcel) SetItemsList(v []Item)`

SetItemsList sets ItemsList field to given value.

### HasItemsList

`func (o *Parcel) HasItemsList() bool`

HasItemsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


