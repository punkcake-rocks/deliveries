# AddressComponents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | Pointer to **string** | Street and house number for the address | [optional] 
**City** | Pointer to **string** | The city for the address | [optional] 
**PostalCode** | Pointer to **string** | The postal code for the address | [optional] 

## Methods

### NewAddressComponents

`func NewAddressComponents() *AddressComponents`

NewAddressComponents instantiates a new AddressComponents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressComponentsWithDefaults

`func NewAddressComponentsWithDefaults() *AddressComponents`

NewAddressComponentsWithDefaults instantiates a new AddressComponents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressLine1

`func (o *AddressComponents) GetAddressLine1() string`

GetAddressLine1 returns the AddressLine1 field if non-nil, zero value otherwise.

### GetAddressLine1Ok

`func (o *AddressComponents) GetAddressLine1Ok() (*string, bool)`

GetAddressLine1Ok returns a tuple with the AddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine1

`func (o *AddressComponents) SetAddressLine1(v string)`

SetAddressLine1 sets AddressLine1 field to given value.

### HasAddressLine1

`func (o *AddressComponents) HasAddressLine1() bool`

HasAddressLine1 returns a boolean if a field has been set.

### GetCity

`func (o *AddressComponents) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressComponents) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressComponents) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *AddressComponents) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetPostalCode

`func (o *AddressComponents) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressComponents) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressComponents) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *AddressComponents) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


