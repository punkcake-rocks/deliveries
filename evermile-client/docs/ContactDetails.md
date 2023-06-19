# ContactDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactName** | **string** | The name of the contact person at address | 
**ContactPhone** | Pointer to **string** | The phone number of the contact person at address | [optional] 
**ContactEmail** | **string** | The email address of the contact person at address | 
**LocationCompanyName** | Pointer to **string** | The name of the company at address | [optional] 
**Instructions** | Pointer to **string** | Instructions for delivery | [optional] 
**AddressLine2** | Pointer to **string** | Additional contact address details such as flat number | [optional] 

## Methods

### NewContactDetails

`func NewContactDetails(contactName string, contactEmail string, ) *ContactDetails`

NewContactDetails instantiates a new ContactDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactDetailsWithDefaults

`func NewContactDetailsWithDefaults() *ContactDetails`

NewContactDetailsWithDefaults instantiates a new ContactDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactName

`func (o *ContactDetails) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *ContactDetails) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *ContactDetails) SetContactName(v string)`

SetContactName sets ContactName field to given value.


### GetContactPhone

`func (o *ContactDetails) GetContactPhone() string`

GetContactPhone returns the ContactPhone field if non-nil, zero value otherwise.

### GetContactPhoneOk

`func (o *ContactDetails) GetContactPhoneOk() (*string, bool)`

GetContactPhoneOk returns a tuple with the ContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPhone

`func (o *ContactDetails) SetContactPhone(v string)`

SetContactPhone sets ContactPhone field to given value.

### HasContactPhone

`func (o *ContactDetails) HasContactPhone() bool`

HasContactPhone returns a boolean if a field has been set.

### GetContactEmail

`func (o *ContactDetails) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *ContactDetails) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *ContactDetails) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.


### GetLocationCompanyName

`func (o *ContactDetails) GetLocationCompanyName() string`

GetLocationCompanyName returns the LocationCompanyName field if non-nil, zero value otherwise.

### GetLocationCompanyNameOk

`func (o *ContactDetails) GetLocationCompanyNameOk() (*string, bool)`

GetLocationCompanyNameOk returns a tuple with the LocationCompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationCompanyName

`func (o *ContactDetails) SetLocationCompanyName(v string)`

SetLocationCompanyName sets LocationCompanyName field to given value.

### HasLocationCompanyName

`func (o *ContactDetails) HasLocationCompanyName() bool`

HasLocationCompanyName returns a boolean if a field has been set.

### GetInstructions

`func (o *ContactDetails) GetInstructions() string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *ContactDetails) GetInstructionsOk() (*string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *ContactDetails) SetInstructions(v string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *ContactDetails) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.

### GetAddressLine2

`func (o *ContactDetails) GetAddressLine2() string`

GetAddressLine2 returns the AddressLine2 field if non-nil, zero value otherwise.

### GetAddressLine2Ok

`func (o *ContactDetails) GetAddressLine2Ok() (*string, bool)`

GetAddressLine2Ok returns a tuple with the AddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressLine2

`func (o *ContactDetails) SetAddressLine2(v string)`

SetAddressLine2 sets AddressLine2 field to given value.

### HasAddressLine2

`func (o *ContactDetails) HasAddressLine2() bool`

HasAddressLine2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


