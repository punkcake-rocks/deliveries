# Parcel1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The id of the parcel object in evermile db | [optional] 
**Dimensions** | [**Dimensions**](Dimensions.md) |  | 
**WeightKg** | **float32** | Weight of the package in Kilograms | 
**ParcelType** | [**ParcelType**](ParcelType.md) |  | [default to PARCELTYPE_PACKAGE]
**ItemsList** | Pointer to [**[]Item1**](Item1.md) |  | [optional] 

## Methods

### NewParcel1

`func NewParcel1(dimensions Dimensions, weightKg float32, parcelType ParcelType, ) *Parcel1`

NewParcel1 instantiates a new Parcel1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParcel1WithDefaults

`func NewParcel1WithDefaults() *Parcel1`

NewParcel1WithDefaults instantiates a new Parcel1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Parcel1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Parcel1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Parcel1) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Parcel1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDimensions

`func (o *Parcel1) GetDimensions() Dimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *Parcel1) GetDimensionsOk() (*Dimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *Parcel1) SetDimensions(v Dimensions)`

SetDimensions sets Dimensions field to given value.


### GetWeightKg

`func (o *Parcel1) GetWeightKg() float32`

GetWeightKg returns the WeightKg field if non-nil, zero value otherwise.

### GetWeightKgOk

`func (o *Parcel1) GetWeightKgOk() (*float32, bool)`

GetWeightKgOk returns a tuple with the WeightKg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightKg

`func (o *Parcel1) SetWeightKg(v float32)`

SetWeightKg sets WeightKg field to given value.


### GetParcelType

`func (o *Parcel1) GetParcelType() ParcelType`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *Parcel1) GetParcelTypeOk() (*ParcelType, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *Parcel1) SetParcelType(v ParcelType)`

SetParcelType sets ParcelType field to given value.


### GetItemsList

`func (o *Parcel1) GetItemsList() []Item1`

GetItemsList returns the ItemsList field if non-nil, zero value otherwise.

### GetItemsListOk

`func (o *Parcel1) GetItemsListOk() (*[]Item1, bool)`

GetItemsListOk returns a tuple with the ItemsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsList

`func (o *Parcel1) SetItemsList(v []Item1)`

SetItemsList sets ItemsList field to given value.

### HasItemsList

`func (o *Parcel1) HasItemsList() bool`

HasItemsList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


