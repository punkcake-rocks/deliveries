# InvalidAddressBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | The error message | [optional] 
**Input** | [**AddressComponents**](AddressComponents.md) |  | 

## Methods

### NewInvalidAddressBody

`func NewInvalidAddressBody(input AddressComponents, ) *InvalidAddressBody`

NewInvalidAddressBody instantiates a new InvalidAddressBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidAddressBodyWithDefaults

`func NewInvalidAddressBodyWithDefaults() *InvalidAddressBody`

NewInvalidAddressBodyWithDefaults instantiates a new InvalidAddressBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InvalidAddressBody) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvalidAddressBody) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvalidAddressBody) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InvalidAddressBody) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetInput

`func (o *InvalidAddressBody) GetInput() AddressComponents`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *InvalidAddressBody) GetInputOk() (*AddressComponents, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *InvalidAddressBody) SetInput(v AddressComponents)`

SetInput sets Input field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


