# HandoffContactInfo1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Phone** | **string** | A phone number of the handoff receiver (leader) | 
**Address** | [**AddressComponents**](AddressComponents.md) |  | 

## Methods

### NewHandoffContactInfo1

`func NewHandoffContactInfo1(name string, phone string, address AddressComponents, ) *HandoffContactInfo1`

NewHandoffContactInfo1 instantiates a new HandoffContactInfo1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandoffContactInfo1WithDefaults

`func NewHandoffContactInfo1WithDefaults() *HandoffContactInfo1`

NewHandoffContactInfo1WithDefaults instantiates a new HandoffContactInfo1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HandoffContactInfo1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandoffContactInfo1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandoffContactInfo1) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *HandoffContactInfo1) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *HandoffContactInfo1) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *HandoffContactInfo1) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetAddress

`func (o *HandoffContactInfo1) GetAddress() AddressComponents`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HandoffContactInfo1) GetAddressOk() (*AddressComponents, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HandoffContactInfo1) SetAddress(v AddressComponents)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


