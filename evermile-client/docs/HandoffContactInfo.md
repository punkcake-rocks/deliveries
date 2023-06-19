# HandoffContactInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Phone** | **string** | A phone number of the handoff receiver (leader) | 
**Address** | [**AddressComponents**](AddressComponents.md) |  | 

## Methods

### NewHandoffContactInfo

`func NewHandoffContactInfo(name string, phone string, address AddressComponents, ) *HandoffContactInfo`

NewHandoffContactInfo instantiates a new HandoffContactInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHandoffContactInfoWithDefaults

`func NewHandoffContactInfoWithDefaults() *HandoffContactInfo`

NewHandoffContactInfoWithDefaults instantiates a new HandoffContactInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HandoffContactInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HandoffContactInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HandoffContactInfo) SetName(v string)`

SetName sets Name field to given value.


### GetPhone

`func (o *HandoffContactInfo) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *HandoffContactInfo) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *HandoffContactInfo) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetAddress

`func (o *HandoffContactInfo) GetAddress() AddressComponents`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HandoffContactInfo) GetAddressOk() (*AddressComponents, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HandoffContactInfo) SetAddress(v AddressComponents)`

SetAddress sets Address field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


