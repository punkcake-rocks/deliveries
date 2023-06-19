# Address2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | **string** | Street and house number for the address | 
**AddressLine2** | Pointer to **string** | Additional address details such as flat number | [optional] 
**City** | **string** | The city for the address | 
**PostalCode** | **string** | The postal code for the address | 
**GeoLocation** | Pointer to [**GeoLocation**](GeoLocation.md) |  | [optional] 
**Type** | [**AddressType**](AddressType.md) |  | [default to ADDRESSTYPE_STORE]

## Methods

### NewAddress2

`func NewAddress2(addressLine1 string, city string, postalCode string, type_ AddressType, ) *Address2`

NewAddress2 instantiates a new Address2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddress2WithDefaults

`func NewAddress2WithDefaults() *Address2`

NewAddress2WithDefaults instantiates a new Address2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *Address2) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *Address2) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *Address2) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.


### GetAddressLine2

`func (o *Address2) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *Address2) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *Address2) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *Address2) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.

### GetCity

`func (o *Address2) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address2) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address2) SetCity(v string)`

SetCity sets City field to given value.


### GetPostalCode

`func (o *Address2) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Address2) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Address2) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetGeoLocation

`func (o *Address2) GetGeoLocation() GeoLocation`

GetGeoLocation returns the GeoLocation field if non-nil, zero value otherwise.

### GetGeoLocationOk

`func (o *Address2) GetGeoLocationOk() (*GeoLocation, bool)`

GetGeoLocationOk returns a tuple with the GeoLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoLocation

`func (o *Address2) SetGeoLocation(v GeoLocation)`

SetGeoLocation sets GeoLocation field to given value.

### HasGeoLocation

`func (o *Address2) HasGeoLocation() bool`

HasGeoLocation returns a boolean if a field has been set.

### GetType

`func (o *Address2) GetType() AddressType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Address2) GetTypeOk() (*AddressType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Address2) SetType(v AddressType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


